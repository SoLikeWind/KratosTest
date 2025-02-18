package service

import "github.com/google/wire"

// ProviderSet is service providers.
<<<<<<< HEAD
var ProviderSet = wire.NewSet(NewGreeterService, NewCustomerService)
=======
var ProviderSet = wire.NewSet(NewCustomerService, NewGreeterService)
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
