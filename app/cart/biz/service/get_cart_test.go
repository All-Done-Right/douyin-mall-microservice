package service_test

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/mocks"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
	//"go.uber.org/mock/gomock"
)

func TestGetCartService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 Mock CartStore
	mockStore := mocks.NewMockCartStore(ctrl)
	/*
		expectedCarts := []model.Cart{
			{ProductID: 1, Qty: 2},
			{ProductID: 2, Qty: 3},
		}
	*/
	// 设置 Mock 预期行为
	mockStore.EXPECT().
		GetCartByUserId(gomock.Any(), gomock.Eq(uint32(1001))).
		Return([]model.Cart{
			{ProductID: 1, Qty: 2},
			{ProductID: 2, Qty: 3},
		}, nil)

	// 初始化服务
	svc := service.NewGetCartService(context.Background(), mockStore)
	resp, err := svc.Run(&cart.GetCartReq{UserId: 1001})

	// 验证结果
	assert.NoError(t, err)
	assert.Len(t, resp.Items, 2)
	assert.Equal(t, int32(1), resp.Items[0].ProductId)
}

func TestGetCartService_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCartStore(ctrl)
	mockStore.EXPECT().
		GetCartByUserId(gomock.Any(), uint32(1001)).
		Return(nil, gorm.ErrRecordNotFound)

	svc := service.NewGetCartService(context.Background(), mockStore)
	resp, err := svc.Run(&cart.GetCartReq{UserId: 1001})

	assert.Error(t, err)
	assert.Nil(t, resp)
}
