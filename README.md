## Overview
Here we have dockerized nginx-server running with Elasticsearch and Kibana also run on the same host (ES and Kibana don't actually need to be on the same host). 
We also have filebeat and metricbeat containers to pull data from nginx (filebeat needs to share a volume with nginx and so either needs to run on the same host or needs to be installed in the same server/container as nginx).

Start by spinning up the stack with `docker-compose up --build -d`, then create some access logs with `curl localhost:8080`. The whole stack might take a few minutes to complete setup and be ready for action; first ES/Kibana will take a moment to be ready for service, then filebeat/metricbeat will run setup to create some ES index patterns and dashboards. 

## Elasticsearch
- List indices: `curl $ES_HOSTNAME:9200/_cat/indices`
- Get mappings: `curl $ES_HOSTNAME:9200/$INDEX/_mapping` (You can use the mapping API to make changes like field type)
- Query index: `curl $ES_HOSTNAME:9200/$INDEX/_search?q=user=kimchy`
- For this setup docker-compose will make the container's port 9200 to to host's port 9200, so `localhost` will work for `ES_HOSTNAME`

## Kibana
- As usual, access the GUI via `localhost:5601`
- You need to create an index pattern (e.g. `filebeat*`) under 'manage' or 'connect to elasticsearch' on the main page and set the @timestamp field
- Then you can query data under discover by using the search and time filters.

## Troubleshooting

### Elasticsearch
- Elasticsearch requires that you set `vm.max_map_count=262144` in your `/etc/sysctl.conf`. Config ad hoc with `sysctl -w vm.max_map_count=262144`, or manually reload the config (will log each setting/file applied, which is helpful) with `sysctl --system`.
- You can run a single node cluster for dev with `discovery.type=single-node` (don't do this in prod)
- For production config, you need to run at least 1 master node and data node, configure them, and set some environment variables. See https://www.elastic.co/guide/en/elasticsearch/reference/7.0/important-settings.html and pay attention to the heap size setting (don't go above the JVM limit for compressed pointers)
- If you don't configure environment properly the container will exit with status 78.


## TODO
logrotate

## Prod usage
- You definitely need to have at least 1 additional data shard. Don't run ES in single-node mode.
- You definitely want to buy the security X-Pack before exposing data to the internet.
