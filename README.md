# Golang Tech Test
After running `docker-compose up` and starting the app, the following commands can be used (requires grpcurl)
```
grpcurl -v --plaintext -d '{"question":"What is pi?", "answers":["~3.14","22/7","a dessert"]}' -proto api/service.proto localhost:9000 VotingService/CreateVoteable
grpcurl -v --plaintext -d '{"page_number":1,"result_per_page":1}' -proto api/service.proto localhost:9000 VotingService/ListVoteables
grpcurl -v --plaintext -d '{"uuid":"<uuid>","answer_index":0}' -proto api/service.proto localhost:9000 VotingService/CastVote
```
Docker-compose will create the database container and create a table `Voteable`. in the DynamoDB instance.

- `CreateVoteable` will takes a question and a list of answers as input and creates a new voteable record in the database. A uuid of the created resource is returned.
- `ListVotables` Fetches a list of voteable records from the database. The list can be controlled by a page number and result per page parameter given as input.
- `Castvote` I'm not sure exactly what this function is supposed to do. It takes a voteable uuid and answer index as input, but doesn't return anything. 
