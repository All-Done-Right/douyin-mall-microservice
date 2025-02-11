package service

import (
	"context"
	"time"

	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/conf"
	auth "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v4"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// 创建一个新的 JWT 令牌
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.UserId
	claims["exp"] = time.Now().Add(conf.JWTExpirationTime).Unix()

	// 生成签名后的令牌
	tokenString, err := token.SignedString([]byte(conf.JWTSecret))
	if err != nil {
		return nil, err
	}

	// 返回令牌响应
	return &auth.DeliveryResp{
		Token: tokenString,
	}, nil
}
