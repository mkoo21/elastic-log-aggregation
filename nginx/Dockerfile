FROM nginx

# Install prereqs
RUN apt-get update \
  && apt-get install -y vim curl

# Remove default log files, which are softlinked to stdout/stderr
RUN rm /var/log/nginx/*.log

# Upload nginx conf

# Setup filebeat
#RUN curl -L -O https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-7.0.0-amd64.deb \
#  && dpkg -i filebeat-7.0.0-amd64.deb \
#  && filebeat modules enable nginx

#COPY filebeat.yml /etc/filebeat/filebeat.yml
#COPY filebeat.nginx.yml /etc/filebeat/modules.d/nginx.yml

