
## NKL and MultiCluster Load Balancing with HTTP/S

<br/>

## Overview

<br/>

>With the NGINX Plus Servers located external to the Cluster, using NGINX's advanced HTTP/S features provide Enterprise class traffic management solutions.
  
- MultiCluster Active/Active Load Balancing
- Horizontal Cluster Scaling
- HTTP Split Clients - for `A/B, Blue/Green, and Canary` test and production traffic steering.  Allows Cluster operations/maintainence like:
  - Node upgrades
  - Software upgrades/security patches
  - Cluster resource expansions
  - Troubleshooting
  - ^^ With NO downtime or reloads
- NGINX Zone Sync of KeyVal data
- API Gateway functions
- Advanced TLS Processing - MutualTLS, OCSP, FIPS, dynamic cert loading
- Advanced Security features - App Protect WAF Firewall, Oauth, JWT, Dynamic Rate and Bandwidth limits, GeoIP, IP block/allow lists
- NGINX Java Script (NJS) for custom solutions

<br/>

## Reference Diagram for NKL HTTP MultiCluster Load Balancing Solution

<br/>

Multiple K8s Clusters, HA NGINX Plus LB Servers, NKL Controllers

![NKL MultiCluster Diagram](../media/nkl-multicluster-config.png)


<br/>

NKL Watching nginx-ingress Service and Updating HTTP Upstreams; Service Type Loadbalancer or NodePort:

![NKL MultiCluster LoadBalancer](../media/nkl-cluster1-add-loadbalancer.png)
or
![NKL MultiCluster NodePort](../media/nkl-cluster1-add-nodeport.png)

<br/>

MultiCluster Load Balancing

![NKL MultiCluster Dashboard](../media/nkl-multicluster-upstreams.png)

<br/>

NGINX HTTP Split Clients with Dynamic Ratio -- 10% Cluster1 : 90% Cluster2 

![NGINX HTTP Split 10](../media/nkl-clusters-10.png)


<br/>

The `Installation Guide` for HTTP MultiCluster Solution is located in the docs/http folder:

[HTTP MultiCluster Loadbalancing Guide](../http/http-installation-guide.md)

<br/>

## Authors
- Chris Akker - Solutions Architect - Community and Alliances @ F5, Inc.
- Steve Wagner - Solutions Architect - Community and Alliances @ F5, Inc.