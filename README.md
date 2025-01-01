# Fly Minecraft Servers

A repo containing all things Minecraft.

I'm not a gamer, but as a system engineer I help my friends by setting up their servers so that they don't have to.

## Minecraft Servers

All deployed to a fly.io organization, the folders all contain a `fly.toml` to deploy.

Deployment's are done manually at the time, no Github Action.

## Telegram Minecraft Manager

The [tmm](./tmm) folder contains a draft of a Telegram bot to manage these Minecraft servers. But it's by far not done.

## Initial setup of server

You will need a fly.io account, then install and setup the [fly.io CLI](https://fly.io/docs/flyctl/install/) then run:

```bash
flyctl launch 
```

If asked to provision a dedicated IP address say yes or use `fly ips allocate-v4` later to get your IP.

As soon as you have an IP, update your DNS A record accordingly. 

### Copy existing world data

If you have an existing world, get it on your local computer first. Open a terminal in the directory you have the zip file.

Then connect using SFTP to the machine in your app: `fly sftp shell -a minecraft-flasche`.

Run a  `cd /data` and `put my-fancy-world.zip` to upload the zip file to the correct directory.

Then run `fly ssh console -a minecraft-flasche` to get a shell in your machine.

Unzip the zip file in the diretory. Finally rename the folder to `world` and thus override the existing world.

Now do a `fly machine restart -a minecraft-flasche` and once the server has restared the world should be loaded.
