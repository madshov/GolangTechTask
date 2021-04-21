package api

import context "context"

type VoteableStorage interface {
	CreateVoteable(ctx context.Context, v *Voteable) error
	GetVoteables(ctx context.Context, page, per_page int32) ([]*Voteable, error)
	GetVoteable(ctx context.Context, ID string) (*Voteable, error)
}

type Voteable struct {
	UUID     string   `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}
