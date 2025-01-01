# Telegram Minecraft Manager

A Telegram Bot to manage Minecraft Servers on fly.io

This is WIP, no idea if it will ever be finished.

## Idea

I'd like to overengineer my job as the tech nerd who can setup minecraft servers for friends.

The idea is to have a bot that can do most jobs I or my friends would like to do on a Minecraft Server.

This includes:
- start a server
- stop a server
- provision IP/DNS for a server 
- create world dump 
- restore world dump
- Optional: provision new server
- Optional: destroy server

### Dependencies & Assumptions

Since I don't want to engineer everything myself I decided to depend on several third-parties to make my life easier:

- fly.io: hosting provider with reasonable prices and all features you need
- [itzg/docker-minecraft-server](https://github.com/itzg/docker-minecraft-server): ready-to-use docker image, avoids having to deal with java and shit directly

### Config

My bot only reads env vars to make things easier. But to keep long-term track of world data and managed servers, the bot needs a dedicated S3 bucket to store things.

See [pkg/config/config.go:12](./pkg/config/config.go:12) for a list of environment variables the bot accepts.

### RBAC

Currently the bot can be configured with a list of admins who can do stuff and everyone else on Telegram can't do anything. Later I'd like to add more fine-graded RBAC to limit people to specific servers and maybe also actions.

## Initial deployment

The following commands will do:


```console
fly launch
fly secrets set TMM_BOT_TOKEN=""
fly secrets set TMM_BOT_WEBHOOK_TOKEN=""
fly secrets set TMM_S3_ACCESS_KEY=""
fly secrets set TMM_S3_SECRET_KEY=""
fly secrets set TMM_S3_BUCKET=""
```