MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAMESPACE:=logging-example
NAME:=$(NAMESPACE)

install-example: uninstall-example
	helm install -n $(NAMESPACE) --create-namespace $(NAME) $(MAKEPATH)/charts/rancher-logging-example

uninstall-example:
	-helm uninstall -n $(NAMESPACE) $(NAME)
	-kubectl delete namespace $(NAMESPACE)

port-forward:
	kubectl port-forward -n $(NAMESPACE) svc/elasticsearch-master 9200:9200

open-elasticsearch:
	open http://localhost:9200/_cat/indices
