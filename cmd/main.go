package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc"

	"github.com/buffup/GolangTechTask/api"
	repo "github.com/buffup/GolangTechTask/api/repository/dynamodb"
	pb "github.com/buffup/GolangTechTask/pkg/api"
)

func main() {
	logger := log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)

	os.Setenv("AWS_ACCESS_KEY_ID", "DUMMYIDEXAMPLE")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "DUMMYIDEXAMPLE")

	creds := credentials.NewEnvCredentials()
	creds.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String("us-west-2"),
		Endpoint:                      aws.String("http://localhost:8000"),
		CredentialsChainVerboseErrors: aws.Bool(true),
		Credentials:                   creds,
	})
	if err != nil {
		logger.Fatalf("could not create session: %s", err)
	}

	db := dynamodb.New(sess)
	repo := repo.NewVoteableRepo(db)

	var s api.VotingServer
	s = api.NewServer(repo)
	s = api.NewLogger(logger, s)

	grpc := grpc.NewServer()
	pb.RegisterVotingServiceServer(grpc, s)

	errs := make(chan error)
	go func() {

		grpcAddr := ":9000"
		// setup listener for gRPC requests
		listener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			errs <- err
			return
		}

		logger.Println("serving requests", "addr", grpcAddr)
		errs <- grpc.Serve(listener)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Println("msg", "app terminated", "reason", <-errs)
}
