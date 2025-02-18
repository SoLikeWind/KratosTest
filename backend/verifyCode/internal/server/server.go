package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
<<<<<<< HEAD

//提供GRPC和HTTP服务，再去server目录中实现具体服务（http和grpc实例的创建和配置）
=======
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
