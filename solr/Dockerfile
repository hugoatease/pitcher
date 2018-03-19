FROM solr:7.1

USER root

RUN apt-get update && apt-get install -y sudo
COPY solr /etc/sudoers.d/solr

RUN mkdir /data
RUN chown -Rf solr /data
COPY postgresql-42.1.4.jar /opt/solr/dist/postgresql-42.1.4.jar

USER solr

VOLUME /data
