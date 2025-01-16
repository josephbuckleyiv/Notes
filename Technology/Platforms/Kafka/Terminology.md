# Terminology
## Broker
A server within a Kafka Cluster that receives messages from producers, stores them in partitions, and serves those messages to consumers.
In diagrams, this 'houses' the topics/partitions from which consumers pull. After connecting to any broker _bootstrap broker__, you will
connect to whole cluster.
## Consumers
Read data from the topic, via 'pull' model. They know which broker to read from. New ability to read from the 'closest' replica.

## Consumer Groups
All the consumers in an application read data as a group. Each consumer reads from exclusive partitions. Multiple consumer groups can
be on the same topic. But in each consumer group, each partition has a unique consumer.
## Topics
A stream of data. Can have as many as you want, and is identifiable via Name. Similar to a table in a database, but without 
constaints. The sequence of data passing to Topic is called a data stream. Cannot query topics. Immutable -- once written, cannot be 
changed. Data is kept for a limited time. MUST NOT CHANGE DATATYPE. Create a new topic.
### Replication Factor
Should be above 1, this way if a broker goes down, another broker can serve the data.
## Partitions
Topics are split in partitions. Messages within each partition are ordered. At any time only one broker can be the leader of a 
partition. If replication happens fast enough, our duplicates will be called In-Sync Replicas.
## Producers
Writes to topics. Producers know to which partition to write to. In case of broker failures, producers will recover. Can
send key with the message. Key=null means data is sent round robin. Key=null, then all messages for that key will always go to
the same partition.
## Producer Acknowledgements
- acks=0 : Producer won't wait for acknowledgment
- acks=1 : Producer will wait for leader acknowledgment.
- acks=all/-1 : Leader + replicas acknowledgement (no data loss)
## Kafka Message
Key-binary, value-binary, Compression Type, Headers, Partition + Offset, Timestamp \
Message Serialization is used because Kafka only accepts bytes. Hash assigns key to partition.
```
targetPosition = Math.abs(Utils.murmur2(keyBytes)) % (numPartitions - 1)
```
## Offset
Incremental ID unique to each partition. Guarantees order within a partition (but not in other partitions).
Kafka stores offsets at which a consumer group has been reading in special topic call __consumer_offets.
Commit Offets:
- At least once -- commit after the message is processed. If processing fails, read again.
- At most once -- once received. If processing fails, message lost.
- Exactly once [this is our goal]

