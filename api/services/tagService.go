package services

import (
	"Brice-Ridings/leaftea_bookmarks_api/models"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

func GetTagById(id string) (*models.Tag, error) {
	svc := dynamodb.New(session.Must(session.NewSession()))
	tagsTableName := os.Getenv("BOOKMARKS_TAGS_TABLE_NAME")

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tagsTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"userId:tag": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	var tag models.Tag
	if result.Item != nil {
		err = dynamodbattribute.UnmarshalMap(result.Item, &tag)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}
	}

	return &tag, err
}
