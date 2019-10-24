
# wireguard

Why ?

Forces mobiles to not connect unless vpn on.
Does not leak DNS or IP address.
Well checked the CVE's and reputable.

## docs
https://www.reddit.com/r/WireGuard/


https://www.wireguard.com/install/
- every desktop and mobile is there.


## libs

RUST
https://github.com/cloudflare/boringtun
- FFI bindings making it easy to use with Flutter FFI ?


Wireguard
- https://github.com/WireGuard

Docker
https://github.com/masipcat/wireguard-go-docker/issues/1
https://github.com/masipcat/wireguard-go-docker


Linuxkit has wireguard support and does multiarch builds to anything.
Infrakit is k8.
https://github.com/linuxkit/linuxkit/blob/master/docs/wireguard.md


Mobile build tools
https://github.com/AOSiP/platform_build_soong
-its uses to build open source Android.



MACOS and IOS
- Mobiles using gomobile
- supports embedding
	- using  -buildmode c-archive for golang compilation. (https://github.com/WireGuard/wireguard-apple/blob/master/wireguard-go-bridge/Makefile)
https://github.com/aequitas/macos-menubar-wireguard
- nice

Android
https://github.com/msfjarvis/viscerion
- uses golang under the hood: https://github.com/msfjarvis/viscerion/tree/master/native/tools
- can wrap with flutter then.
- on play store: https://play.google.com/store/apps/details?id=me.msfjarvis.viscerion
	- it can exclude other apps from the tunnel. SO the other apps go through the normal connection.
	- allows QR code to be used to transfer tunnel :)
- Telegram group: https://t.me/joinchat/DKHRHEvMR2kz7IW3LB46dQ
- His recommended Server: https://github.com/trailofbits/algo


Windows
- https://github.com/WireGuard/wireguard-windows
- all goalng :)
- supports embedding into another app: https://github.com/WireGuard/wireguard-windows/tree/master/embeddable-dll-service


- Is has a Roaming Server to allow Clients to find each other.
- https://github.com/mdlayher/wgipam
	- https://github.com/mdlayher/wgdynamic-go
	- this is a server and a client. The idea is to allow the IP addresses of all clients to be exchanged via the Server.
	- DOes it support NAted client ?https://github.com/mdlayher/wgdynamic-go/issues

- Easy to run the server with docker and k8
	- https://github.com/masipcat/wireguard-go-docker

Rest API over Wireguard Server
- https://github.com/gun1x/wireguard_rest_api
- This would allow remote management aspects to be automated.
- ACTIVE

- golang.zx2c4.com/wireguard/wgctrl
- core one

Configuration
https://github.com/pirate/wireguard-docs
- good examples
	- https://github.com/pirate/wireguard-docs/tree/master/example-internet-browsing-vpn
		- Server, laptop and mobile. Exactly what i want !!

GUI
https://github.com/subspacecloud/subspace
- golang based web gui for a wireguard server
- creates the WG keys n install: https://github.com/subspacecloud/subspace/blob/master/entrypoint.sh
- Users: https://github.com/soundscapecloud/soundscape
	- cool :)
- Run it with Caddy and simple docker:
	- https://github.com/subspacecloud/subspace/issues/31#issuecomment-510152383

https://github.com/usrpro/wg-share
- manage peers

Metrics
https://github.com/mdlayher/wireguard_exporter
- export wireguard metrics over prom
- export machine metrics in general
	- https://github.com/prometheus/node_exporter
	- https://github.com/martinlindhe/wmi_exporter
		- has an MSI, setups a service and makes a firewall exception



Dynamic peers
https://github.com/mdlayher/wgdynamic-go

https://github.com/mdlayher/wgipam



