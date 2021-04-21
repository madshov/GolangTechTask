package dynamodb

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/buffup/GolangTechTask/api"
)

type mockDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	mockPutItem func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	mockScan    func(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}

func (m *mockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if m.mockPutItem != nil {
		return m.mockPutItem(input)
	}

	return nil, nil
}

func (m *mockDynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if m.mockScan != nil {
		return m.mockScan(input)
	}

	return nil, nil
}

func TestCreateVoteableSuccess(t *testing.T) {
	db := &mockDynamoDB{
		mockPutItem: func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
			return &dynamodb.PutItemOutput{}, nil
		},
	}

	repo := NewVoteableRepo(db)

	err := repo.CreateVoteable(context.Background(), &api.Voteable{
		UUID:     "uuid",
		Question: "question",
		Answers:  []string{"answer1, answer2"},
	})

	if err != nil {
		t.Errorf("unexpected error returned '%s'", err.Error())
	}
}

func TestCreateVoteableError(t *testing.T) {
	db := &mockDynamoDB{
		mockPutItem: func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
			return &dynamodb.PutItemOutput{}, errors.New("some error")
		},
	}

	repo := NewVoteableRepo(db)

	err := repo.CreateVoteable(context.Background(), &api.Voteable{
		UUID:     "uuid",
		Question: "question",
		Answers:  []string{"answer1, answer2"},
	})

	if err == nil {
		t.Errorf("expected error not returned '%s'", err.Error())
	}
}
