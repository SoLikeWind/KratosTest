package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
<<<<<<< HEAD
var ProviderSet = wire.NewSet(NewGreeterUsecase)
=======
var ProviderSet = wire.NewSet(NewGreeterUsecase, NewCustomerBiz)
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
