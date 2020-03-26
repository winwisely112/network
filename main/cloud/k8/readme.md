# k8

Use knaitve and istio

https://cloud.google.com/run#section-2

can ALSO use with gloo maybe for easy bebugging

# Tiers

knative is only serverless for stateless systems and all our microservices are
- SO the SAAS layer runs on Knaitve
- The PAAS layer is NATS, MINIO, etc and runs all the time.

cloudrun
