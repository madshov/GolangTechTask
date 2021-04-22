package dynamodb

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dyndb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/buffup/GolangTechTask/api"
	"github.com/google/go-cmp/cmp"
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

func TestGetVoteables1Success(t *testing.T) {
	want := []*api.Voteable{
		&api.Voteable{
			UUID:     "uuid1",
			Question: "question1",
			Answers:  []string{"answer1", "answer2"},
		},
	}

	db := &mockDynamoDB{
		mockScan: func(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
			return &dynamodb.ScanOutput{
				Count: aws.Int64(2),
				Items: []map[string]*dyndb.AttributeValue{
					map[string]*dyndb.AttributeValue{
						"Id": {
							S: aws.String("uuid1"),
						},
						"Question": {
							S: aws.String("question1"),
						},
						"Answers": {
							SS: aws.StringSlice([]string{"answer1", "answer2"}),
						},
					},
					map[string]*dyndb.AttributeValue{
						"Id": {
							S: aws.String("uuid2"),
						},
						"Question": {
							S: aws.String("question2"),
						},
						"Answers": {
							SS: aws.StringSlice([]string{"answer3", "answer4"}),
						},
					},
				},
			}, nil
		},
	}

	repo := NewVoteableRepo(db)

	got, err := repo.GetVoteables(context.Background(), 1, 1)

	if err != nil {
		t.Errorf("unexpected error returned '%s'", err.Error())
	}

	if !cmp.Equal(want, got) {
		t.Errorf("expected return value to be %v, got %v", want, got)
	}
}

func TestGetVoteables2Success(t *testing.T) {
	want := []*api.Voteable{
		&api.Voteable{
			UUID:     "uuid2",
			Question: "question2",
			Answers:  []string{"answer3", "answer4"},
		},
	}

	db := &mockDynamoDB{
		mockScan: func(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
			return &dynamodb.ScanOutput{
				Count: aws.Int64(2),
				Items: []map[string]*dyndb.AttributeValue{
					map[string]*dyndb.AttributeValue{
						"Id": {
							S: aws.String("uuid1"),
						},
						"Question": {
							S: aws.String("question1"),
						},
						"Answers": {
							SS: aws.StringSlice([]string{"answer1", "answer2"}),
						},
					},
					map[string]*dyndb.AttributeValue{
						"Id": {
							S: aws.String("uuid2"),
						},
						"Question": {
							S: aws.String("question2"),
						},
						"Answers": {
							SS: aws.StringSlice([]string{"answer3", "answer4"}),
						},
					},
				},
			}, nil
		},
	}

	repo := NewVoteableRepo(db)

	got, err := repo.GetVoteables(context.Background(), 2, 1)

	if err != nil {
		t.Errorf("unexpected error returned '%s'", err.Error())
	}

	if !cmp.Equal(want, got) {
		t.Errorf("expected return value to be %v, got %v", want, got)
	}
}

func TestGetVoteables3Success(t *testing.T) {
	want := []*api.Voteable{
		&api.Voteable{
			UUID:     "uuid1",
			Question: "question1",
			Answers:  []string{"answer1", "answer2"},
		},
		&api.Voteable{
			UUID:     "uuid2",
			Question: "question2",
			Answers:  []string{"answer3", "answer4"},
		},
	}

	db := &mockDynamoDB{
		mockScan: func(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
			return &dynamodb.ScanOutput{
				Count: aws.Int64(2),
				Items: []map[string]*dyndb.AttributeValue{
					map[string]*dyndb.AttributeValue{
						"Id": {
							S: aws.String("uuid1"),
						},
						"Question": {
							S: aws.String("question1"),
						},
						"Answers": {
							SS: aws.StringSlice([]string{"answer1", "answer2"}),
						},
					},
					map[string]*dyndb.AttributeValue{
						"Id": {
							S: aws.String("uuid2"),
						},
						"Question": {
							S: aws.String("question2"),
						},
						"Answers": {
							SS: aws.StringSlice([]string{"answer3", "answer4"}),
						},
					},
				},
			}, nil
		},
	}

	repo := NewVoteableRepo(db)

	got, err := repo.GetVoteables(context.Background(), 1, 2)

	if err != nil {
		t.Errorf("unexpected error returned '%s'", err.Error())
	}

	if !cmp.Equal(want, got) {
		t.Errorf("expected return value to be %v, got %v", want, got)
	}
}
