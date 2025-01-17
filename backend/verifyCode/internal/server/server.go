package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

//提供GRPC和HTTP服务，再去server目录中实现具体服务（http和grpc实例的创建和配置）
