MAKEPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAMESPACE:=logging-example
NAME:=$(NAMESPACE)

install-example: uninstall-example
	helm install -n $(NAMESPACE) --create-namespace $(NAME) $(MAKEPATH)/charts/rancher-logging-example

install-example-no-limits-or-requests: uninstall-example
	helm install -n $(NAMESPACE) --create-namespace $(NAME) $(MAKEPATH)/charts/rancher-logging-example \
		--set elasticsearch.resources.requests.cpu="0" \
		--set elasticsearch.resources.requests.memory="0" \
		--set elasticsearch.resources.limits.cpu="0" \
		--set elasticsearch.resources.limits.memory="0"

uninstall-example:
	-helm uninstall -n $(NAMESPACE) $(NAME)
	-kubectl delete namespace $(NAMESPACE)

port-forward:
	kubectl port-forward -n $(NAMESPACE) svc/elasticsearch-master 9200:9200

open-elasticsearch:
	open http://localhost:9200/_cat/indices
