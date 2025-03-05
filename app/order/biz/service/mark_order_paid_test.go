package service

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试 MarkOrderPaidService.Run
func TestMarkOrderPaidService_Run(t *testing.T) {
	// 设置 logrus 为测试模式
	logrus.SetLevel(logrus.DebugLevel)

	testCases := []struct {
		name           string
		userID         uint32
		orderID        string
		redisSetup     func(*miniredis.Miniredis)
		repoSetup      func(*MockOrderRepo)
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name:    "成功标记订单已支付",
			userID:  1001,
			orderID: "order123",
			redisSetup: func(mr *miniredis.Miniredis) {
				mr.Set("order123", "1001")
			},
			repoSetup: func(repo *MockOrderRepo) {
				repo.On("MarkOrderPaid", "order123").Return(nil)
			},
			expectedError: false,
		}, {
			name:    "订单已过期",
			userID:  1001,
			orderID: "expired_order",
			redisSetup: func(mr *miniredis.Miniredis) {
				// redis 中没有该订单即为过期

			},
			repoSetup: func(repo *MockOrderRepo) {

			},
			expectedError:  true,
			expectedErrMsg: "order not found",
		},
		{
			name:    "订单不存在",
			userID:  1001,
			orderID: "nonexistent",
			redisSetup: func(mr *miniredis.Miniredis) {
			},
			repoSetup: func(repo *MockOrderRepo) {
			},
			expectedError:  true,
			expectedErrMsg: "order not found",
		},
		{
			name:    "订单不属于该用户",
			userID:  1001,
			orderID: "order456",
			redisSetup: func(mr *miniredis.Miniredis) {
				mr.Set("order456", "2002") // 不同的用户ID
			},
			repoSetup: func(repo *MockOrderRepo) {
				// 不会调用到仓库方法
			},
			expectedError:  true,
			expectedErrMsg: "order not belong to the user",
		},
	}

	// 执行测试
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建 miniredis 模拟 Redis 服务器
			mr, err := miniredis.Run()
			if err != nil {
				t.Fatalf("Failed to create miniredis: %v", err)
			}
			defer mr.Close()

			// 创建连接到 miniredis 的 Redis 客户端
			rdb := redis.NewClient(&redis.Options{
				Addr: mr.Addr(),
			})
			defer rdb.Close()

			// 设置 Redis 测试数据
			tc.redisSetup(mr)

			// 创建 mock 订单仓库
			mockRepo := new(MockOrderRepo)
			tc.repoSetup(mockRepo)

			// 保存和替换全局 Redis 客户端
			global.RDB = rdb

			// 创建服务实例
			service := MarkOrderPaidService{
				ctx: context.Background(),
				db:  mockRepo,
			}

			// 构造请求
			req := &order.MarkOrderPaidReq{
				UserId:  tc.userID,
				OrderId: tc.orderID,
			}

			// 调用被测试方法
			resp, err := service.Run(req)

			// 验证结果
			if tc.expectedError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErrMsg, err.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
			}

			// 验证 mock 对象的方法是否被按预期调用
			mockRepo.AssertExpectations(t)
		})
	}
}
