# Border Gateway Protocol
No assumption about network topology (versus Exterior Gateway Protocol which assumed Tree-like structure)
Manages interconnection between Autonomous systems. 
Local traffic is distinct from transit traffic. 

## Terminology
### Types of Autonomous Systems
- Stub - only a single connection to one other AS
- Multihomed - connections between more than one other AS, but REFUSES transit traffic
- Transit - connections between more than one other AS, but carries both transit and local traffic.
### Border routers
Packets AS ingress/egress. At one time called a gateway.
### BGP speaker 
Communicates/co-ordinates with BGP speakers in other AS.
## Goal
Scalability and retrieval of policy-compliant path. Only ensures reachability -- cannot comment on weight of path.

## Structure

The result of BGP is a complete-path.
Policy is made the concern of whoever gets the path.
Loop-prevention is naively done by seeing if I (the node who is calculating the path) see myself in the complete path.
This requires uniquity of AS idenitifer -- currently these are 32-bit ints assigned by central authority. Speaker only advertises the best one according to its own policies. Furthmore, a speaker is under no obligation to let others know of an extant path - hence, transit traffic may be denied.
