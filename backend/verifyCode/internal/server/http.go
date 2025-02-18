package server

import (
	v1 "verifyCode/api/helloworld/v1"
<<<<<<< HEAD
	"verifyCode/api/verifyCode"
=======
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	"verifyCode/internal/conf"
	"verifyCode/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

<<<<<<< HEAD
// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	greeter *service.GreeterService,
	verifyCodeService *service.VerifyCodeService,
	logger log.Logger) *http.Server {
=======
// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	v1.RegisterGreeterHTTPServer(srv, greeter)
<<<<<<< HEAD
	verifyCode.RegisterVerifyCodeHTTPServer(srv, verifyCodeService)
=======
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	return srv
}
