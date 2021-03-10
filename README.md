# Rancher Logging Examples

This repository contains an example chart for testing your logging setup.

## Charts

### Rancher Logging Example
The `rancher-logging-example` is an end to end setup of an application, multiple configurations for log collection, and a simple output. This chart is recommended for running an example payload for the [rancher-logging](https://rancher.com/docs/rancher/v2.x/en/logging/v2.5/) chart.  

This chart demonstrates 3 scenarios:

1. `all-logs` collects all logs in the cluster
2. `logs-by-label` collects all logs for pods with the label `rancher.logging.example/app: log-generator`
3. `namespace-logs` collects all logs for the namespace this chart is deployed in.

#### Installation
To install this chart run:

```shell
git clone https://github.com/paynejacob/rancher-logging-examples
cd rancher-logging-examples
make install
```


#### Checking the output
 
To see information about the logs being collected you will need a proxy to the log output service.  The helm notes will include instructions for setting this up after a successful installation. Then, you can view your logs by going to [http://localhost:8080](http://localhost:8080).  It can take a few seconds for logs to show up.

## Log Generator

`docker run paynejacob/log-generator:latest`

The log generator is a simple application that will constantly output logs.  

You can tweak the log output by setting the following environment variables:

- `LOGFORMAT` a go template string for the generated logs
- `FREQUENCY` the maximum delay between log messages
- `CONSTANTRATE` by default the log generator will output messages every `FREQUENCY` seconds.  If this variable is set to true messages will be logged every `FREQUENCY` seconds.

## Log Output

`docker run -p 80:80 paynejacob/log-output:latest`

The log output is a simple http service that can receive logs and be used to test a logging setup. A `GET /` will return a summary of all logs received by the service. Logs are organized into indices.  Think of an index as a folder for logs. The logs are written by preforming `POST /<index name>/` with one or more log messages.  A `GET /<index name>/` will return the last 100 logs received by the service for the given index. **Logs are stored in memory and all statistics and logs will be lost if the pod is restarted.**
