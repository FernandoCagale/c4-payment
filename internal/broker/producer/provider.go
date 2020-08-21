package producer

import "github.com/google/wire"

var Set = wire.NewSet(NewProducerKafka, wire.Bind(new(Producer), new(*ProducerKafka)))
