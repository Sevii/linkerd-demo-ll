# Linkerd-demo-ll

Linkerd-demo-ll is a demo comparing kube-dns and linkerd with a slightly complicated microservice graph. Linkderd (http://linkerd.io) is a proxy which adds service discovery, routing and failure handling to services. It integrates with kubernetes for service discovery. 


## Microservices

The microservices directory contains the microservices used in this demo.

* Edge: ui templates, takes public traffic
* Login: verifies user logins
* Reports: generates reports
* Weather: produces weather data
* Stock: produces stocks data


![ScreenShot](https://github.com/Sevii/linkerd-demo-ll/blob/master/Graph.png)



## kube-dns-test

The kube-dns-test folder contains the kubernetes manifests to setup a functional version of the system using kube-dns for discovery.

### Contents
* Prometheus
* Microservice service definitions

## linkerd-test

The linkerd-test folder contains the kubernetes manifests to setup a functional version of the system using linkerd for discovery. Do not use this config inside minikube[1].

### Contents
* Prometheus
* Microservice service definitions
* Linkerd daemonset


## minikube-linkerd-test

The minikube-linkerd-test contains the kubernetes manifests to setup a functional version of the system using linkerd for discovery inside of minikube. 

### Contents

* Prometheus
* Microservice service definitions
* Linkerd daemonset

## Prometheus
Prometheus is setup so that we have a stable platform to compare metrics on. Linkerd provides its own metrics collection and dashboard, but it is obviously not avaliable for kube-dns. Prometheus is configured to perform scraping jobs against the kubernetes services, pulling both node metrics and pod intrementation metrics. The microservice http endpoints are instrumented using the prometheus go libraries [2]. 

### Instrumenting
http://the-hobbes.github.io/update/prometheus/metrics/instrumentation/monitoring/2016/03/27/instrumenting-with-prometheus.html

### Configuring prometheus inside kubernetes
https://coreos.com/blog/monitoring-kubernetes-with-prometheus.html

## Container Hosting
The docker containers used for this demo are hosted at http://quay.io/sevii. They can be built from source using the build.sh file in each microservice folder. This file ensure that the go compiler target linux when it creates the binary. The binary is then added to a docker container. 

## References
* [1] https://github.com/kubernetes/minikube/issues/757
* [2] https://github.com/prometheus/client_golang/tree/master/prometheus

