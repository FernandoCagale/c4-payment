package routers

import "github.com/google/wire"

var Set = wire.NewSet(NewSystem)
