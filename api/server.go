package api

import (
	context "context"
	"errors"

	uuid "github.com/satori/go.uuid"

	pb "github.com/buffup/GolangTechTask/pkg/api"
)

type VotingServer interface {
	CreateVoteable(ctx context.Context, req *pb.CreateVoteableRequest) (*pb.CreateVoteableResponse, error)
	ListVoteables(ctx context.Context, req *pb.ListVoteableRequest) (*pb.ListVoteablesResponse, error)
	CastVote(ctx context.Context, req *pb.CastVoteRequest) (*pb.CastVoteResponse, error)
}

type server struct {
	repo VoteableStorage
}

func NewServer(r VoteableStorage) *server {
	return &server{
		repo: r,
	}
}

// CreateVoteable receives a request and calls CreateVoteable method on the
// repository. It then returns a response with a created UUID.
func (s *server) CreateVoteable(ctx context.Context, req *pb.CreateVoteableRequest) (*pb.CreateVoteableResponse, error) {
	v := Voteable{
		UUID:     uuid.NewV4().String(),
		Question: req.Question,
		Answers:  req.Answers,
	}

	err := s.repo.CreateVoteable(ctx, &v)
	if err != nil {
		return nil, err
	}

	return &pb.CreateVoteableResponse{
		Uuid: v.UUID,
	}, nil
}

// ListVoteables receieves a requests and gets all voteable records from the
// repository. It then returns a response with the list of records.
func (s *server) ListVoteables(ctx context.Context, req *pb.ListVoteableRequest) (*pb.ListVoteablesResponse, error) {
	vs, err := s.repo.GetVoteables(ctx, req.PageNumber, req.ResultPerPage)
	if err != nil {
		return nil, err
	}

	vos := make([]*pb.Voteable, len(vs))
	for k, v := range vs {
		vo := &pb.Voteable{
			Uuid:     v.UUID,
			Question: v.Question,
			Answers:  v.Answers,
		}

		vos[k] = vo
	}

	return &pb.ListVoteablesResponse{
		Votables: vos,
	}, nil
}

// CastVote get a specific voteable record from the repository and checks if
// the given answer index exist.
func (s *server) CastVote(ctx context.Context, req *pb.CastVoteRequest) (*pb.CastVoteResponse, error) {
	v, err := s.repo.GetVoteable(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	if int(req.AnswerIndex) > len(v.Answers)-1 {
		return nil, errors.New("Answer index is out of bounds")
	}

	// do something with answer
	//ans := v.Answers[req.AnswerIndex]

	return &pb.CastVoteResponse{}, nil
}
