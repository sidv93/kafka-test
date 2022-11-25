# Kafka Setup

This repository is a simple reproduction of a event broker setup using Kafka. We monitor a system's key metrics and stream them using Kafka.

The following are the services in this repo

Service		| Description
------------|------------
system-stats 		| Gathers key system metrics
producer 		    | Gets metics as a json from _system-stats_ and pushes it into Kafka
kafka-broker		| Kafka broker
consumer 		    | Consumes data from kafka and streams it to UI
ui		            | A simple web UI which renders system metrics

## Setup instructions (_for deployment_)

The following needs to be present in a _host machine_ chosen for deployment.

| | | |
|-|-|-|
| Docker | |
| Docker compose | |

## Deployment steps
* Create a `.env` by adjusting the `.env.sample`
* `cd` to the root folder where this repo is cloned
* Execute the following commands in sequence
```
	docker-compose build
    docker-compose up -d
```

This command would bring up all containers in detached mode. Navigate to [http://localhost:6000](http://localhost:6000) and the UI should render system stats from the host machine.


### Improvements
* Dynamic way of creating Kafka topics
* Multiple Kafka brokers with partitions
* As of now, only one socket can connect consumer service. This has to be enhanced to support a connection pool, which should broadcast stats to all subsribers
* Other key metrics like GPU stats

