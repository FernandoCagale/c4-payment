package consumer

import "github.com/google/wire"

var Set = wire.NewSet(NewConsumerKafka, wire.Bind(new(Consumer), new(*ConsumerKafka)))
