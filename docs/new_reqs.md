# New Work

## Overview
There's been a good amount of work done to collect machine stats, and the product is doing well, after having been
deployed to a kubernetes cluster. That said, there's been a change in requirements. While the RESTful http server is 
going to stay and be continued, there's been a formal request to include gRPC endpoints. This is namely to take 
advantage of http/2 channels, and familiarity of gRPC with our desktop developers.

Furthermore, we want to include tracing from the request through to the database, as our SREs have reported 
inconsistency in the system, and they'd like to see where this issue is happening.

Finally, We can no longer limit users based on their IP addresses. We need to introduce some basic authorization and
authentication, ideally in the form of a token.

## The Job
**_we aren't writing any code here._** Our job is to see what changes to the service you wrote need to be made, and how 
we can satisfy the requests. Feel free to use any diagraming/whiteboarding tool, or we can simply talk things through.

### More details
- Assume we're running the service in a kubernetes cluster, and that we can deploy additional services as needed.
- Assume the service is behind some type of api gateway (i.e: nginx). And that we can modify or hook into it at will.
- Assume the database is actually being used in the cloud (SQL, mongo, doesn't matter).
