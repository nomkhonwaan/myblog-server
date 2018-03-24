package mongodb

import mgo "gopkg.in/mgo.v2"

// Collection is a wrapper interface of mgo.Collection that allows to mock for testing
type Collection interface {
	Find(query interface{}) Query
}

type wrappedCollection struct {
	*mgo.Collection
}

func (w *wrappedCollection) Find(query interface{}) Query {
	return &wrappedQuery{Query: w.Collection.Find(query)}
}
