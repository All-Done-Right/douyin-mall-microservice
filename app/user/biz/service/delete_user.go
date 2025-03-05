package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
)

type DeleteUserService struct {
	ctx context.Context
}

func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	if err = mysql.DB.Delete(&model.User{}, req.UserId).Error; err != nil {
		return nil, err
	}
	return &user.DeleteUserResp{Success: true}, nil
}
