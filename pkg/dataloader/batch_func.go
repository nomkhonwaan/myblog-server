package dataloader

import (
	"context"

	"github.com/nicksrandall/dataloader"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Placeholder is an interface of any models that can resolve a key as a string
type Placeholder interface {
	Key() string
}

// NewBatchFunc returns a new dataloader.BatchFunc
func NewBatchFunc(c *mgo.Collection, newPlaceholderFunc func() Placeholder) dataloader.BatchFunc {
	return func(_ context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		keyValuePairs := make(map[string]*dataloader.Result)

		q := bson.M{"_id": bson.M{"$in": make([]bson.ObjectId, len(keys))}}
		for i, key := range keys {
			q["_id"].(bson.M)["$in"].([]bson.ObjectId)[i] = bson.ObjectIdHex(key.String())
		}

		placeholder := newPlaceholderFunc()
		iter := c.Find(q).Iter()
		for iter.Next(placeholder) {
			keyValuePairs[placeholder.Key()] = &dataloader.Result{
				Data:  placeholder,
				Error: nil,
			}
			placeholder = newPlaceholderFunc()
		}

		for i, key := range keys {
			if val, ok := keyValuePairs[key.String()]; ok {
				results[i] = val
			}
		}

		return results
	}
}
