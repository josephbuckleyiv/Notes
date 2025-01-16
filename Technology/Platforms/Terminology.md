# Terminology
## Broker
A server within a Kafka Cluster that receives messages from producers, stores them in partitions, and serves those messages to consumers.
In diagrams, this 'houses' the topics/partitions from which consumers pull.
## Consumers
Read data from the topic, via 'pull' model. They know which broker to read from.
## Topics
A stream of data. Can have as many as you want, and is identifiable via Name. Similar to a table in a database, but without 
constaints. The sequence of data passing to Topic is called a data stream. Cannot query topics. Immutable -- once written, cannot be 
changed. Data is kept for a limited time.
## Partitions
Topics are split in partitions. Messages within each partition are ordered.
##Producers
Writes to topics. Producers know to which partition to write to. In case of broker failures, producers will recover. Can
send key with the message. Key=null means data is sent round robin. Key=null, then all messages for that key will always go to
the same partition.
## Kafka Message
Key-binary, value-binary, Compression Type, Headers, Partition + Offset, Timestamp \
Message Serialization is used because Kafka only accepts bytes. Hash assigns key to partition.
```
targetPosition = Math.abs(Utils.murmur2(keyBytes)) % (numPartitions - 1)
```
## Offset
Incremental ID unique to each partition. Guarantees order within a partition (but not in other partitions).
