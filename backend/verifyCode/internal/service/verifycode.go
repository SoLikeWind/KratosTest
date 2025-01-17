package service

import (
	"context"
	"math/rand"
	pb "verifyCode/api/verifyCode" //pb protobuf的缩写
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

// GetVerifyCode 获取VerifyCode验证码
func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code:    RandCode(int(req.Length), req.Type),
		Success: "success",
	}, nil //实现业务逻辑的地方，
}

// RandCode （开放）生成随机字符，针对不同类型,先做一个不同类型差异
func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		return randCode("0123456789", l)
	case pb.TYPE_LETIER:
		return randCode("abcdefghijkmnopqrstuvwxyz", l)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijkmnopqrstuvwxyz", l)
	default:
	}
	return ""
}

// （内部）随机生成数字
func randCode(chars string, l int) string {
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		randIndex := rand.Intn(len(chars)) //返回0-n之间的随机数
		result[i] = chars[randIndex]
	}
	return string(result)
}
