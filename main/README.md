# Main PASS


SAAS
- k8. But does not have all of the Paasy stuff we need.

- ion (MS). new build 1 day away
 - yedis/redis 
 - nats 

PAAS
- Helm of : MINIO, NATS, REDIS (TCPIP), TSDB
- no CI, all from make files

- security
 - opa

IAAS
- something to run on baremtal, etc etc 
- GCE.

---

## Sprint
- PASS (nats, redis, minio )
	- GCN Helm (same as in packages, with without the SAAS) 
	- LATER: kustom (KOTS = https://github.com/replicatedhq/kots)
- ion SAAS
	- docker and k8 deployed into own Cluster and calling our PAAS.
- TEST and PRAY
  - Give them access to our makefiles ( they try out our PASS AND SAAS)
- sec
 - opa (https://github.com/open-policy-agent)
	- stage 1: username and password
	- stage later: rings security policy. you auth differently depending on the org you want access to. Ring level 1 = super hard.. Ring level 5 = easy entry.
