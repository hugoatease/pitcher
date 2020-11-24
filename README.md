# Pitcher 🎶
Track search and music metadata retrieval service, based on [Musicbrainz](http://musicbrainz.org/) data.

The main project's goal is to be able to match a simple `Artist - Track` string combination with full metadata retrieved from open databases, in a reasonably fast manner.

This repository includes :
- The Pitcher microservice, which offers a simple API around music metadata search and retrieval.
- A _musicbrainz_ Helm chart to set up database import and replication Kubernetes jobs.
- A [SolrCloud](https://lucene.apache.org/solr/) core configuration with the search schema for track matching with Musicbrainz IDs.
- [Strimzi](https://strimzi.io/) Kubernetes Custom Resources for the Kafka Connect setup allowing the Musicbrainz database and Solr track matching collection to stay in sync.