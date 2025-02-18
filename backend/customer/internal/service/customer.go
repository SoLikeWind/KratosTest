package service

import (
	"context"
	"regexp"

	pb "customer/api/customer"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	//1.校验手机号
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$` //手机号的标准正则表达式
	regexpPattern := regexp.MustCompile(pattern)                                        //校验，MustCompile类似于[Compile]，但如果表达式不能被解析则会panic。
	if !regexpPattern.MatchString(req.Telephone) {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "电话号码格式错误",
		}, nil
	}
	return &pb.GetVerifyCodeResp{}, nil
}
