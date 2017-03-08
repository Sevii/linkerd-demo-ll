#Linkerd-demo-ll

Linkerd-demo-ll is a demo comparing kube-dns and linkerd. Linkderd (linkerd.io) is a proxy which adds service discovery, routing and failure handling to services. It integrates with kubernetes for service discovery. 


##Microservices

The microservices directory contains the microservices used in this demo.

* Edge: ui templates, takes public traffic
* Login: verifies user logins
* Reports: generates reports
* Weather: produces weather data
* Stock: produces stocks data


![ScreenShot](https://github.com/Sevii/linkerd-demo-ll/blob/master/Graph.png)



## kube-dns-test

The kube-dns-test folder contains the kubernetes manifests to setup a functional version of the system using kube-dns for discovery. It contains:

###Contents
Prometheus
Microservice service definitions

## linkerd-test

The linkerd-test folder contains the kubernetes manifests to setup a functional version of the system using linkerd for discovery. Do not use this config inside minikube[1].

###Contents
Prometheus
Microservice service definitions
Linkerd daemonset


## minikube-linkerd-test

The minikube-linkerd-test contains the kubernetes manifests to setup a functional version of the system using linkerd for discovery inside of minikube. 

###Contents

Prometheus
Microservice service definitions
Linkerd daemonset





References
[1] https://github.com/kubernetes/minikube/issues/757
