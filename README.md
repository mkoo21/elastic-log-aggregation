
## Troubleshooting
- Elasticsearch requires that you set `vm.max_map_count=262144` in your `/etc/sysctl.conf`. Config ad hoc with `sysctl -w vm.max_map_count=262144`, or manually reload the config (will log each setting/file applied) with `sysctl --system`.
If you do not the container will exit with status 78 after failing bootstrap checks.
