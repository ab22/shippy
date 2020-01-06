package db

import (
	mgo "gopkg.in/mgo.v2"
)

// NewMongoClient --
func NewMongoClient(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
