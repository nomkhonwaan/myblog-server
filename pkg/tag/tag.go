package tag

import (
	"context"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	mgo "gopkg.in/mgo.v2"
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

// Repositorier is a Tag's repository interface
type Repositorier interface {
	FindByID(id string) (*Tag, error)
}

// Repository is an implemented of Tag's Repositorier interface
type Repository struct {
	db     *mgo.Database
	loader *dld.Loader
}

// NewRepository returns a new Tag's repository with dataloader configured
func NewRepository(db *mgo.Database) Repository {
	c, _ := cache.New(100)

	return Repository{
		db: db,
		loader: dld.NewBatchedLoader(
			dataloader.NewBatchFunc(db.C("tags"), NewPlaceholder),
			dld.WithCache(c),
		),
	}
}

// FindByID finds a single Tag from its ID
func (repo Repository) FindByID(id string) (*Tag, error) {
	t, err := repo.loader.Load(context.TODO(), dld.StringKey(id))()
	if err != nil {
		return nil, err
	}

	return t.(*Tag), nil
}
