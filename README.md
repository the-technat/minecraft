# Minecraft Servers

Minecraft Server hosting for my friends.

> I'm not a gamer, but as a system engineer I help my friends by setting up their servers so that they don't have to.

## Goal

Setup as many minecraft servers as I want with as few effort as possible and as low cost as possible. If there's effort then in the begining, but there shouldn't be any maintenance effort.

The server used for this must be cattle, in case it dies we just got downtime till a new server is setup.

## Setup

We need some sort of server, either an old PC sitting under your desk, a VPS or something else.

On this box I ensure:
- docker is installed
- this repo is cloned at `~/minecraft`
- the following directories are populated with the current world data:
  - `~/minecraft/fische_data`
  - `~/minecraft/flasche_data`
  - see the "Restore" chapter how to get that data back to the server from the latest backup
- start the stack using `docker compose up -d`
- expose the container `mc_router`'s ``25565/udp` is exposed to the world, and the following domains are pointing to it:
  - `flasche.alleaffengaffen.ch 300 IN A IP`
  - `fische.alleaffengaffen.ch 300 IN A IP`
  - `_minecraft._tcp.flasche.alleaffengaffen.ch IN SRV 10 100 PORT FQDN.`
  _ `_minecraft._tcp.fische.alleaffengaffen.ch IN SRV 10 100 PORT FQDN.`

### playit.gg

One way to expose `mc_router` to the internet without paying for a public IP.

Setup:
- Register at `playit.gg/login/create`
- Run the following commands on your server:
  ```console
  curl -SsL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/playit.gpg >/dev/null
  echo "deb [signed-by=/etc/apt/trusted.gpg.d/playit.gpg] https://playit-cloud.github.io/ppa/data ./" | sudo tee /etc/apt/sources.list.d/playit-cloud.list
  sudo apt update
  sudo apt install playit
  sudo systemctl enable --now playit
  ```
- Configure the agent by running `playit setup` (this will prompt you to enter a code in your browser and will then configure the playit agent) 
- Create a tunnel in their UI using "Global Anycast" and type "Minecraft Java (game)".
- Use the IP and Port they give you to configure DNS

### Backups

There are preconfigured backup containers doing regular backups to Openstack Swift (Infomaniak). To finish their setup, create a `.env` file with the credentials needed:

```console
cat <<EOF | tee ./.env
RESTIC_PASSWORD=<restic password>
RESTIC_ADDITIONAL_TAGS=banana
OS_AUTH_URL=https://api.pub1.infomaniak.cloud/identity/v3
OS_PROJECT_ID="<openstack project id>"
OS_PROJECT_NAME="<openstack project name>"
OS_USER_DOMAIN_NAME="Default"
OS_PROJECT_DOMAIN_ID="default"
OS_USERNAME="<openstack username>"
OS_PASSWORD="<password of openstack user>"
OS_REGION_NAME="dc4-a"
OS_INTERFACE=public
OS_IDENTITY_API_VERSION=3
OS_AUTH_VERSION=3
EOF
```

Also set the `RCON_PASSWORD` variable to something static:

```console
cat <<EOF | tee .fische-env 
RCON_PASSWORD=""
EOF
cat <<EOF | tee .flasche-env 
RCON_PASSWORD=""
EOF
``` 

### Restore

Use the following oneshot container to restore the latest world data on a fresh/existing server:
```console
docker run --rm -ti -v ./restore_location:/restore --env-file .env -e RESTIC_REPOSITORY="swift:minecraft-backups:/server_name" restic/restic restore latest --target /restore
```
