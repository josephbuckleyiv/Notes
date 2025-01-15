# Service Health
Often times, various concerns will become interdependent -- telemetry might depend on session and session on telemetry or whatnot! -- and we would like to poke the system to see if it's healthy. In that case, we need to check on our dependencies to see if th . Isn't it circular? Well, we can break the chain by making a decision. And decision comes in two flavors:
- I have just received a request asking if I'm okay. Only someone who isn't in distress themselves, and isnt worried about someone else already can ask me that.
- The only person I'm worried about asked me if I'm okay. And I'm okay. Everything is hunky dory
This corresponds to the situation where,
  - Node receives request and has multiple outgoing : DECISION: The Node asking me is working, and no-one should ask it again. I'll ask those around other Nodes I depend on to see if I'm working.
  - Node receives request and has no other dependencies: DECISION: The node asking me must be fine, and I depend on it. And I'm fine. So tell the Node I'm fine.

  ## Our goal is to avoid caching or state.
  In one fell swoop, we can see if the entire system is up and running. If it isn't, something's afoot, and we leave it to telemetry and another system to do the forensics and damage control.
