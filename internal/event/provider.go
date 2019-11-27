package event

import "github.com/google/wire"

var Set = wire.NewSet(New, wire.Bind(new(Event), new(*EventRepository)))
