package api

import (
	"context"
	"log"
	"time"

	pb "github.com/buffup/GolangTechTask/pkg/api"
)

type logger struct {
	log *log.Logger
	VotingServer
}

func NewLogger(log *log.Logger, vs VotingServer) VotingServer {
	return &logger{log, vs}
}

func (l *logger) CreateVoteable(ctx context.Context, req *pb.CreateVoteableRequest) (resp *pb.CreateVoteableResponse, err error) {
	defer func(begin time.Time) {
		if err == nil {
			l.log.Println(
				"method", "create-voteable",
				"took", time.Since(begin),
			)
		} else {
			l.log.Println(
				"method", "create-voteable",
				"error", err,
				"took", time.Since(begin),
			)
		}
	}(time.Now())

	return l.VotingServer.CreateVoteable(ctx, req)
}

func (l *logger) ListVoteables(ctx context.Context, req *pb.ListVoteableRequest) (resp *pb.ListVoteablesResponse, err error) {
	defer func(begin time.Time) {
		if err == nil {
			l.log.Println(
				"method", "list-voteables",
				"took", time.Since(begin),
			)
		} else {
			l.log.Println(
				"method", "list-voteables",
				"error", err,
				"took", time.Since(begin),
			)
		}
	}(time.Now())

	return l.VotingServer.ListVoteables(ctx, req)
}

func (l *logger) CastVote(ctx context.Context, req *pb.CastVoteRequest) (resp *pb.CastVoteResponse, err error) {
	defer func(begin time.Time) {
		if err == nil {
			l.log.Println(
				"method", "cast-vote",
				"took", time.Since(begin),
			)
		} else {
			l.log.Println(
				"method", "cast-vote",
				"error", err,
				"took", time.Since(begin),
			)
		}
	}(time.Now())

	return l.VotingServer.CastVote(ctx, req)
}
