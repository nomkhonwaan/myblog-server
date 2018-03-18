package tag

import (
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"gopkg.in/mgo.v2/bson"
)

type Tag struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Slug string        `bson:"slug"`
}

func NewPlaceholder() dataloader.Placeholder {
	return dataloader.Placeholder(&Tag{})
}

func (t *Tag) Key() string {
	return t.ID.Hex()
}
