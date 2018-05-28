package tag

import (
	"context"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	"github.com/nomkhonwaan/myblog-server/pkg/mongodb"
	"gopkg.in/mgo.v2/bson"
)

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

type Tag struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Slug string        `bson:"slug"`
}

func (t Tag) Key() string {
	return t.ID.Hex()
}

type Tags []*Tag

func (ts Tags) Keys() []string {
	keys := make([]string, len(ts))
	for i, t := range ts {
		keys[i] = t.Key()
	}
	return keys
}

func NewPlaceholder() dataloader.Placeholder {
	return dataloader.Placeholder(&Tag{})
}

type Repository struct {
	Database mongodb.Database
	Loader   dld.Interface
}

func NewRepository(db mongodb.Database) Repository {
	c, _ := cache.New(100)

	return Repository{
		Database: db,
		Loader: dld.NewBatchedLoader(
			dataloader.NewBatchFunc(db.C("tags"), NewPlaceholder),
			dld.WithCache(c),
		),
	}
}

func (repo Repository) FindByID(id string) (*Tag, error) {
	t, err := repo.Loader.Load(context.TODO(), dld.StringKey(id))()
	if err != nil {
		return nil, err
	}
	if t != nil {
		return t.(*Tag), nil
	}
	return nil, nil
}

func (repo Repository) FindAll(
	offset, limit int,
	orderBy struct {
		Field     string
		Direction string
	},
) ([]*Tag, error) {
	q := repo.Database.C("tags").Find(nil).Select(bson.M{"_id": 1})
	q.Skip(offset)
	if limit > 0 {
		q.Limit(limit)
	}

	var ts Tags
	err := q.All(&ts)
	if err != nil {
		return nil, err
	}

	data, _ := repo.Loader.LoadMany(context.TODO(), dld.NewKeysFromStrings(ts.Keys()))()
	for i := range ts {
		ts[i] = data[i].(*Tag)
	}

	return ts, nil
}
