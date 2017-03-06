#!/bin/bash

edgeUrl=$(eval minikube ip)
edgePort=$(kubectl get service edge -o=json | jq '.spec.ports[0].nodePort')
printf "%s\n" "$edgeUrl"

for ((i=1;i<=20;i++)); do 
	http -a Admin:admin "${edgeUrl}:${edgePort}";
	http -a Guest:pass23word "${edgeUrl}:${edgePort}";
	http -a Ad2min:admin "${edgeUrl}:${edgePort}";
	http -a Guest:password "${edgeUrl}:${edgePort}";

	http -a Admin:ad23in "${edgeUrl}:${edgePort}";
	http -a Stern:mrp23aulstern "${edgeUrl}:${edgePort}";
	http -a Guefst:password "${edgeUrl}:${edgePort}";

	http -a Stern:mrpaulstern "${edgeUrl}:${edgePort}";

	http -a Stsern:mrpaulstern "${edgeUrl}:${edgePort}";
	sleep 1;
 done
