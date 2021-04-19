package handlers

import (
	"Brice-Ridings/leaftea_bookmarks_api/models"
	"Brice-Ridings/leaftea_bookmarks_api/services"
	"Brice-Ridings/leaftea_bookmarks_api/utils"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

func CreateBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit Create Bookmark")
	userId := r.Header.Get("user")

	// Decode json body into struct, ToDo: utilize different names for request and database?
	var newBookmark models.NewBookmark
	err := json.NewDecoder(r.Body).Decode(&newBookmark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookmark := models.Bookmark{
		Id:     uuid.NewString(),
		Url:    newBookmark.Url,
		UserId: userId,
		Tags:   newBookmark.Tags,
	}

	// Create request and store item in dynamodb
	svc := dynamodb.New(session.Must(session.NewSession()))
	bookmarkDbItem, err := dynamodbattribute.MarshalMap(bookmark)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}

	if bookmark.Tags == nil || len(bookmark.Tags) <= 0 {
		bookmarksTableName := os.Getenv("BOOKMARKS_TABLE_NAME")
		_, err = svc.PutItem(&dynamodb.PutItemInput{
			Item:      bookmarkDbItem,
			TableName: &bookmarksTableName,
		})
		if err != nil {
			panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
		}
	} else {
		addBookmarkAndTags(svc, bookmarkDbItem, bookmark)
	}
	utils.ResponseJSON(w, http.StatusOK, bookmark)
}

// Add bookmark item via transaction to include tags request
func addBookmarkAndTags(svc *dynamodb.DynamoDB, bookmarkDbItem map[string]*dynamodb.AttributeValue, bookmark models.Bookmark) {
	bookmarksTableName := os.Getenv("BOOKMARKS_TABLE_NAME")
	tagsTableName := os.Getenv("BOOKMARKS_TAGS_TABLE_NAME")

	transactionRequestItems := []*dynamodb.TransactWriteItem{
		{
			Put: &dynamodb.Put{
				TableName: aws.String(bookmarksTableName),
				Item:      bookmarkDbItem,
			},
		},
	}

	// gather tags from db and store in list to be updated
	for _, tag := range bookmark.Tags {
		id := fmt.Sprintf("%v:%v", bookmark.UserId, tag)
		tagItem, err := services.GetTagById(id)
		if err != nil {
			panic(fmt.Sprintf("failed to get Tag %v from DynamoDB, %v", id, err))
		}

		if tagItem == nil {
			*tagItem = models.Tag{
				Id:          id,
				BookmarkIds: []string{},
			}
		}
		tagItem.BookmarkIds = append(tagItem.BookmarkIds, bookmark.Url)
		tagDbItem, err := dynamodbattribute.MarshalMap(tagItem)
		if err != nil {
			panic(fmt.Sprintf("failed to Marshal Tag %v, %v", id, err))
		}
		transactionItem := dynamodb.TransactWriteItem{
			Put: &dynamodb.Put{
				TableName: aws.String(tagsTableName),
				Item:      tagDbItem,
			},
		}
		transactionRequestItems = append(transactionRequestItems, &transactionItem)
	}

	// Run transaction request
	_, errDb := svc.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		TransactItems: transactionRequestItems,
	})
	if errDb != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", errDb))
	}
}
