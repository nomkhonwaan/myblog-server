package mongodb

import mgo "gopkg.in/mgo.v2"

// Query is a wrapper interface of mgo.Query that allows to mock for testing
type Query interface {
	All(result interface{}) error
	Iter() Iter
	Limit(n int) Query
	Select(selector interface{}) Query
	Skip(n int) Query
}

type wrappedQuery struct {
	*mgo.Query
}

func (w *wrappedQuery) All(result interface{}) error {
	return w.Query.All(result)
}

func (w *wrappedQuery) Iter() Iter {
	return &wrappedIter{Iter: w.Query.Iter()}
}

func (w *wrappedQuery) Limit(n int) Query {
	w.Query.Limit(n)
	return w
}

func (w *wrappedQuery) Select(selector interface{}) Query {
	w.Query.Select(selector)
	return w
}

func (w *wrappedQuery) Skip(n int) Query {
	w.Query.Skip(n)
	return w
}
