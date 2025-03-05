package service

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/global"
	"strings"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/model"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPlaceOrderService_Run(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	testCases := []struct {
		name           string
		req            *order.PlaceOrderReq
		repoSetup      func(*MockOrderRepo) // 设置mock仓库行为
		expectedResp   *order.PlaceOrderResp
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name: "成功下单",
			req: &order.PlaceOrderReq{
				UserId:       1001,
				UserCurrency: "USD",
				Email:        "user1@example.com",
				Address: &order.Address{
					StreetAddress: "123 Main St",
					City:          "New York",
					State:         "NY",
					Country:       "USA",
					ZipCode:       10001,
				},
				OrderItems: []*order.OrderItem{
					{
						Cost: 49.99,
						Item: &cart.CartItem{
							ProductId: 1,
							Quantity:  2,
						},
					},
					{
						Cost: 29.99,
						Item: &cart.CartItem{
							ProductId: 2,
							Quantity:  1,
						},
					},
				},
			},
			repoSetup: func(mockRepo *MockOrderRepo) {
				mockRepo.On("CreateOrder", mock.MatchedBy(func(order model.Order) bool {
					// 验证传递给CreateOrder的参数是否正确
					return order.UserID == 1001 &&
						order.UserCurrency == "USD" &&
						order.Email == "user1@example.com" &&
						len(order.OrderItems) == 2
				})).Return("fixed-uuid-12345", nil)
			},
			expectedResp: &order.PlaceOrderResp{
				Order: &order.OrderResult{
					OrderId: "fixed-uuid-12345",
				},
			},
			expectedError: false,
		},
		{
			name: "地址为空",
			req: &order.PlaceOrderReq{
				UserId:       1001,
				UserCurrency: "USD",
				Email:        "user1@example.com",
				Address:      nil, // 地址为空
				OrderItems: []*order.OrderItem{
					{
						Cost: 49.99,
						Item: &cart.CartItem{
							ProductId: 1,
							Quantity:  2,
						},
					},
				},
			},
			repoSetup: func(mockRepo *MockOrderRepo) {
				// 不会调用CreateOrder
			},
			expectedResp:   nil,
			expectedError:  true,
			expectedErrMsg: "address is empty",
		},
		{
			name: "订单项为空",
			req: &order.PlaceOrderReq{
				UserId:       1001,
				UserCurrency: "USD",
				Email:        "user1@example.com",
				Address: &order.Address{
					StreetAddress: "123 Main St",
					City:          "New York",
					State:         "NY",
					Country:       "USA",
					ZipCode:       10001,
				},
				OrderItems: []*order.OrderItem{}, // 空订单项
			},

			repoSetup: func(mockRepo *MockOrderRepo) {
				// 不会调用CreateOrder
			},
			expectedResp:   nil,
			expectedError:  true,
			expectedErrMsg: "order items is empty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建 miniredis 实例
			mr, err := miniredis.Run()
			if err != nil {
				t.Fatalf("Failed to create miniredis: %v", err)
			}
			defer mr.Close()

			// 创建Redis客户端
			var rdb *redis.Client
			rdb = redis.NewClient(&redis.Options{
				Addr: mr.Addr(),
			})

			// 保存和替换全局Redis客户端
			global.RDB = rdb

			// 创建mock订单仓库
			mockRepo := new(MockOrderRepo)
			tc.repoSetup(mockRepo)

			// 创建服务实例
			service := PlaceOrderService{
				ctx: context.Background(),
				db:  mockRepo,
			}

			// 调用被测试方法
			resp, err := service.Run(tc.req)

			// 验证结果
			if tc.expectedError {
				assert.Error(t, err)
				if tc.expectedErrMsg != "" && !strings.Contains(err.Error(), "dial tcp") { // 忽略Redis连接错误的具体消息
					assert.Contains(t, err.Error(), tc.expectedErrMsg)
				}
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResp.Order.OrderId, resp.Order.OrderId)
			}

			// 验证mock对象的方法是否按预期调用
			mockRepo.AssertExpectations(t)
		})
	}
}
