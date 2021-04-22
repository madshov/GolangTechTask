package api

import (
	"context"
	"testing"

	pb "github.com/buffup/GolangTechTask/pkg/api"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type mockRepo struct {
	mockCreateVoteable func() error
	mockGetVoteables   func() ([]*Voteable, error)
	mockGetVoteable    func() (*Voteable, error)
}

func (m *mockRepo) CreateVoteable(ctx context.Context, v *Voteable) error {
	if m.mockCreateVoteable != nil {
		return m.mockCreateVoteable()
	}
	return nil
}

func (m *mockRepo) GetVoteables(ctx context.Context, page, per_page int32) ([]*Voteable, error) {
	if m.mockGetVoteables != nil {
		return m.mockGetVoteables()
	}
	return nil, nil
}

func (m *mockRepo) GetVoteable(ctx context.Context, ID string) (*Voteable, error) {
	if m.mockGetVoteable != nil {
		return m.mockGetVoteable()
	}
	return nil, nil
}

func TestCreateVoteable_Success(t *testing.T) {
	have := &pb.CreateVoteableRequest{
		Question: "question",
		Answers:  []string{"answer1", "answer2"},
	}

	repo := &mockRepo{
		mockCreateVoteable: func() error {
			return nil
		},
	}

	s := NewServer(repo)

	got, err := s.CreateVoteable(context.Background(), have)

	if err != nil {
		t.Errorf("unexpected error returned '%s'", err.Error())
	}

	if got.Uuid == "" {
		t.Errorf("unexpected return value")
	}
}

func TestListVoteables_Success(t *testing.T) {
	have := &pb.ListVoteableRequest{
		PageNumber:    0,
		ResultPerPage: 0,
	}
	want := &pb.ListVoteablesResponse{
		Votables: []*pb.Voteable{
			&pb.Voteable{
				Uuid:     "uuid1",
				Question: "question1",
				Answers:  []string{"answer1", "answer2"},
			},
			&pb.Voteable{
				Uuid:     "uuid2",
				Question: "question2",
				Answers:  []string{"answer3", "answer4"},
			},
		},
	}

	repo := &mockRepo{
		mockGetVoteables: func() ([]*Voteable, error) {
			return []*Voteable{
				&Voteable{
					UUID:     "uuid1",
					Question: "question1",
					Answers:  []string{"answer1", "answer2"},
				},
				&Voteable{
					UUID:     "uuid2",
					Question: "question2",
					Answers:  []string{"answer3", "answer4"},
				},
			}, nil
		},
	}

	s := NewServer(repo)

	got, err := s.ListVoteables(context.Background(), have)

	if err != nil {
		t.Errorf("unexpected error returned '%s'", err.Error())
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(pb.ListVoteablesResponse{}, pb.Voteable{})) {
		t.Errorf("expected return value to be %v, got %v", want, got)
	}
}
