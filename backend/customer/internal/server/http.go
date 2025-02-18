package server

import (
<<<<<<< HEAD
	"customer/api/customer"
	v1 "customer/api/helloworld/v1"
	"customer/internal/conf"
	"customer/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	customerService *service.CustomerService,
	greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
=======
	"context"
	"customer/api/customer"
	v1 "customer/api/helloworld/v1"
	"customer/internal/biz"
	"customer/internal/conf"
	"customer/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, customerService *service.CustomerService, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			// 自己设置的中间件
			// CORS，全部的请求（响应）都使用该中间件
			selector.Server(MWCors()).Match(func(ctx context.Context, operation string) bool {
				return true
			}).Build(),
			// JWT，除了特定几个，其他请求（响应）都是用该中间件
			selector.Server(
				jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
					return []byte(biz.CustomerSecret), nil
				}),
				customerJWT(customerService),
			).Match(func(ctx context.Context, operation string) bool {
				// 根据自己的需要完成是否启用该中间件的校验
				noJWT := map[string]struct{}{
					"/api.customer.Customer/Login":         {},
					"/api.customer.Customer/GetVerifyCode": {},
				}
				if _, exists := noJWT[operation]; exists {
					return false
				}
				return true
			}).Build(),
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
<<<<<<< HEAD
	v1.RegisterGreeterHTTPServer(srv, greeter)
	customer.RegisterCustomerHTTPServer(srv, customerService)
=======
	// 注册customer的http服务
	customer.RegisterCustomerHTTPServer(srv, customerService)
	v1.RegisterGreeterHTTPServer(srv, greeter)
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	return srv
}
