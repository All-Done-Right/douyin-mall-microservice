package service_test

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/mocks"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"google.golang.org/grpc/status"
)

func TestAddItemService_Run_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 Mock 对象
	mockProductClient := mocks.NewMockClient(ctrl)
	mockCartStore := mocks.NewMockCartStore(ctrl)

	// 设置 Mock 预期
	productID := uint32(100)
	quantity := 2
	userID := uint32(2001)

	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), &product.GetProductReq{Id: productID}).
		Return(&product.GetProductResp{Product: &product.Product{Id: productID}}, nil)

	mockCartStore.EXPECT().
		AddItem(gomock.Any(), &model.Cart{
			UserID:    userID,
			ProductID: productID,
			Qty:       quantity,
		}).
		Return(nil)

	// 初始化服务
	svc := NewAddItemService(
		context.Background(),
		mockCartStore,
		mockProductClient,
	)

	// 准备请求
	req := &cart.AddItemReq{
		UserId: userID,
		Item: &cart.CartItem{
			ProductId: productID,
			Quantity:  quantity,
		},
	}

	// 执行测试
	resp, err := svc.Run(req)

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestAddItemService_Run_ProductClientNil(t *testing.T) {
	// 初始化服务，不传入 ProductClient
	svc := NewAddItemService(
		context.Background(),
		nil,
		nil,
	)

	// 准备请求
	req := &cart.AddItemReq{
		UserId: 2001,
		Item: &cart.CartItem{
			ProductId: 100,
			Quantity:  2,
		},
	}

	// 执行测试
	resp, err := svc.Run(req)

	// 验证结果
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, 50000, kerrors.GetErrorCode(err))
}

func TestAddItemService_Run_ProductNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 Mock 对象
	mockProductClient := mocks.NewMockClient(ctrl)
	mockCartStore := mocks.NewMockCartStore(ctrl)

	// 设置 Mock 预期
	productID := uint32(100)
	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), &product.GetProductReq{Id: productID}).
		Return(&product.GetProductResp{Product: nil}, nil)

	// 初始化服务
	svc := NewAddItemService(
		context.Background(),
		mockCartStore,
		mockProductClient,
	)

	// 准备请求
	req := &cart.AddItemReq{
		UserId: 2001,
		Item: &cart.CartItem{
			ProductId: productID,
			Quantity:  2,
		},
	}

	// 执行测试
	resp, err := svc.Run(req)

	// 验证结果
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, 40004, kerrors.GetErrorCode(err))
}

func TestAddItemService_Run_RPCError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 Mock 对象
	mockProductClient := mocks.NewMockClient(ctrl)
	mockCartStore := mocks.NewMockCartStore(ctrl)

	// 设置 Mock 预期
	productID := uint32(100)
	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), &product.GetProductReq{Id: productID}).
		Return(nil, status.Error(500, "RPC error"))

	// 初始化服务
	svc := NewAddItemService(
		context.Background(),
		mockCartStore,
		mockProductClient,
	)

	// 准备请求
	req := &cart.AddItemReq{
		UserId: 2001,
		Item: &cart.CartItem{
			ProductId: productID,
			Quantity:  2,
		},
	}

	// 执行测试
	resp, err := svc.Run(req)

	// 验证结果
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestAddItemService_Run_CartStoreError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 Mock 对象
	mockProductClient := mocks.NewMockClient(ctrl)
	mockCartStore := mocks.NewMockCartStore(ctrl)

	// 设置 Mock 预期
	productID := uint32(100)
	quantity := 2
	userID := uint32(2001)

	mockProductClient.EXPECT().
		GetProduct(gomock.Any(), &product.GetProductReq{Id: productID}).
		Return(&product.GetProductResp{Product: &product.Product{Id: productID}}, nil)

	mockCartStore.EXPECT().
		AddItem(gomock.Any(), &model.Cart{
			UserID:    userID,
			ProductID: productID,
			Qty:       quantity,
		}).
		Return(kerrors.NewBizStatusError(50000, "database error"))

	// 初始化服务
	svc := NewAddItemService(
		context.Background(),
		mockCartStore,
		mockProductClient,
	)

	// 准备请求
	req := &cart.AddItemReq{
		UserId: userID,
		Item: &cart.CartItem{
			ProductId: productID,
			Quantity:  quantity,
		},
	}

	// 执行测试
	resp, err := svc.Run(req)

	// 验证结果
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, 50000, kerrors.GetErrorCode(err))
}
