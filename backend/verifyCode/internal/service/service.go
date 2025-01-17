package service

//每当service.go中代码发生改变(业务增减)就生成代码(go generate ./...)
//因为是引入了wire依赖注入
import "github.com/google/wire"

// ProviderSet is service providers.（它是服务提供者）
var ProviderSet = wire.NewSet(NewGreeterService, NewVerifyCodeService)

//将verifyCode.go中的NewVerifyCodeService构建函数放到service.go注册服务，
//该语句有一套构建注册服务的方法。
