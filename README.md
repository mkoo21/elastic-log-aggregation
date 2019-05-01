## Elasticsearch
- List indices: `curl $ES_HOST:9200/_cat/indices`
- Get mappings: `curl $ES_HOST:9200/$INDEX/_mapping` (You can use the mapping API to make changes like field type)
- Query index: `curl $ES_HOST:9200/$INDEX/_search?q=user=kimchy`

## Kibana
- You need to create an index pattern (e.g. `filebeat*`) under 'manage' or 'connect to elasticsearch' on the main page and set the @timestamp field
- Then you can query data under discover by using the search and time filters.

## Troubleshooting

### Elasticsearch
- Elasticsearch requires that you set `vm.max_map_count=262144` in your `/etc/sysctl.conf`. Config ad hoc with `sysctl -w vm.max_map_count=262144`, or manually reload the config (will log each setting/file applied) with `sysctl --system`.
- You can run a single node cluster for dev with `discovery.type=single-node` (don't do this in prod)
- For production config, you need to run at least 1 master node and data node, configure them, and set some environment variables. See https://www.elastic.co/guide/en/elasticsearch/reference/7.0/important-settings.html and pay attention to the heap size setting (don't go above the JVM limit for compressed pointers)
- If you do configure properly the container will exit with status 78.


## TODO
logrotate
