package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/conf"
	auth "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

// Mock 时间函数和配置变量
var now = time.Now

func TestDeliverTokenByRPCService_Run(t *testing.T) {
	defer func() {
		now = time.Now
	}()

	testCases := []struct {
		name           string
		userID         int64
		mockSecret     string
		mockExpiration time.Duration
		mockTime       time.Time
		expectedError  bool
		expectedErrMsg string
		validateResp   func(*testing.T, *auth.DeliveryResp, time.Time, time.Duration)
	}{
		{
			name:           "成功生成令牌",
			userID:         1001,
			mockSecret:     conf.GetConf().JWT.Secret,
			mockExpiration: time.Duration(conf.GetConf().JWT.ExpireTime) * time.Hour,
			mockTime:       time.Now(),
			expectedError:  false,
			validateResp: func(t *testing.T, resp *auth.DeliveryResp, mockTime time.Time, mockExpiration time.Duration) {
				// 解析并验证令牌
				token, err := jwt.Parse(resp.Token, func(token *jwt.Token) (interface{}, error) {
					return []byte(conf.GetConf().JWT.Secret), nil
				})
				fmt.Println(token)
				assert.NoError(t, err)
				assert.True(t, token.Valid)

				claims := token.Claims.(jwt.MapClaims)
				assert.Equal(t, float64(1001), claims["user_id"])
				exp := int64(claims["exp"].(float64))

				expectedExp := mockTime.Add(mockExpiration).Unix()
				// 验证过期时间是否一致（允许 ±1 分钟的误差）
				delta := int64(60) // 允许的时间误差范围（秒）
				assert.InDelta(t, expectedExp, exp, float64(delta), "过期时间不匹配，允许误差范围为 ±1 分钟")
			},
		},
		{
			name: "userid为空导致签名失败",
			// userID:         ,
			mockSecret:     conf.GetConf().JWT.Secret,
			mockExpiration: time.Duration(conf.GetConf().JWT.ExpireTime),
			mockTime:       time.Now(),
			expectedError:  true,
			expectedErrMsg: "Userid is empty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			now = func() time.Time { return tc.mockTime }

			// 创建服务实例
			service := &DeliverTokenByRPCService{}

			// 构造请求
			req := &auth.DeliverTokenReq{
				UserId: int32(tc.userID),
			}
			fmt.Println("输出密钥")
			fmt.Println(conf.GetConf().JWT.Secret)

			// 调用被测试方法
			resp, err := service.Run(req)

			// 验证结果
			if tc.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErrMsg)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				tc.validateResp(t, resp, tc.mockTime, tc.mockExpiration)
			}
		})
	}
}
