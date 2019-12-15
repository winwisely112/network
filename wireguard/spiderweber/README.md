# spiderweb

nest = docs

- TODO: Make it use hugo and have a Web Server host it, so it can be docker deployed.

---

wasp = admin cli

- uses GRPC, so we can then make a GUI Management API using Flutter
- In order to support Flutter Web, we probably have to start to use the GRPC gateway.

---


mosquito = wireguard client ( desktop)

- Make a Flutter GUI

---

knot = wireguard server

- TODO: currently needs wireguard server to be present. SO run in docker, so adapt make file to use docker and expose ports, etc

---

spider = wiregiard server ops.

- connects to ETCD to get the "knots - wg-servers"


# TODO

- A docker compose to we can all run the bits locally.

- A K8 to run this way. ETDC can then scale globally.

