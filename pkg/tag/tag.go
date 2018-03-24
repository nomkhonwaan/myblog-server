package tag

import (
	"context"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	"github.com/nomkhonwaan/myblog-server/pkg/mongodb"
	"gopkg.in/mgo.v2/bson"
)

// Repositorier is a Tag's repository interface
type Repositorier interface {
	FindByID(id string) (*Tag, error)
	FindAll(
		offset, limit int,
		orderBy struct {
			Field     string
			Direction string
		},
	) ([]*Tag, error)
}

// Tag represent an entity of Tag
type Tag struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Slug string        `bson:"slug"`
}

// Key returns a Placeholder key string
func (t Tag) Key() string {
	return t.ID.Hex()
}

// Tags represent a list of Tags
type Tags []*Tag

// Keys returns a list of Placeholder key string
func (ts Tags) Keys() []string {
	keys := make([]string, len(ts))
	for i, t := range ts {
		keys[i] = t.Key()
	}
	return keys
}

// NewPlaceholder returns a new Tag's object
func NewPlaceholder() dataloader.Placeholder {
	return dataloader.Placeholder(&Tag{})
}

// Repository is an implemented of Tag's Repositorier interface
type Repository struct {
	db     mongodb.Database
	loader *dld.Loader
}

// NewRepository returns a new Tag's repository with dataloader configured
func NewRepository(db mongodb.Database) Repository {
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

// FindAll finds all Tags
func (repo Repository) FindAll(
	offset, limit int,
	orderBy struct {
		Field     string
		Direction string
	},
) ([]*Tag, error) {
	q := repo.db.C("tags").Find(nil).Select(bson.M{"_id": 1})
	q.Skip(offset)
	if limit > 0 {
		q.Limit(limit)
	}

	var ts Tags
	err := q.All(&ts)
	if err != nil {
		return nil, err
	}

	data, _ := repo.loader.LoadMany(context.TODO(), dld.NewKeysFromStrings(ts.Keys()))()
	for i := range ts {
		ts[i] = data[i].(*Tag)
	}

	return ts, nil
}
