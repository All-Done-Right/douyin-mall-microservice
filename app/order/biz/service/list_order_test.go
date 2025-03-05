package service

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/dal/repo/model"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order"
	"github.com/stretchr/testify/assert"
)

// 创建一个模拟的订单仓库

func TestListOrderService_Run(t *testing.T) {
	// 测试用例定义
	testCases := []struct {
		name           string
		userID         uint32
		mockSetup      func(*MockOrderRepo)
		expectedOrders int
		expectedError  bool
		expectedErrMsg string
		validateResp   func(*testing.T, *order.ListOrderResp)
	}{
		{
			name:   "用户有多个订单",
			userID: 1001,
			mockSetup: func(mockRepo *MockOrderRepo) {
				// 构造测试数据
				orders := []model.Order{
					{
						OrderID:      "order1",
						UserID:       1001,
						UserCurrency: "USD",
						Email:        "user1@example.com",
						Address: model.Address{
							StreetAddress: "123 Main St",
							City:          "New York",
							State:         "NY",
							Country:       "USA",
							ZipCode:       10001,
						},
						OrderItems: []model.OrderItem{
							{
								ProductID: 1,
								Quantity:  2,
								Cost:      49.99,
							},
							{
								ProductID: 2,
								Quantity:  1,
								Cost:      29.99,
							},
						},
					},
					{
						OrderID:      "order2",
						UserID:       1001,
						UserCurrency: "USD",
						Email:        "user1@example.com",
						Address: model.Address{
							StreetAddress: "123 Main St",
							City:          "New York",
							State:         "NY",
							Country:       "USA",
							ZipCode:       10001,
						},
						OrderItems: []model.OrderItem{
							{
								ProductID: 3,
								Quantity:  1,
								Cost:      99.99,
							},
						},
					},
				}
				mockRepo.On("ListOrders", uint32(1001)).Return(orders, nil)
			},
			expectedOrders: 2,
			expectedError:  false,
			validateResp: func(t *testing.T, resp *order.ListOrderResp) {
				assert.Equal(t, 2, len(resp.Orders))

				// 验证第一个订单
				assert.Equal(t, "order1", resp.Orders[0].OrderId)
				assert.Equal(t, uint32(1001), resp.Orders[0].UserId)
				assert.Equal(t, "USD", resp.Orders[0].UserCurrency)
				assert.Equal(t, "user1@example.com", resp.Orders[0].Email)
				assert.Equal(t, "123 Main St", resp.Orders[0].Address.StreetAddress)

				// 验证第一个订单的订单项
				assert.Equal(t, 2, len(resp.Orders[0].OrderItems))
				assert.Equal(t, uint32(1), resp.Orders[0].OrderItems[0].Item.ProductId)
				assert.Equal(t, int32(2), resp.Orders[0].OrderItems[0].Item.Quantity)
				assert.Equal(t, float32(49.99), resp.Orders[0].OrderItems[0].Cost)

				// 验证第二个订单
				assert.Equal(t, "order2", resp.Orders[1].OrderId)
				assert.Equal(t, 1, len(resp.Orders[1].OrderItems))
			},
		},
		{
			name:   "用户没有订单",
			userID: 1002,
			mockSetup: func(mockRepo *MockOrderRepo) {
				// 返回空订单列表
				mockRepo.On("ListOrders", uint32(1002)).Return([]model.Order{}, nil)
			},
			expectedOrders: 0,
			expectedError:  false,
			validateResp: func(t *testing.T, resp *order.ListOrderResp) {
				assert.Equal(t, 0, len(resp.Orders))
			},
		},
	}

	// 执行测试
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 创建 mock 对象
			mockRepo := new(MockOrderRepo)

			// 设置 mock 行为
			tc.mockSetup(mockRepo)

			// 创建服务实例
			service := NewListOrderService(context.Background(), mockRepo)

			// 构造请求
			req := &order.ListOrderReq{
				UserId: tc.userID,
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
				assert.NotNil(t, resp)
				if tc.validateResp != nil {
					tc.validateResp(t, resp)
				}
			}

			// 验证 mock 对象的方法是否被按预期调用
			mockRepo.AssertExpectations(t)
		})
	}
}
