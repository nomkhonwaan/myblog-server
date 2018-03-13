package post

import (
	"github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/dataloader/cache"
	mgo "gopkg.in/mgo.v2"
)

type Repositorier interface {
	FindByID(id string) (*Post, error)
	FindAll(
		offset, limit int,
		orderBy *struct {
			Direction *string
			Field     *string
		},
	) ([]*Post, error)
}

type Repository struct {
	db                *mgo.Database
	loader            *dataloader.Loader
	collectionName    string
	batchQueryResults []*Post
}

func NewRepository(db *mgo.Database) (*Repository, error) {
	c, err := cache.New(100)
	if err != nil {
		return nil, err
	}

	repo := Repository{db: db, collectionName: "posts"}
	repo.loader = dataloader.NewBatchedLoader(
		repo.batch,
		dataloader.WithCache(c),
	)

	return &repo, nil
}

// func (repo *Repository) batch(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
// 	results := make([]*dataloader.Result, len(keys))

// 	q := bson.M{"$or": make([]bson.M, len(keys))}
// 	for i, key := range keys {
// 		q["$or"].([]bson.M)[i] = bson.M{"_id": bson.ObjectIdHex(key.String())}
// 	}

// 	err := repo.db.C(repo.collectionName).Find(q).All(&repo.batchQueryResults)

// 	keyMap := make(map[string]interface{})
// 	for _, val := range repo.batchQueryResults {
// 		keyMap[val.ID()]
// 	}

// 	// results := make([]*dataloader.Result, len(keys))

// 	// q := bson.M{"$or": make([]bson.M, 0)}
// 	// for _, key := range keys {
// 	// 	q["$or"] = append(q["$or"].([]bson.M), bson.M{"_id": bson.ObjectIdHex(key.String())})
// 	// }

// 	// var posts []*Post
// 	// err := repo.db.C("posts").Find(q).All(&posts)

// 	// keyMap := make(map[string]*Post)
// 	// for _, post := range posts {
// 	// 	keyMap[post.ID.Hex()] = post
// 	// }

// 	// for i, key := range keys {
// 	// 	if post, ok := keyMap[key.String()]; ok {
// 	// 		results[i] = &dataloader.Result{
// 	// 			Data:  post,
// 	// 			Error: err,
// 	// 		}
// 	// 	} else {
// 	// 		results[i] = nil
// 	// 	}
// 	// }

// 	// return results
// }

// func (repo *Repository) FindByID(id string) (*Post, error) {
// 	post, err := repo.loader.Load(context.TODO(), dataloader.StringKey(id))()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return post.(*Post), nil
// }

// func (repo *Repository) FindAll(
// 	orderBy *struct {
// 		Direction *string
// 		Field     *string
// 	},
// ) ([]*Post, error) {
// 	var sort string

// 	if orderBy == nil {
// 		sort = "-publishedAt"
// 	} else {
// 		direction := ""
// 		field := "publishedAt"

// 		if orderBy.Direction != nil && *orderBy.Direction == "DESC" {
// 			direction = "-"
// 		}

// 		if orderBy.Field != nil {
// 			switch *orderBy.Field {
// 			case "SLUG":
// 				field = "slug"
// 				break
// 			case "CREATED_AT":
// 				field = "createdAt"
// 				break
// 			case "UPDATED_AT":
// 				field = "updatedAt"
// 				break
// 			case "PUBLISHED_AT":
// 				field = "publishedAt"
// 				break
// 			}
// 		}

// 		sort = direction + field
// 	}

// 	var ids []struct {
// 		ID bson.ObjectId `bson:"_id"`
// 	}
// 	err := repo.db.C("posts").Find(nil).Sort(sort).Select(bson.M{"_id": 1}).All(&ids)
// 	if err != nil {
// 		return nil, err
// 	}

// 	keys := make([]string, len(ids))
// 	for i, val := range ids {
// 		keys[i] = val.ID.Hex()
// 	}

// 	posts, errs := repo.loader.LoadMany(context.TODO(), dataloader.NewKeysFromStrings(keys))()
// 	if errs != nil {
// 		return nil, errs[0]
// 	}

// 	assertedPosts := make([]*Post, len(posts))
// 	for i, post := range posts {
// 		assertedPosts[i] = post.(*Post)
// 	}

// 	return assertedPosts, nil
// }
