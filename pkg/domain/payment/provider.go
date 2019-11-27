package payment

import "github.com/google/wire"

var Set = wire.NewSet(NewUseCase, NewInMemoryRepository, NewMongodbRepository,
	//wire.Bind(new(Repository), new(*InMemoryRepository)), // in InMemoryRepository
	wire.Bind(new(Repository), new(*MongodbRepository)),    // in MongoDB
	wire.Bind(new(UseCase), new(*PaymentUseCase)))

