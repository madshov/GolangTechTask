package dynamodb

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	dyndb "github.com/aws/aws-sdk-go/service/dynamodb"
	dyndbattr "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/buffup/GolangTechTask/api"
)

type voteableRepo struct {
	DB dynamodbiface.DynamoDBAPI
}

// NewVoteableRepo creates and returns a new instance of a voteable repository.
func NewVoteableRepo(DB dynamodbiface.DynamoDBAPI) *voteableRepo {
	return &voteableRepo{
		DB: DB,
	}
}

// CreateViteable creates a DynamoDB putItem request and persists to DB.
func (vr *voteableRepo) CreateVoteable(ctx context.Context, v *api.Voteable) error {

	input := &dyndb.PutItemInput{
		Item: map[string]*dyndb.AttributeValue{
			"Id": {
				S: aws.String(v.UUID),
			},
			"Question": {
				S: aws.String(v.Question),
			},
			"Answers": {
				SS: aws.StringSlice(v.Answers),
			},
		},
		TableName: aws.String("Voteable"),
	}

	_, err := vr.DB.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

// GetVoteables creates a DynamoDB Scan request and returns all records.
// Total can be controlled via page and per_page parameters.
func (vr *voteableRepo) GetVoteables(ctx context.Context, page, per_page int32) ([]*api.Voteable, error) {
	input := &dyndb.ScanInput{
		TableName: aws.String("Voteable"),
	}

	result, err := vr.DB.Scan(input)
	if err != nil {
		return nil, err
	}

	var vo []*api.Voteable

	count := int(*result.Count)
	stIdx := 0
	endIdx := 0

	if count > 0 {
		endIdx = count
	}

	// translate page and per_page into a start index and end index.
	// page and per_page are 1-indexed.
	if page > 0 && per_page > 0 {
		p := int(page)
		pp := int(per_page)

		if p > count {
			p = 1
		}

		stIdx = ((p - 1) * pp)
		if stIdx > count {
			stIdx = count - pp
			if stIdx < 0 {
				stIdx = 0
			}
		}

		endIdx = stIdx + pp
		if endIdx > count {
			endIdx = count
		}
	}

	items := result.Items[stIdx:endIdx]
	for _, i := range items {
		var v *api.Voteable

		err = dyndbattr.UnmarshalMap(i, &v)
		if err != nil {
			return nil, err
		}

		vo = append(vo, v)
	}

	return vo, nil
}

// GetVoteable retrieves a voteable record for a given ID in the repositoru.
func (vr *voteableRepo) GetVoteable(ctx context.Context, ID string) (*api.Voteable, error) {
	input := &dyndb.QueryInput{
		ExpressionAttributeValues: map[string]*dyndb.AttributeValue{
			":v1": {
				S: aws.String(ID),
			},
		},
		KeyConditionExpression: aws.String("Id = :v1"),
		TableName:              aws.String("Voteable"),
	}

	result, err := vr.DB.Query(input)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, errors.New("voteable not found")
	}

	var v *api.Voteable
	err = dyndbattr.UnmarshalMap(result.Items[0], &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
