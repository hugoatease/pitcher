FROM solr:7.1

USER root

RUN apt-get update && apt-get install -y sudo
COPY solr /etc/sudoers.d/solr

RUN mkdir -p /opt/solr/server/solr/musicbrainz
RUN chown -Rf solr /opt/solr/server/solr/musicbrainz
COPY postgresql-42.1.4.jar /opt/solr/dist/postgresql-42.1.4.jar

USER solr

VOLUME /opt/solr/server/solr/musicbrainz/data
