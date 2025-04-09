# Minecraft Servers

A repo containing all things Minecraft.

I'm not a gamer, but as a system engineer I help my friends by setting up their servers so that they don't have to.

## Setup

0. Install docker on a server
1. Clone the repo to a new server.
2. Get data to the directories (e.g scp,unzip,move)
3. `docker compose up -d`
4. Expose the IP/Port of the router container somehow to the world (check docker labels for the domains to use)

### PlayIT

 nice little reverse-tunnel service to expose the servers (if the host is running behind a NAT).

Register, login, install their agent:

```console
curl -SsL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/playit.gpg >/dev/null
echo "deb [signed-by=/etc/apt/trusted.gpg.d/playit.gpg] https://playit-cloud.github.io/ppa/data ./" | sudo tee /etc/apt/sources.list.d/playit-cloud.list
sudo apt update
sudo apt install playit
sudo systemctl enable --now playit
playit setup # will prompt, then write config file for systemd service
```

Lastly create a tunnel in ther UI.

### DNS

The following two records are used for a server:
```
@ 300  IN A  147.185.221.27 -> subzone flasche.alleaffengaffen.ch 
_minecraft 300  IN SRV 10 100 23373 flasche.alleaffengaffen.ch. -> subzone _tcp.flasche.alleaffengaffen.ch
```

## Backups

There are preconfigured backup containers doing regular backups to S3.

To finish the setup, create the `.env` file:

```console
cat <<EOF | tee ./.env
RESTIC_REPOSITORY=swift:minecraft-backups:/
RESTIC_PASSWORD=<password for backups>
RESTIC_ADDITIONAL_TAGS=banana
## + all the vars from openstack.rc you got from the openstack dashboard
EOF
```
