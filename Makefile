MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAMESPACE:=logging-example
NAME:=$(NAMESPACE)

install-example: uninstall-example
	helm install -n $(NAMESPACE) --create-namespace $(NAME) $(MAKEPATH)/charts/rancher-logging-example

uninstall-example:
	-helm uninstall -n $(NAMESPACE) $(NAME)
	-kubectl delete namespace $(NAMESPACE)
