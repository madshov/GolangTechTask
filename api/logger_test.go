package api

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	pb "github.com/buffup/GolangTechTask/pkg/api"
)

type mockServer struct {
	mockCreateVoteable func() (*pb.CreateVoteableResponse, error)
	mockListVoteables  func() (*pb.ListVoteablesResponse, error)
	mockCastVote       func() (*pb.CastVoteResponse, error)
}

func (m *mockServer) CreateVoteable(ctx context.Context, req *pb.CreateVoteableRequest) (*pb.CreateVoteableResponse, error) {
	if m.mockCreateVoteable != nil {
		return m.mockCreateVoteable()
	}
	return nil, nil
}

func (m *mockServer) ListVoteables(ctx context.Context, req *pb.ListVoteableRequest) (*pb.ListVoteablesResponse, error) {
	if m.mockListVoteables != nil {
		return m.mockListVoteables()
	}
	return nil, nil
}

func (m *mockServer) CastVote(ctx context.Context, req *pb.CastVoteRequest) (*pb.CastVoteResponse, error) {
	if m.mockCastVote != nil {
		return m.mockCastVote()
	}
	return nil, nil
}

func TestCreateVoteable_LoggerSuccess(t *testing.T) {
	have := &pb.CreateVoteableRequest{
		Question: "question",
		Answers:  []string{"answer1", "answer2"},
	}
	want := &pb.CreateVoteableResponse{
		Uuid: "uuid",
	}

	server := &mockServer{
		mockCreateVoteable: func() (*pb.CreateVoteableResponse, error) {
			return want, nil
		},
	}

	logger := log.New(ioutil.Discard, "", 0)
	s := NewLogger(logger, server)

	got, err := s.CreateVoteable(context.Background(), have)

	if err != nil {
		t.Errorf("unexpected error returned '%s'", err.Error())
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(pb.CreateVoteableResponse{})) {
		t.Errorf("expected return value to be %v, got %v", want, got)
	}
}

func TestCreateVoteable_LoggerError(t *testing.T) {
	have := &pb.CreateVoteableRequest{
		Question: "question",
		Answers:  []string{"answer1", "answer2"},
	}

	server := &mockServer{
		mockCreateVoteable: func() (*pb.CreateVoteableResponse, error) {
			return nil, errors.New("some error")
		},
	}

	logger := log.New(ioutil.Discard, "", 0)
	s := NewLogger(logger, server)

	_, err := s.CreateVoteable(context.Background(), have)

	if err == nil {
		t.Errorf("expected error not returned '%s'", err.Error())
	}
}
