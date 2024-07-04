package service

import (
	"context"
	verifyCode "customer/api/verfiyCode"
	"customer/internal/biz"
	"customer/internal/data"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/consul/api"
	"log"
	"regexp"
	"time"

	pb "customer/api/customer"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	CD *data.CustomerData
	cb *biz.CustomerBiz
}

func NewCustomerService(cd *data.CustomerData, cb *biz.CustomerBiz) *CustomerService {
	return &CustomerService{
		CD: cd,
		cb: cb,
	}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	// 一，校验手机号
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$`
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(req.Telephone) {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "电话号码格式错误",
		}, nil
	}

	// 二，通验证码服务生成验证码（服务间通信，grpc）
	// 使用服务发现
	// 1.获取consul客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "localhost:8500"
	consulClient, err := api.NewClient(consulConfig)
	// 2.获取服务发现管理器
	dis := consul.New(consulClient)
	if err != nil {
		log.Fatal(err)
	}
	//selector.SetGlobalSelector(random.NewBuilder())
	selector.SetGlobalSelector(wrr.NewBuilder())
	//selector.SetGlobalSelector(p2c.NewBuilder())
	//log.Println(selector.GlobalSelector())
	// 2.1,连接目标grpc服务器
	endpoint := "discovery:///VerifyCode"
	conn, err := grpc.DialInsecure(
		//context.Background(),
		ctx,
		grpc.WithEndpoint(endpoint), // 目标服务的名字
		// 使用服务发现
		grpc.WithDiscovery(dis),
		//grpc.WithEndpoint("localhost:9000"), // verifyCode grpc service 地址
	)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码服务不可用",
		}, nil
	}
	//关闭
	defer func() {
		_ = conn.Close()
	}()

	// 2.2,发送获取验证码请求
	client := verifyCode.NewVerifyCodeClient(conn)
	reply, err := client.GetVerifyCode(context.Background(), &verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type:   1,
	})
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码获取错误",
		}, nil
	}

	// 三，redis的临时存储
	const life = 60
	if err := s.CD.SetVerifyCode(req.Telephone, reply.Code, life); err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码存储错误（Redis的操作错误）",
		}, nil
	}

	// 生成响应
	return &pb.GetVerifyCodeResp{
		Code:           0,
		VerifyCode:     reply.Code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: life,
	}, nil
}

// 登录的application
func (s *CustomerService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	// 一，校验电话和验证码
	// redis
	code := s.CD.GetVerifyCode(req.Telephone)
	if code == "" || code != req.VerifyCode {
		return &pb.LoginResp{
			Code:    1,
			Message: "验证码不匹配",
		}, nil
	}

	// 二，判定电话号码是否注册，来获取顾客信息
	customer, err := s.CD.GetCustomerByTelephone(req.Telephone)
	if err != nil {
		return &pb.LoginResp{
			Code:    1,
			Message: "顾客信息获取错误",
		}, nil
	}

	// 三，设置token，jwt-token
	token, err := s.CD.GenerateTokenAndSave(customer, biz.CustomerDuration*time.Second, biz.CustomerSecret)
	log.Println(err)
	if err != nil {
		return &pb.LoginResp{
			Code:    1,
			Message: "Token生成失败",
		}, nil
	}

	// 四，响应token
	return &pb.LoginResp{
		Code:          0,
		Message:       "login success",
		Token:         token,
		TokenCreateAt: time.Now().Unix(),
		TokenLife:     biz.CustomerDuration,
	}, nil
}

// 顾客退出
func (s *CustomerService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
	// 一，获取用户id
	// 获取，断言使用
	claims, _ := jwt.FromContext(ctx)
	claimsMap := claims.(jwtv4.MapClaims)

	// 二，删除用户的token
	if err := s.CD.DelToken(claimsMap["jti"]); err != nil {
		return &pb.LogoutResp{
			Code:    1,
			Message: "Token删除失败",
		}, nil
	}

	// 三，成功，响应
	return &pb.LogoutResp{
		Code:    0,
		Message: "logout success",
	}, nil
}

func (s *CustomerService) EstimatePrice(ctx context.Context, req *pb.EstimatePriceReq) (*pb.EstimatePriceResp, error) {
	price, err := s.cb.GetEstimatePrice(ctx, req.Origin, req.Destination)
	if err != nil {
		return &pb.EstimatePriceResp{
			Code:    1,
			Message: err.Error(),
		}, nil
	}

	return &pb.EstimatePriceResp{
		Code:        0,
		Message:     "SUCCESS",
		Origin:      req.Origin,
		Destination: req.Destination,
		Price:       price,
	}, nil
}
