package post

import (
	"context"
	"time"

	dld "github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	"github.com/nomkhonwaan/myblog-server/pkg/tag"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Post represent an entity of Post
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

// Key returns a Placeholder key string
func (p *Post) Key() string {
	return p.ID.Hex()
}

// NewPlaceholder returns a new Post's object
func NewPlaceholder() dataloader.Placeholder {
	return dataloader.Placeholder(&Post{})
}

// Repositorier is a Post's repository interface
type Repositorier interface {
	FindByID(id string) (*Post, error)
	FindPublishedByID(id string) (*Post, error)
}

// Repository is an implemented of Post's Repositorier interface
type Repository struct {
	db     *mgo.Database
	loader *dld.Loader
}

// NewRepository returns a new Post's repository with dataloader configured
func NewRepository(db *mgo.Database) Repository {
	c, _ := cache.New(100)

	return Repository{
		db: db,
		loader: dld.NewBatchedLoader(
			dataloader.NewBatchFunc(db.C("posts"), NewPlaceholder),
			dld.WithCache(c),
		),
	}
}

// FindByID finds a single Post from its ID
func (repo Repository) FindByID(id string) (*Post, error) {
	p, err := repo.loader.Load(context.TODO(), dld.StringKey(id))()
	if err != nil {
		return nil, err
	}

	return p.(*Post), nil
}

// FindPublishedByID finds a single published Post from its ID
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
