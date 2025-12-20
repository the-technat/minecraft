# Minecraft Servers

A repo containing all things Minecraft.

I'm not a gamer, but as a system engineer I help my friends by setting up their servers so that they don't have to.

All deployed to a fly.io organization, each folder a server.

## Initial deployment

Initial deployment was done according to [this doc](https://fly.io/docs/launch/continuous-deployment-with-github-actions/). The workflow runs against the `fische` env on the `main` branch or any other env, given the server.

An organization token was generated manually and added to the Github repository.

Don't forget to specify the right organization using the  `-o Minecraft` flag.

A dedicated IPv6 address was manaually allocated to the app using `fly ips allocate-v6` and there's an AAAA record pointing to that IP. This means the servers are IPv6 only which means we are not paying for a dedicated IPv4 address. This works as plain TCP traffic is supported by fly.io on IPv6.

### Copy existing world data

If you have an existing world, get it on your local computer first. Open a terminal in the directory you have the zip file.

Then connect using SFTP to the machine in your app: `fly sftp shell -a minecraft-flasche`.

Run a  `cd /data` and `put my-fancy-world.zip` to upload the zip file to the data directory on the machine.

Then run `fly ssh console -a minecraft-flasche` to get a shell in your machine.

Unzip the zip file in the `/data` diretory. Finally rename the folder to `world` and thus override the existing world.

Now do a `fly machine restart -a minecraft-flasche` and once the server has restared the world should be loaded.

## Credits

[@yamatt](https://github.com/yamatt) for his [fly-minecraft-server](https://github.com/yamatt/fly-minecraft-server) project which served as inspiration how to deploy the servers