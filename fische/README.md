# fische

Minecraft Server for friends.

## Initial setup of server

In the fly org for minecraft servers, with the fly cli installed and authenticated I did:

```bash
flyctl launch
fly ips allocate-v4
```

A CNAME record points my friends to the domain fly.io allocates for you which in turn resolves to the dedicated IP. 

A dedicated IP is required for UPD traffic on fly.io.

### Copy existing world data

If you have an existing world, get it on your local computer first. Open a terminal in the directory you have the zip file.

Then connect using SFTP to the machine in your app: `fly sftp shell -a minecraft-fische`.

Run a  `cd /data` and `put my-fancy-world.zip` to upload the zip file to the data directory on the machine.

Then run `fly ssh console -a minecraft-fische` to get a shell in your machine.

Unzip the zip file in the `/data` diretory. Finally rename the folder to `world` and thus override the existing world.

Now do a `fly machine restart -a minecraft-fische` and once the server has restared the world should be loaded.
