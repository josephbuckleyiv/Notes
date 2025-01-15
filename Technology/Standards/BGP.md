# Border Gateway Protocol
No assumption about network topology (versus Exterior Gateway Protocol which assumed Tree-like structure)
Manages interconnection between Autonomous systems. 
Local traffic is distinct from transit traffic. 

## Types of Autonomous Systems
- Stub - only a single connection to one other AS
- Multihomed - connections between more than one other AS, but REFUSES transit traffic
- Transit - connections between more than one other AS, but carries both transit and local traffic.
## Goal
Scalability and retrieval of policy-compliant path. Only ensures reachability -- cannot comment on weight of path.

## Structure
Border routers -- packets AS ingress/egress. At one time called a gateway.
BGP speaker - communicates/co-ordinates with BGP speakers in other AS.
The result of BGP is a complete-path.
Policy is made the concern of whoever gets the path.
Loop-prevention is naively done by seeing if I (the node who is calculating the path) see myself in the complete path.
