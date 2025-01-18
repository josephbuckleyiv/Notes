# Implementation
## Pay Attention to Partitions Count and Replication
If partitions count increases during a topic lifecycle, you will break your keys ordering guarantees.
If the replication factor increases during a topic lifecycle, you put more pressure on your cluster.
### Partitions
More partitions means better parallelism, better throughput. Ability to leverage more brokers. Ability to run more consumers in a group to scale.
BUT more elections in zookeeper AND more files opened. The questions is how many partitions per topic: Small cluster(<6) -- 3 x # Brokers
Big cluster (>12) 2 x # brokers.
### Replication
2,3 (usually) or 4. But more replication means higher latency and more disk space on your systems.
Soft limit -- 4000 partitions per broker. 200, 000 total.
## Pay Careful attention to advertised host.
Initial handshake differs from ADV_HOST -- the former might expose a public IP address while the latter a private one. This discrepancy can cause connectivity issues. So advertised.listeners -- if clients are on private network, set inernal private IP or the internal private DNS hostname. The client resolves IP.
If on a public network, set to the external public IP, or the external public hostname pointing to the public IP.
## Topic Configuration
Impacts perfomance and topic behavior.
### Partitions
are made of segments. Only one segment is ACTIVE at a time. Maximum size is 1GB (log.segment.bytes). Segments come with two indexes -- an offset to position index, and a timestamp to offset index. Smaller log.segment.bytes means more segments per partitions, log compaction happens more frequently, BUT more files need to be opened. If you have a smaller log.segment.ms, you are setting a max frequency for log compaction.
### Log Compaction

### Log Cleanup Policies
log.cleanup.policy=delete (delete based on age of data or delete based on max size of log) \
log.cleanup.polcy=compact ( delete based on keys or delete old duplicate keys after the active segment is committed) \
Ability to control the size of the data on the disk, delete obsolete data. Overall: limit maintenance work on the kafka cluster. \
log.retention.hours (# hours to keep data, higher number means more disk space, log retention
Dedepulication is done after a segment is committed.
### Large Messages
Broker-side, set message.max.bytes. CLient side max.message.bytes

## Dangerous
unclean.leader.election.enable: enabled to start producing to non ISR partitions -> this improves availability, but you could lose data because other messages on ISR will be discard when they come back online and replicate data from the new leader. USE CASES: metrics collection and log collection, whre loss of data is acceptable.
## CLI
```
--describe
kafka-configs --bootstrap-server localhost:9092 --entity-type topics --entity-name configured-topic --describe
kafka-topics --bootstrap-server localhost:9092 --describe  --topic __consumer_offsets.

```
