package dataloader

import (
	"context"

	"github.com/nicksrandall/dataloader"
	"github.com/nomkhonwaan/myblog-server/pkg/mongodb"
	"gopkg.in/mgo.v2/bson"
)

// Placeholder is an interface of any models that can resolve a key as a string
type Placeholder interface {
	Key() string
}

// NewBatchFunc returns a new dataloader.BatchFunc
func NewBatchFunc(c mongodb.Collection, newPlaceholderFn func() Placeholder) dataloader.BatchFunc {
	return func(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		keyValuePairs := make(map[string]*dataloader.Result)

		q := bson.M{"_id": bson.M{"$in": make([]bson.ObjectId, len(keys))}}
		for i, key := range keys {
			q["_id"].(bson.M)["$in"].([]bson.ObjectId)[i] = bson.ObjectIdHex(key.String())
		}

		placeholder := newPlaceholderFn()
		iter := c.Find(q).Iter()
		for iter.Next(placeholder) {
			keyValuePairs[placeholder.Key()] = &dataloader.Result{
				Data:  placeholder,
				Error: nil,
			}
			placeholder = newPlaceholderFn()
		}

		for i, key := range keys {
			if val, ok := keyValuePairs[key.String()]; ok {
				results[i] = val
			}
		}

		return results
	}
}
