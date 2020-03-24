# Main PASS

Helm of : MINIO, NATS, REDIS (TCPIP), TSDB

Joe Webb, [24. Mar 2020 at 08:54:14]:
...The Paas Layer ( minio, etc) MUST be made in network repo going forward.

1. Make file to kick it off

2. CI not needed because its "for ever".

3. SAAS layer ( each MS like chat, etc) is still in CI of course and runs on a different kluster. This cleanly seperates stateful from non stateful layrs.

4. Make file must also have the benchmark stuff for the PAAS

5. Keep the OLD PAAS makes files. So suggest you just number them using folders above each make file ( 01, 02, )

- add a readme for each that lists the stack

6. Make file to open the CLI and Web tools for accessing the kluster management tools and metrics, logging tools !

ALL this is so that we are REPEATABLE and DECOUPLED properly so anyone can bring up a cluster !!!! 