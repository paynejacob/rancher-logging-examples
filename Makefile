UNAME:=$(shell uname -s)
MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAMESPACE:=rancher-logging-example
NAME:=$(NAMESPACE)

install: uninstall
	helm install -n $(NAMESPACE) --create-namespace $(NAME) $(MAKEPATH)/charts/rancher-logging-example

uninstall:
	-helm uninstall -n $(NAMESPACE) $(NAME)
	-kubectl delete namespace $(NAMESPACE)

port-forward:
	kubectl port-forward -n $(NAMESPACE) svc/$(NAMESPACE)-log-output 8080:80

open-log-output:
ifeq ($(UNAME), Darwin)
	open http://localhost:8080
else ifeq ($(UNAME), Linux)
	xdg-open http://localhost:8080
else
	explorer http://localhost:8080
endif
