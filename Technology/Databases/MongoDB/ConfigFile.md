# Configuration
Ideally, use config file to spin up replica sets.
## How to
Make a keyfile. For production use X.509 certificates. The following is a conf file:
```
storage:
  dbPath: m1/db
net:
  bindIp: localhost
  port: 27017
security:
  authorization: enabled
  keyFile: keyfile
systemLog:
  destination: file
  path: m1/mongod.log
processManagement:
  fork: true
replication:
  replSetName: mongodb-essentials

```

