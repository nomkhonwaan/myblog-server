package post

import (
	"context"

	"github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repositorier interface {
	FindByID(id string) (*Post, error)
}

type Repository struct {
	db     *mgo.Database
	loader *dataloader.Loader
}

func NewRepository(db *mgo.Database) (*Repository, error) {
	c, err := cache.New(100)
	if err != nil {
		return nil, err
	}

	repo := Repository{db: db}
	repo.loader = dataloader.NewBatchedLoader(
		repo.batch,
		dataloader.WithCache(c),
	)

	return &repo, nil
}

func (repo *Repository) batch(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))

	q := bson.M{"$or": make([]bson.M, 0)}
	for _, key := range keys {
		q["$or"] = append(q["$or"].([]bson.M), bson.M{"_id": bson.ObjectIdHex(key.String())})
	}

	var posts []*Post
	err := repo.db.C("posts").Find(q).All(&posts)

	for i, post := range posts {
		results[i] = &dataloader.Result{
			Data:  post,
			Error: err,
		}
	}

	return results
}

func (repo *Repository) FindByID(id string) (*Post, error) {
	post, err := repo.loader.Load(context.TODO(), dataloader.StringKey(id))()
	if err != nil {
		return nil, err
	}

	return post.(*Post), nil
}
