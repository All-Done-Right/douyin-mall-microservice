package service

import (
	"context"
	"testing"
	"time"

	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/conf"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestRenewTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRenewTokenByRPCService(ctx)

	// 定义测试用例
	testCases := []struct {
		name        string
		req         *auth.RenewTokenReq
		expectError bool
		errorMsg    string
		setup       func() *auth.RenewTokenReq
	}{
		{
			name: "正常情况",
			setup: func() *auth.RenewTokenReq {
				oldTokenString := generateToken(conf.GetConf().JWT.Secret, time.Now().Add(time.Hour))
				return &auth.RenewTokenReq{OldToken: oldTokenString}
			},
			expectError: false,
		},
		{
			name: "缺失 Token",
			setup: func() *auth.RenewTokenReq {
				return &auth.RenewTokenReq{OldToken: ""}
			},
			expectError: true,
			errorMsg:    "missing token in request",
		},
		{
			name: "无效的 Token",
			setup: func() *auth.RenewTokenReq {
				return &auth.RenewTokenReq{OldToken: "invalid-token"}
			},
			expectError: true,
			errorMsg:    "token contains an invalid number of segments",
		},
		{
			name: "Token 过期",
			setup: func() *auth.RenewTokenReq {
				oldClaims := jwt.MapClaims{
					"user_id": 123,
					"exp":     time.Now().Add(-time.Hour).Unix(), // 过期时间设置为过去
					"iat":     time.Now().Unix(),
				}
				oldToken := jwt.NewWithClaims(jwt.SigningMethodHS256, oldClaims)
				oldTokenString, _ := oldToken.SignedString([]byte(conf.GetConf().JWT.Secret))
				return &auth.RenewTokenReq{OldToken: oldTokenString}
			},
			expectError: true,
			errorMsg:    "Token is expired",
		},
	}

	// 遍历测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.setup()
			resp, err := s.Run(req)

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, tc.errorMsg, err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Token)
			}
		})
	}
}
