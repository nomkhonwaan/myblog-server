package post

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	Draft     = Status("DRAFT")
	Published = Status("PUBLISHED")
)

type Status string

type Post struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string        `bson:"title"`
	Slug        string        `bson:"slug"`
	Status      Status        `bson:"status"`
	HTML        string        `bson:"html"`
	Markdown    string        `bson:"markdown"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
	PublishedAt time.Time     `bson:"publishedAt"`
}
