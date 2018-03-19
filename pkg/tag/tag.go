package tag

import (
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"gopkg.in/mgo.v2/bson"
)

// Tag represent an entity of Tag
type Tag struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Slug string        `bson:"slug"`
}

// Key returns a Placeholder key string
func (t *Tag) Key() string {
	return t.ID.Hex()
}

// NewPlaceholder returns a new Tag's object
func NewPlaceholder() dataloader.Placeholder {
	return dataloader.Placeholder(&Tag{})
}
