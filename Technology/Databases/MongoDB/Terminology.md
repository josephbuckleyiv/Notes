# Terminology
## Mongod
Daemon Process for mongod. Handles data requests from the MongoDB shell and performs background management
operations.
```
mongod -dbpath mongod_only
```
Default settings listen on port 27017.
Then run mongosh
## Durability
Can configure using 
```
{
  w: "majority"
  j: true
  wtimeout:
}
```
property when inserting an object. Write-concern means more than the "majority" must have the write, before an acknowledgement is sent to the client.
## Atomicity
Require MultiDocument Transactions
