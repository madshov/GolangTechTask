# Golang Tech Test
After running `docker-compose up` and starting the app, the following commands can be used (requires grpcurl)
```
grpcurl -v --plaintext -d '{"question":"What is pi?", "answers":["~3.14","22/7","a dessert"]}' -proto api/service.proto localhost:9000 VotingService/CreateVoteable
grpcurl -v --plaintext -d '{"page_number":1,"result_per_page":1}' -proto api/service.proto localhost:9000 VotingService/ListVoteables
grpcurl -v --plaintext -d '{"uuid":"<uuid>","answer_index":0}' -proto api/service.proto localhost:9000 VotingService/CastVote
```
Docker-compose will create the database container and create a table `Voteable`.

As part of the recruitment process we want to know how you think, code and structure your work.
In order to do that, we're going to ask you to complete this coding challenge.

**Please do not spend more than 5 hours on this task**, as this would not be
respectful of your time.

**Please do not submit work as a PR as these are publicly visible**

## Task

In the `api/` folder of this repo there is a basic `grpc` service definition for a voting service.
This service contains RPCs for creating, listing and voting on `voteable` items.

There is also a "hello world" go application in `cmd/` and a `docker-compose.yml` for running
`Amazon DynamoDB` locally.

We need you to:

- Provide a `Go` implementation of the `GRPC` service in the `cmd/` directory of this repo.
- Implement a `DynamoDB` based store for this `GRPC` service
- Add pagination to the `ListVoteables` RPC call
- Provide adequate test coverage for this simple service


## How to impress us

There are a few optional tasks you can complete if you really want to show off.

1. Adding Observability 

    Adding structured logging and/or tracing and metrics.
    (The current tech used should be considered when choosing technologies)

2. Adding Configuration and Secrets management
