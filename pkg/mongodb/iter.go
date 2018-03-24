package mongodb

import (
	"gopkg.in/mgo.v2"
)

// Iter is a wrapper interface of mgo.Iter that allows to mock for testing
type Iter interface {
	Next(result interface{}) bool
}

type wrappedIter struct {
	*mgo.Iter
}

func (w *wrappedIter) Next(result interface{}) bool {
	return w.Iter.Next(result)
}
