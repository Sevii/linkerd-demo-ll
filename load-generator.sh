#!/bin/bash

edgeUrl=$(eval minikube ip)
edgePort=$(kubectl get service edge -o=json | jq '.spec.ports[0].nodePort')
printf "%s\n" "$edgeUrl"

for ((i=1;i<=20;i++)); do 
	http -a Admin:admin "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Guest:pass23word "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Ad2min:admin "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Guest:password "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Admin:ad23in "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Stern:mrp23aulstern "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Guefst:password "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Stern:mrpaulstern "${edgeUrl}:${edgePort}";
	sleep 1;
	http -a Stsern:mrpaulstern "${edgeUrl}:${edgePort}";
	sleep 1;
 done
