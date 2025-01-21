# CLI
### Display
```
show databases
```
### Insert something.
```
db.xXXx.insertOne({"hello", "world"})
```
## Linux
### Create a database in a Replica Set
```
mongod --replSet myReplSet --dbpath ./m1/db --logpath ./m1/mongodb.log --port 27017 --fork --keyFile ./keyfile
```
### Initiate Set
```
rs.initiate()
use admin
```
Make sure admin has the ability to create users.
```
db.createUser({user: 'naomi', pwd: passwordPrompt(), roles: ["root"]}
db.getSiblingDB("admin).auth("naomi", passwordPrompt())
rs.add("localhost:27018") 
```
