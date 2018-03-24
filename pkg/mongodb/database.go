package mongodb

import mgo "gopkg.in/mgo.v2"

// Database is a wrapper interface of mgo.Database that allows to mock for testing
type Database interface {
	C(name string) Collection
}

type wrappedDatabase struct {
	*mgo.Database
}

func (w *wrappedDatabase) C(name string) Collection {
	return &wrappedCollection{Collection: w.Database.C(name)}
}
