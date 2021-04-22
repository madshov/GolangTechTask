# Golang Tech Test
After running `docker-compose up` and starting the app, the following commands can be used (requires grpcurl)
```
grpcurl -v --plaintext -d '{"question":"What is pi?", "answers":["~3.14","22/7","a dessert"]}' -proto api/service.proto localhost:9000 VotingService/CreateVoteable
grpcurl -v --plaintext -d '{"page_number":1,"result_per_page":1}' -proto api/service.proto localhost:9000 VotingService/ListVoteables
grpcurl -v --plaintext -d '{"uuid":"<uuid>","answer_index":0}' -proto api/service.proto localhost:9000 VotingService/CastVote
```
Docker-compose will create the database container and create a table `Voteable`.
