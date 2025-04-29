# Ayanami-Fan-Controller
My fan speed controller for my **Rei-chan** server — my Dell PowerEdge R230 server running TrueNAS Scale.

# Background
So I got my new addition to my HomeLab, a Dell PowerEdge R230 server for a while now (I impulsively bought it). I have finally decided to put it to use as my NAS server due to its 4-bay hotswap drive bays and I also have 4 1TB 2.5" SAS drives lying around (I also impulsively bought them 😅). The server is amazing the Xeon E3 literally sip just couple watt of power running at IDLE with also a iDRAC8 IPMI it is amazing. The only issued I have with it is the crazy fan noise to place it in my office to something like a synology NAS. It's literally always quiet and methodical compared to my other server hence why I name it Rei-chan from Evangelion - amazing server, just always calm and silent. After a little bit of digging around I found out that on a Linux system there is a program called `ipmitool` which allow us to manage and configure ipmi interface like HPE iLO or Dell iDRAC but the process of configuring a Fan Speed isn't really practical it's involved typing quite a bit of command on my TrueNAS Scale shell so I decided to create this project to enable me to easily adjust my fan speed using TrueNAS Docker Apps.

# Requirements
- Computer or Server with an IPMI interface **THIS DO NOT WORK ON NORMAL PC** I only tested this on my Dell PowerEdge R230 with iDRAC8 Enterprise
- Linux OS (If you are using TrueNAS like me be sure to use TrueNAS Scale which is linux based not TrueNAS Core which is FreeBSD based)
- Docker installed
- ipmitool installed (`apt install ipmitool` on Debian/Ubuntu or `yum install ipmitool` on CentOS/RHEL)

# Features
- Beautiful UI inspired from Rei Ayanami's EVA-00 from the anime Neon Genesis Evangelion
- Super easy to setup
- REST API to control the fan speed or just give control back to the iPMI
- Extremely lightweight (thanks to Golang)

<figure>
    <img src="https://raw.githubusercontent.com/wasinuddy/ayanami-fan-controller/main/images/screenshot.png" alt="Ayanami Fan Controller" width="600"/>
    <figcaption>Screenshot of the WebUI</figcaption>
</figure>

# Installation
The installation is very straightforward mount the `/dev/ipmi0` device to the container and run the container. You can use the following command to run the container:
```bash
docker run -d \
    --name ayanami-fan-controller \
    --restart unless-stopped \
    -p 8080:8080 \
    -v /dev/ipmi0:/dev/ipmi0 \
    ghcr.io/wasinuddy/ayanami-fan-controller:latest
```
On TrueNAS Scale just create a new app according to the docker run command above make sure the apps is **privilaged** if you are unsure of the ipmi device you can just mount the whole `/dev` directory to the container but this is not recommended for security reason.

# Usage
After the container is running you can access the web UI on `http://<your-server-ip>:8080` and you will see a simple UI with a slider to control the fan speed if you use manual mode or you can just click the **Auto** button to let the iDRAC control the fan speed. If you want Rei-chan to operate with precision, you can click on **Silent** Mode which will set the fan speed to minimal levels. Turning the fan off completely can also be done but I do not recommend it.

**Note: I created this project for my own personal server. Use it at your own risk. I am not responsible if your server transforms into an Evangelion**