# Asuka-Fan-Controller
My TrueNAS Scale Dell PowerEdge R230 (Asuka chan) Fan Controller powered by ipmitool

# Disclaimer
As this container used ipmitool it is meant to be used on Linux systems with an IPMI interface(In my case DELL IDRAC).

# Description
I have recently set up a TrueNAS Scale server with a Dell PowerEdge R230 which I named it Asuka-chan from Neon Genesis Evangelion. I wanted to control the fan speed of the server using ipmitool, but I couldn't find a good solution. So I created this container to do it. This container uses ipmitool to control the fan speed of the server. It also has a web interface to control the fan speed.

# Features
- Control fan speed using ipmitool
- Web interface to control fan speed (Evangelion EVA-02 theme)

# Basic Usage
mount /dev/ipmi0 of the container to /dev/ipmi0 of the host to give access to the ipmi device

# TODO
- Real time temperature monitoring
- Temperature thresholds to roll back to automatic fan speed
- EVA-01 theme if I am free
- Better documentation
