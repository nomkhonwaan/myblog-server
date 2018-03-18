package post

import (
	"time"

	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
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
	Tags        []tag.Tag     `bson:"tags"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
	PublishedAt time.Time     `bson:"publishedAt"`
}

func NewPlaceholder() dataloader.Placeholder {
	return dataloader.Placeholder(&Post{})
}

func (p *Post) Key() string {
	return p.ID.Hex()
}

func (p *Post) IsPublished() bool {
	if p.Status == Published || !p.PublishedAt.IsZero() {
		return true
	}

	return false
}

type OrderField string

func (o OrderField) ToCamelCase() string {
	switch o {
	case "SLUG":
		return "slug"
	case "CREATED_AT":
		return "createdAt"
	case "UPDATED_AT":
		return "updatedAt"
	case "PUBLISHED_AT":
		return "publishedAt"
	default:
		return string(o)
	}
}
