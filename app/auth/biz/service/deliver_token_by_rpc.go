package service

import (
	"context"
	"errors"
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
	// 如果用户 ID 为空，返回错误
	if req.UserId == 0 {
		return nil, errors.New("Userid is empty")
	}

	claims := jwt.MapClaims{
		"user_id": req.UserId,
		"exp":     time.Now().Add(time.Hour * time.Duration(conf.GetConf().JWT.ExpireTime)).Unix(), // 过期时间
		"iat":     time.Now().Unix(),                                                               // 创建时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.GetConf().JWT.Secret))
	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{Token: tokenString}, nil
}
