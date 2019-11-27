package datastore

import "github.com/google/wire"

var Set = wire.NewSet(NewMongoDB)
