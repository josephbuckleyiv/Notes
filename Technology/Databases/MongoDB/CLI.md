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
### Find Something
```
db.records.findOne( {document})
db.records.find( {documen})
```
### Read Concern
Configure the readConcern by calling method on find or insert
```
readConcern("local" | "available" | "majority" | "linearizable")
```
### Read Preference
```
readPreference("primary" | "secondary" | "nearest")
```
Fine for analytics, but don't use to scale normal traffic. 
### Comparisons
```
$eq, $gt, $gte, $lt, $lte, $in, $nin
```
### Logical Operators
```
db.inventory.findOne( $and: [{ "variations.quantity" : {$ne : 0}}] )
```
### Sort, Limit, Skyp
Sort via an index. If you can't define an index, use limit. Remember MongoDB always performs sort, skip and limit in that order.

### Element Match
```
db.inventory.find({variations: {$elemMatch: { varitian: "Blue", quantity: {$gte : 7}} })
```

### Push, Pop
```
db.movies.updateOne({, { $push" { genres: "Test"}) --> Adds "test" to genres array
db.movies.updateOne({, { $pop" { genres: "Test"}) --> Removes it.

```
