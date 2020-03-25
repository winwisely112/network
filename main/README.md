# Main PASS

Helm of : MINIO, NATS, REDIS (TCPIP), TSDB


1. Make file to kick it off.

2. CI not needed because its "for ever".

- a dev ops person manages updates manually for now.

3. SAAS layer ( packages repo) is still in CI of course and runs on a different Cluster. This cleanly separates stateful from stateless layers.

4. Make file must also have the benchmark stuff for the PAAS

- For each store
- Can be simple for now. Just need to know each works and its limitations.

5. Keep the OLD PAAS makes files. 

- This is so that when we change how we run something we can go back to it later.

- So suggest you just number them using folders above each make file ( 01, 02,

- add a readme for each that lists the stack and the reason for it being the way it is.

6. Make file to open the CLI and Web tools for accessing the Cluster management tools and metrics, logging tools !

ALL this is so that we are REPEATABLE and DECOUPLED properly so anyone can bring up a cluster !!!! 

## IAAS

https://github.com/weaveworks/ignite

ARM:
https://github.com/weaveworks/ignite/search?q=arm&unscoped_q=arm

https://github.com/weaveworks/ignite/blob/master/docs/installation.md

https://github.com/weaveworks/ignite/blob/master/docs/cloudprovider.md

https://github.com/weaveworks/wks-quickstart-firekube

https://github.com/inlets/inlets#stretch-goals

https://github.com/weaveworks/wks-quickstart-firekube

CoreOS Fork:
https://www.flatcar-linux.org/releases/

Packet Roadmap:
https://www.packet.com/developers/roadmap/

Packet Github:
https://github.com/packethost