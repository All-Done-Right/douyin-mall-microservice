package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/conf"
	auth "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v4"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// 解析 JWT 令牌
	if req.Token == "" {
		return &auth.VerifyResp{
			Res: false,
		}, nil
	}
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		klog.Info("jwt secret读取：" + conf.GetConf().JWT.Secret)
		return []byte(conf.GetConf().JWT.Secret), nil
	})
	// klog.Info(err)
	// klog.Info(token.Valid)
	if err != nil {
		return &auth.VerifyResp{
			Res: false,
		}, nil
	}

	// 检查令牌是否有效
	if token.Valid {
		return &auth.VerifyResp{
			Res: true,
		}, nil
	}

	return &auth.VerifyResp{
		Res: false,
	}, nil
}
