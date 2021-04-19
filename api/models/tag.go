package models

type Tag struct {
	Id          string   `json:"id" dynamodbav:"userId:tag"`
	BookmarkIds []string `json:"bookmarkIds"`
}
