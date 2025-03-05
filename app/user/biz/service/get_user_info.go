package service

import (
	"context"
	"time"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	userModel, err := model.GetByID(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}

	// 构建响应
	return &user.GetUserInfoResp{
		UserId:     int32(userModel.ID),
		Email:      userModel.Email,
		Nickname:   userModel.Nickname,
		Avatar:     userModel.Avatar,
		Phone:      userModel.Phone,
		Address:    userModel.Address,
		CreateTime: userModel.CreatedAt.Format(time.RFC3339),
	}, nil
}
