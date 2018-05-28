package post

import (
	"context"
	"time"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	"github.com/nomkhonwaan/myblog-server/pkg/mongodb"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
	"gopkg.in/mgo.v2/bson"
)

type Repositorier interface {
	FindByID(id string) (*Post, error)
	FindPublishedByID(id string) (*Post, error)
	FindAllPublished(
		offset, limit int,
		orderBy struct {
			Field     string
			Direction string
		},
	) ([]*Post, error)
}

type Post struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string        `bson:"title"`
	Slug        string        `bson:"slug"`
	Status      string        `bson:"status"`
	HTML        string        `bson:"html"`
	Markdown    string        `bson:"markdown"`
	Tags        []*tag.Tag    `bson:"tags"`
	CreatedAt   time.Time     `bson:"createdAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
	PublishedAt time.Time     `bson:"publishedAt"`
}

func (p Post) Key() string {
	return p.ID.Hex()
}

type Posts []*Post

func (ps Posts) Keys() []string {
	keys := make([]string, len(ps))
	for i, p := range ps {
		keys[i] = p.Key()
	}
	return keys
}

func NewPlaceholder() dataloader.Placeholder {
	return dataloader.Placeholder(&Post{})
}

type Repository struct {
	db     mongodb.Database
	loader dld.Interface
}

func NewRepository(db mongodb.Database) Repository {
	c, _ := cache.New(100)

	return Repository{
		db: db,
		loader: dld.NewBatchedLoader(
			dataloader.NewBatchFunc(db.C("posts"), NewPlaceholder),
			dld.WithCache(c),
		),
	}
}

func (repo Repository) FindByID(id string) (*Post, error) {
	p, err := repo.loader.Load(context.TODO(), dld.StringKey(id))()
	if err != nil {
		return nil, err
	}
	if p != nil {
		return p.(*Post), nil
	}
	return nil, nil
}

func (repo Repository) FindPublishedByID(id string) (*Post, error) {
	p, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if p.Status != "PUBLISHED" && p.PublishedAt.IsZero() {
		return nil, nil
	}

	return p, nil
}

func (repo Repository) FindAllPublished(
	offset, limit int,
	orderBy struct {
		Field     string
		Direction string
	},
) ([]*Post, error) {
	return nil, nil
}
