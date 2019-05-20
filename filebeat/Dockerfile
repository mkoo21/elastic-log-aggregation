FROM docker.elastic.co/beats/filebeat:7.0.0

RUN filebeat modules enable nginx
COPY ./http_probe /usr/share/filebeat
COPY ./setup /usr/share/filebeat

ENTRYPOINT /usr/share/filebeat/setup
