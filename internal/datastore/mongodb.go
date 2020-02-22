package datastore

import (
	"gopkg.in/mgo.v2"
	"os"
)

func NewMongoDB() (*mgo.Session, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{os.Getenv("MONGO_ADDRS")},
		FailFast: true,
		Database: os.Getenv("MONGO_DATABASE"),
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
