package tag

import (
	"context"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repositorier interface {
	FindByID(id string) (*Tag, error)
	FindAll(
		offset, limit *int32,
		orderBy *struct {
			Direction *string
			Field     *string
		},
	) ([]*Tag, error)
}

type Repository struct {
	db     *mgo.Database
	loader *dld.Loader
}

func NewRepository(db *mgo.Database) *Repository {
	c, _ := cache.New(100)

	return &Repository{
		db: db,
		loader: dld.NewBatchedLoader(
			dataloader.NewBatchFunc(db.C("tags"), NewPlaceholder),
			dld.WithCache(c),
		),
	}
}

func (repo *Repository) FindByID(id string) (*Tag, error) {
	t, err := repo.loader.Load(context.TODO(), dld.StringKey(id))()
	if err != nil {
		return nil, err
	}

	return t.(*Tag), err
}

func (repo *Repository) FindAll(
	offset, limit *int32,
	orderBy *struct {
		Direction *string
		Field     *string
	},
) ([]*Tag, error) {
	q := repo.db.C("tags").Find(nil).Select(bson.M{"_id": 1})
	if offset != nil {
		q.Skip(int(*offset))
	}
	if limit != nil {
		q.Limit(int(*limit))
	}

	iter := q.Iter()
	var t struct {
		ID bson.ObjectId `bson:"_id"`
	}

	keys := make([]string, 0)
	for iter.Next(&t) {
		keys = append(keys, t.ID.Hex())
	}

	ts, errs := repo.loader.LoadMany(context.TODO(), dld.NewKeysFromStrings(keys))()
	if errs != nil {
		return nil, errs[0]
	}

	assertedTags := make([]*Tag, len(ts))
	for i, p := range ts {
		assertedTags[i] = p.(*Tag)
	}

	return assertedTags, nil
}
