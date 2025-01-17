# Kafka Ecosystem
## Kafka Connect 
### Source Connectors
To get data from Common Data Sources
### Sink Connectors
To publish data in Common Data Stores \
\
Easy way to make Kafka part of ETL pipeline. Scaling made easy. Leverage code that acheives fault tolerance, idempotence and
scalability. On ConfluentHub.

## Kafka Streams
Easy dataprocessing and transformation library, and written as standard application. Highly scalable, and one record at a
time processing. Elastic and fault tolerant. Used for Transformations, Enrichment, Fraud detection and Monitoring. Distributed co-ordination
and operational simplicity leveraged.
### Topology

## Misc.
### Need for Schema Registry
Kafka works on bytes -- no data verification. Deal with the case: producer sends bad data, field gets renamed, data changes format from one day to another. Deserializer fails: this introduces the need for a schema registry. Kafka brokers shouldn't do this as it undermines one of Kafka's fundamental benefits: a broker shouldn't read data, never loads anything into memory [zero memory]: broker sees 1s and 0s. \
With registry:
- Store and retrieve schemas for Producers / Consumers
- Enforce Backward / Forward / Full compatibility on topics
- Decrease the size of the payload of data \
Producer sends schema to schema registry. 
