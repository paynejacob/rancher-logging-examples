# Rancher Logging Examples

This repository contains 2 example charts for testing your logging setup and a simple application for generating logs.

## Charts

### Simple Application
The `simple-application` chart will deploy 2 pods that will constantly out logs.  This chart is recommended for when you want to configure your own logging setup and need a simple application to collect logs from.

### Rancher Logging Example
The `rancher-logging-example` is an end to end setup of an application, multiple configurations for log collection, and an elasticsearch cluster as an output.  This chart is recommended if you want to see a working example of the [rancher-logging](https://rancher.com/docs/rancher/v2.x/en/logging/v2.5/) chart.  

This chart demonstrates 3 scenarios:

1. `all-logs` collects all logs in the cluster
2. `logs-by-label` collects all logs for pods with the label `rancher.logging.example/app: log-generator`
3. `namespace-logs` collects all logs for the namespace this chart is deployed in.

#### Resource Requests and Limits

You can change `elasticsearch` resource requests and limits from the upstream chart by editing `values.yaml` as follows...

```yaml
elasticsearch:
  resources:
    requests:
      cpu: "0"
      memory: "0"
    limits:
      cpu: "0"
      memory: "0"
```

#### Checking ElasticSearch

You can if `rancher-logging` is working by forwarding port `9200` and accessing the instance via `localhost`.

```bash
kubectl port-forward -n logging-example svc/elasticsearch-master 9200:9200
```

Now, navigate to your browser and access `http://localhost:9200/_cat/indices`.
You should see three indices from the Flow(s), Output(s), ClusterFlow(s), and ClusterOuput(s) created.

## Log Generator

`docker run paynejacob/log-generator:latest`

The log generator is a simple application that will constantly output logs.  

You can tweak the log output by setting the following environment variables:

- `LOGFORMAT` a go template string for the generated logs
- `FREQUENCY` the maximum delay between log messages
- `CONSTANTRATE` by default the log generator will output messages ever 0 - `FREQUENCY` seconds.  If this variable is set to true messages will be logged ever `FREQUENCY` seconds.
