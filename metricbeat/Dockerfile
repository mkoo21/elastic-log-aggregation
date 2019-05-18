FROM docker.elastic.co/beats/metricbeat:7.0.0

RUN metricbeat modules enable nginx
COPY nginx.yml /usr/share/metricbeat/modules.d/nginx.yml
