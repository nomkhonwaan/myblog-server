package mongodb

import (
	mgo "gopkg.in/mgo.v2"
)

// Session is a wrapper interface of mgo.Session that allows to mock for testing
type Session interface {
	DB(name string) Database
	Clone() Session
	Close()
}

// DialWithInfo is a wrapper function of mgo.DialWithInfo
// which returns a Session interface that allows to mock for testing
func DialWithInfo(dialInfo *mgo.DialInfo) (Session, error) {
	session, err := mgo.DialWithInfo(dialInfo)
	return &wrappedSession{Session: session}, err
}

type wrappedSession struct {
	*mgo.Session
}

func (w *wrappedSession) DB(name string) Database {
	return &wrappedDatabase{Database: w.Session.DB(name)}
}

func (w *wrappedSession) Clone() Session {
	// logrus.Info(w.Session.Clone())
	return w
	// return &wrappedSession{Session: w.Session.Clone()}
}

func (w *wrappedSession) Close() {
	w.Session.Close()
}
