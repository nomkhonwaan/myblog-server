package post

import (
	"context"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repositorier interface {
	FindByID(id string) (*Post, error)
	FindPublishedPostByID(id string) (*Post, error)
	FindAll(
		offset, limit *int32,
		orderBy *struct {
			Direction *string
			Field     *string
		},
	) ([]*Post, error)
	FindAllPublishedPosts(
		offset, limit *int32,
		orderBy *struct {
			Direction *string
			Field     *string
		},
	) ([]*Post, error)
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
			dataloader.NewBatchFunc(db.C("posts"), NewPlaceholder),
			dld.WithCache(c),
		),
	}
}

func (repo *Repository) FindByID(id string) (*Post, error) {
	p, err := repo.loader.Load(context.TODO(), dld.StringKey(id))()
	if err != nil {
		return nil, err
	}

	return p.(*Post), err
}

func (repo *Repository) FindPublishedPostByID(id string) (*Post, error) {
	p, err := repo.loader.Load(context.TODO(), dld.StringKey(id))()
	if err != nil {
		return nil, err
	}

	if p.(*Post).IsPublished() {
		return p.(*Post), nil
	}

	return nil, nil
}

func (repo *Repository) FindAll(
	offset, limit *int32,
	orderBy *struct {
		Direction *string
		Field     *string
	},
) ([]*Post, error) {
	sort := "-publishedAt"

	if orderBy != nil {
		if orderBy.Field != nil {
			sort = OrderField(*orderBy.Field).ToCamelCase()
		}

		if orderBy.Direction != nil && *orderBy.Direction == "DESC" {
			sort = "-" + sort
		}
	}

	q := repo.db.C("posts").Find(nil).Sort(sort).Select(bson.M{"_id": 1})
	if offset != nil {
		q.Skip(int(*offset))
	}
	if limit != nil {
		q.Limit(int(*limit))
	}

	iter := q.Iter()
	var p struct {
		ID bson.ObjectId `bson:"_id"`
	}

	keys := make([]string, 0)
	for iter.Next(&p) {
		keys = append(keys, p.ID.Hex())
	}

	ps, errs := repo.loader.LoadMany(context.TODO(), dld.NewKeysFromStrings(keys))()
	if errs != nil {
		return nil, errs[0]
	}

	assertedPosts := make([]*Post, len(ps))
	for i, p := range ps {
		assertedPosts[i] = p.(*Post)
	}

	return assertedPosts, nil
}

func (repo *Repository) FindAllPublishedPosts(
	offset, limit *int32,
	orderBy *struct {
		Direction *string
		Field     *string
	},
) ([]*Post, error) {
	ps, err := repo.FindAll(offset, limit, orderBy)
	if err != nil {
		return nil, err
	}

	publishedPost := make([]*Post, 0)
	for _, p := range ps {
		if p.IsPublished() {
			publishedPost = append(publishedPost, p)
		}
	}

	return publishedPost, nil
}
