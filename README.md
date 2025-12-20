# Minecraft Servers

A repo containing all things Minecraft.

I'm not a gamer, but as a system engineer I help my friends by setting up their servers so that they don't have to.

## Minecraft Servers

All deployed to a fly.io organization, the folders all contain a `fly.toml` to deploy.

Deployment's are done manually at the time, no Github Action.

## Initial deployment

Initial deployment was done according to [this doc](https://fly.io/docs/launch/continuous-deployment-with-github-actions/). The workflow runs against the `fische` env on the `main` branch.

Don't forget to specify the right organization!

Secrets have been all generated manually and added to this repository.

A CNAME record points my friends to the domain fly.io allocates for you which in turn resolves to the dedicated IPv4 and shared IPv4. 

### Copy existing world data

If you have an existing world, get it on your local computer first. Open a terminal in the directory you have the zip file.

Then connect using SFTP to the machine in your app: `fly sftp shell -a minecraft-flasche`.

Run a  `cd /data` and `put my-fancy-world.zip` to upload the zip file to the data directory on the machine.

Then run `fly ssh console -a minecraft-flasche` to get a shell in your machine.

Unzip the zip file in the `/data` diretory. Finally rename the folder to `world` and thus override the existing world.

Now do a `fly machine restart -a minecraft-flasche` and once the server has restared the world should be loaded.

## Credits

[@yamatt](https://github.com/yamatt) for his [fly-minecraft-server](https://github.com/yamatt/fly-minecraft-server) project which served as inspiration how to deploy the servers