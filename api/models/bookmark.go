package models

type Bookmark struct {
	Id     string   `json:"id"`
	Url    string   `json:"url"` // TODO: Handle base64 decode/encode?
	UserId string   `json:"userId" dynamodbav:"user_id"`
	Tags   []string `json:"tags"`
}

type NewBookmark struct {
	Url  string   `json:"url"`
	Tags []string `json:"tags"`
}
