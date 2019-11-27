package datastore

import (
	"gopkg.in/mgo.v2"
	"os"
)

func NewMongoDB() (*mgo.Session, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{os.Getenv("MONGO_ADDRS")},
		FailFast: true,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
