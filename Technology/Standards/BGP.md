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
### Withdrawn Route
Negative advertisement that instructs other speakers to scrub paths from their tables.
## Goal
Scalability and retrieval of policy-compliant path. Only ensures reachability -- cannot comment on weight of path.

## Structure

The result of BGP is a complete-path.
Policy is made the concern of whoever gets the path.
Loop-prevention is naively done by seeing if I (the node who is calculating the path) see myself in the complete path.
This requires uniquity of AS idenitifer -- currently these are 32-bit ints assigned by central authority. Speaker only advertises the best one according to its own policies. Furthmore, a speaker is under no obligation to let others know of an extant path - hence, transit traffic may be denied. BGP runs on top of TCP, hence assumes reliability -- so upkeep is the occassional keepalive message.
## Common AS Relationships
If an AS has no provider ASes it's known as a Tier-1 Provider.
### Provider-Customer: Providers connect customers to internet. Advertise all my routes to my cutomers, and advertise everything I learn about my customer to everyone.
### Customer-Provider: Advertise my own prefixes and routes to my customers and my provider. DON'T advertise to other provider. I DON'T WANT TO BE A MIDDLE MAN.

### Peer: Symmetrical peering between autonomous systems. Providers view each other as equals and advertise routes from peer to customer and vice-versa. DON'T advertise routes from my peer to any provider. \
## Integrating Inter- and Intra-
A stub AS might inject a default route into the intradomain protocol. Recall how IP forwarding works.
A provider might inject the route to a customer into its routing protocol. The problem is backbone AS which have too much routing information. Use variant of BGP called interior BGP: calculates best border router to use when sending packets.
