This repo is an example that I came up with to showcase how to easily start up
services in order. This solves a common problem I've observed: without more
complex controls, it's difficult to start up a database, wait for that database
to be ready, then run migrations, and *only then* start up services. There are
lots of hacky, non-Docker-native solutions, but within the last 6 months or so
Docker has added native features that make this easy to implement in a
`docker-compose` file without writing more code outside of Docker.

### What's going on

The services are listed in the order that they need to be brought up in.

First, a mock database starts. This mock database is a small Go program that
responds to the `/db` endpoint on port 1212.

Second, the mock database's migrations run. This is just a `wget` command to hit
the database's `/db` endpoint, but Docker keeps retrying it until it can
successfully reach the endpoint, and then `wget returns a 0, indicating a
successful migration run.

Third, a web API that relies on that database to be set up properly runs. This
web API will only start up once the mock database migrations successfully
finish. This web API has a healthcheck endpoint, so we know when it is ready to
respond to web requests.

Fourth, a web service that uses the web API starts once Docker receives a
"healthy" response from the web API. This way we know that the web service can
make requests from the API.

### Other solutions I've used in the past

* Repeated manual retries
* Using `docker compose` to bring up one tier of services at a time - databases,
then migrations, then services that use those migrations, etc.
* https://github.com/ufoscout/docker-compose-wait
