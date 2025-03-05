package service_test

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/mocks"
	"gorm.io/gorm"

	service "github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/service"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestEmptyCartService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCartStore(ctrl)
	mockStore.EXPECT().
		EmptyCart(gomock.Any(), uint32(1001)).
		Return(nil)

	svc := service.NewEmptyCartService(context.Background(), mockStore)
	resp, err := svc.Run(&cart.EmptyCartReq{UserId: 1001})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestEmptyCartService_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCartStore(ctrl)
	mockStore.EXPECT().
		EmptyCart(gomock.Any(), uint32(1001)).
		Return(gorm.ErrInvalidTransaction)

	svc := service.NewEmptyCartService(context.Background(), mockStore)
	resp, err := svc.Run(&cart.EmptyCartReq{UserId: 1001})

	assert.Error(t, err)
	assert.Nil(t, resp)
}
