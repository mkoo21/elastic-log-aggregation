FROM docker.elastic.co/beats/metricbeat:7.0.0

RUN metricbeat modules enable nginx
COPY nginx.yml /usr/share/metricbeat/modules.d/nginx.yml
COPY ./http_probe /usr/share/metricbeat
COPY ./setup /usr/share/metricbeat

# Run setup script
ENTRYPOINT /usr/share/metricbeat/setup
