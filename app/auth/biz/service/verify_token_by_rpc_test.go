package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	auth "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"

	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/conf"
)

// 初始化测试配置
//
//	func setupTestConfig() func() {
//		originalConf := conf.GetConf()
//		testConf := &conf.Config{
//			JWT: conf.JWTConfig{
//				Secret: "test_secret",
//			},
//		}
//		conf.SetConf(testConf)
//		return func() { conf.SetConf(originalConf) }
//	}
func tamperToken(token string) string {
	// 分割 JWT 的三个部分
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		panic("invalid token format")
	}

	// 解码 Payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		panic("failed to decode payload")
	}

	// 解析 Payload 为 map
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		panic("failed to unmarshal payload")
	}

	// 篡改 Payload
	claims["user_id"] = 456 // 修改 user_id
	claims["admin"] = true  // 添加新的字段

	// 重新编码 Payload
	tamperedPayload, err := json.Marshal(claims)
	if err != nil {
		panic("failed to marshal tampered payload")
	}
	parts[1] = base64.RawURLEncoding.EncodeToString(tamperedPayload)

	// 重新拼接篡改后的令牌
	tamperedToken := strings.Join(parts, ".")
	return tamperedToken
}

func generateToken(secret string, expiresAt time.Time) string {
	claims := jwt.MapClaims{
		"user_id": 1003,
		"exp":     expiresAt.Unix(),  // 过期时间
		"iat":     time.Now().Unix(), // 创建时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		klog.Error("生成token失败")
		return ""
	}
	return tokenString
}

func TestVerifyTokenByRPCService_Run(t *testing.T) {
	ctx := context.Background()
	service := NewVerifyTokenByRPCService(ctx)

	now := time.Now()
	validToken := generateToken(conf.GetConf().JWT.Secret, now.Add(1*time.Hour))
	expiredToken := generateToken(conf.GetConf().JWT.Secret, now.Add(-1*time.Hour))
	wrongSigToken := tamperToken(validToken)

	testCases := []struct {
		name        string
		token       string
		exceptans   bool
		description string
	}{
		{
			name:        "valid_token",
			token:       validToken,
			exceptans:   true,
			description: "有效JWT令牌应验证通过",
		},
		{
			name:        "invalid_signature",
			token:       wrongSigToken,
			exceptans:   false,
			description: "错误签名的令牌应验证失败",
		},
		{
			name:        "expired_token",
			token:       expiredToken,
			exceptans:   false,
			description: "过期令牌应验证失败",
		},
		{
			name:        "empty_token",
			token:       "",
			exceptans:   false,
			description: "空令牌应验证失败",
		},
		{
			name:        "malformed_token",
			token:       "invalid.token.string",
			exceptans:   false,
			description: "格式错误的令牌应验证失败",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := &auth.VerifyTokenReq{Token: tt.token}
			resp, err := service.Run(req)
			fmt.Println(err)

			assert.NoError(t, err, "错误应为nil")
			assert.Equal(t, tt.exceptans, resp.Res, tt.description)
		})
	}
}

func testCasesecretConfiguration(t *testing.T) {
	// 验证配置读取是否正确
	assert.Equal(t, "test_secret", conf.GetConf().JWT.Secret,
		"测试配置应正确设置JWT密钥")
}
