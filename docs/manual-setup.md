# Manual Minecraft Setup

Server: any VPS with fixed IPv4 and ~10GB system volume.

Default Installation of Ubuntu Server 24.04 with cloud-init file:

```yaml
#cloud-config fische

locale: en_US.UTF-8
timezone: UTC
package_update: true
package_upgrade: true
packages:
- vim
- git
- wget
- curl
- dnsutils
- net-tools
- unzip

users:
  - name: technat
    groups: sudo
    sudo: ALL=(ALL) NOPASSWD:ALL # Allow any operations using sudo
    gecos: "Admin user created by cloud-init"
    shell: /usr/bin/bash

write_files:
- path: /etc/sysctl.d/99-tailscale.conf
  content: |
    net.ipv4.ip_forward          = 1
    net.ipv6.conf.all.forwarding = 1      
runcmd:
  - sysctl -p /etc/sysctl.d/99-tailscale.conf
  - systemctl mask ssh
  - curl -fsSL https://tailscale.com/install.sh | sh
  - tailscale up --ssh --auth-key "<single-use-pre-approved-key>"
  - ufw default deny incoming
  - ufw default allow outgoing
  - ufw allow in on tailscale0
  - ufw --force enable
```

Additional firewall rules:
- `sudo ufw allow 25565`

Minecraft Server setup:
```console
sudo apt install openjdk-21-jre-headless
sudo useradd -r -U -d /opt/minecraft -s /usr/sbin/nologin minecraft
sudo mkdir -p /opt/minecraft
wget <url_of_server_jar> -O /tmp/server.jar
sudo mv /tmp/server.jar /opt/minecraft
sudo chown -R minecraft: /opt/minecraft
sudo usermod -aG minecraft technat # add me to the minecraft group for easy edditing of files
sudo -u minecraft bash 
cd /opt/minecraft
java -Xmx1024M -Xms1024M -jar server.jar nogui # run as minecraft user in new shell
sed -i 's/\bfalse\b/TRUE/' eula.txt
```

Also installed mcrcon according to it’s README: https://github.com/Tiiffi/mcrcon (requires git, gcc and make).

My server.properties:
```
# https://minecraft.fandom.com/wiki/Server.properties
#Minecraft server properties
#Thu Dec 26 12:28:27 UTC 2024
accepts-transfers=false
allow-flight=true
allow-nether=true
broadcast-console-to-ops=true
broadcast-rcon-to-ops=true
bug-report-link=
difficulty=normal
enable-command-block=true
enable-jmx-monitoring=false
enable-query=false
enable-rcon=true
enable-status=true
enforce-secure-profile=true
enforce-whitelist=true
entity-broadcast-range-percentage=100
force-gamemode=false
function-permission-level=2
gamemode=survival
generate-structures=true
generator-settings={}
hardcore=false
hide-online-players=false
initial-disabled-packs=
initial-enabled-packs=vanilla
level-name=world
level-seed=
level-type=minecraft\:normal
log-ips=true
max-chained-neighbor-updates=1000000
max-players=2
max-tick-time=60000
max-world-size=29999984
motd=Fischers Fritz fischt frische Fische
network-compression-threshold=256
online-mode=true
op-permission-level=4
pause-when-empty-seconds=60
player-idle-timeout=0
prevent-proxy-connections=false
pvp=true
query.port=25565
rate-limit=0
rcon.password=password
rcon.port=25575
region-file-compression=deflate
require-resource-pack=false
resource-pack=
resource-pack-id=
resource-pack-prompt=
resource-pack-sha1=
server-ip=
server-port=25565
simulation-distance=10
spawn-monsters=true
spawn-protection=16
sync-chunk-writes=false
text-filtering-config=
text-filtering-version=0
use-native-transport=true
view-distance=10
white-list=true
```

My whitelist.json:
```json
[
  {
    "uuid": "f1a2157e-a000-2ad8-9b96-c220c7019385",
    "name": "example"
  }
]
```

My ops.json:
```json
[
  {
    "uuid": "7f67870f-57d6-406c-9ad1-e848715a9453",
    "name": "technat",
    "level": 4,
    "bypassesPlayerLimit": true
  }
]
```

And my server is started through systemd:
```
[Unit]
Description=Minecraft Server
After=network.target

[Install]
WantedBy=multi-user.target

[Service]
User=minecraft
Group=minecraft
Nice=1
SuccessExitStatus=0 1
PrivateDevices=true
NoNewPrivileges=true
WorkingDirectory=/opt/minecraft
# https://docs.papermc.io/misc/tools/start-script-gen
ExecStart=java -Xms2867M -Xmx2867M -XX:+AlwaysPreTouch -XX:+DisableExplicitGC -XX:+ParallelRefProcEnabled -XX:+PerfDisableSharedMem -XX:+UnlockExperimentalVMOptions -XX:+UseG1GC -XX:G1HeapRegionSize=8M -XX:G1HeapWastePercent=5 -XX:G1MaxNewSizePercent=40 -XX:G1MixedGCCountTarget=4 -XX:G1MixedGCLiveThresholdPercent=90 -XX:G1NewSizePercent=30 -XX:G1RSetUpdatingPauseTimePercent=5 -XX:G1ReservePercent=20 -XX:InitiatingHeapOccupancyPercent=15 -XX:MaxGCPauseMillis=200 -XX:MaxTenuringThreshold=1 -XX:SurvivorRatio=32 -Dusing.aikars.flags=https://mcflags.emc.gs -Daikars.new.flags=true -jar server.jar nogui 
ExecStop=/usr/local/bin/mcrcon -H 127.0.0.1 -P 25575 -p password stop
ExecReload=/usr/local/bin/mcrcon -H 127.0.0.1 -P 25575 -p password reload
```

For easier maintenance a bashrc alias helps sending commands to the server:

alias mc=`mcrcon -H 127.0.0.1 -P 25575 -p password`

