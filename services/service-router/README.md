# SERVICE-ROUTER

service-router is a service that can be used to route requests based on X-TargetService
header.
It is basically a reverse proxy used to route requests to a service without the need to
know about the service's address or port. In this context it is used for services to call
each other by name (X-TargetService header).
This allows us to quickly swap services, redeploy them somewhere else and stuff like that,
without the need of having to reconfigure all other services