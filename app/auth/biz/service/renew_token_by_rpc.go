package service

import (
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/conf"
	auth "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v4"
)

type RenewTokenByRPCService struct {
	ctx context.Context
} // NewRenewTokenByRPCService new RenewTokenByRPCService
func NewRenewTokenByRPCService(ctx context.Context) *RenewTokenByRPCService {
	return &RenewTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *RenewTokenByRPCService) Run(req *auth.RenewTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	// 1. 从请求中获取旧的 Token
	// 1. 从请求中获取旧的 Token
	oldTokenStr := req.OldToken
	if oldTokenStr == "" {
		return nil, errors.New("missing token in request")
	}

	// 2. 解析旧的 JWT Token
	token, err := jwt.Parse(oldTokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.GetConf().JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	// 3. 验证旧的 JWT Token
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// 4. 提取旧 Token 中的用户 ID
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to extract claims from token")
	}
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("failed to extract user_id from claims")
	}

	// 5. 生成新的 JWT Token
	newClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(conf.GetConf().JWT.ExpireTime)).Unix(), // 过期时间
		"iat":     time.Now().Unix(),                                                               // 创建时间
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newTokenString, err := newToken.SignedString([]byte(conf.GetConf().JWT.Secret))
	if err != nil {
		return nil, err
	}

	// 6. 返回新的 JWT Token
	return &auth.DeliveryResp{Token: newTokenString}, nil

}
