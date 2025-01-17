// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.29.0
// source: api/verifyCode/verifyCode.proto

package verifyCode

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationVerifyCodeGetVerifyCode = "/api.verifyCode.VerifyCode/GetVerifyCode"

type VerifyCodeHTTPServer interface {
	// GetVerifyCode 定义 VerifyCode 服务
	GetVerifyCode(context.Context, *GetVerifyCodeRequest) (*GetVerifyCodeReply, error)
}

func RegisterVerifyCodeHTTPServer(s *http.Server, srv VerifyCodeHTTPServer) {
	r := s.Route("/")
	r.POST("/verifycode", _VerifyCode_GetVerifyCode0_HTTP_Handler(srv))
}

func _VerifyCode_GetVerifyCode0_HTTP_Handler(srv VerifyCodeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVerifyCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationVerifyCodeGetVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVerifyCode(ctx, req.(*GetVerifyCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVerifyCodeReply)
		return ctx.Result(200, reply)
	}
}

type VerifyCodeHTTPClient interface {
	GetVerifyCode(ctx context.Context, req *GetVerifyCodeRequest, opts ...http.CallOption) (rsp *GetVerifyCodeReply, err error)
}

type VerifyCodeHTTPClientImpl struct {
	cc *http.Client
}

func NewVerifyCodeHTTPClient(client *http.Client) VerifyCodeHTTPClient {
	return &VerifyCodeHTTPClientImpl{client}
}

func (c *VerifyCodeHTTPClientImpl) GetVerifyCode(ctx context.Context, in *GetVerifyCodeRequest, opts ...http.CallOption) (*GetVerifyCodeReply, error) {
	var out GetVerifyCodeReply
	pattern := "/verifycode"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationVerifyCodeGetVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
