## Overview
Here we have an nginx server in a docker container that shares a volume with a filebeat container on the same host. Elasticsearch and Kibana also run on the same host in this case, although there is no need for this since there's no need for them to share a volume with any other components of this system. 

Start by spinning up the stack with `docker-compose up --build -d`, then create some access logs with `curl localhost:8080`.

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
