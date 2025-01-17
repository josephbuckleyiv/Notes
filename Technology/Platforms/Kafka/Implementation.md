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
