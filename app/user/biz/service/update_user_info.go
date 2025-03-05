package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
)

type UpdateUserInfoService struct {
	ctx context.Context
} // NewUpdateUserInfoService new UpdateUserInfoService
func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	// 获取用户
	userModel, err := model.GetByID(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return &user.UpdateUserInfoResp{
			Success: false,
			Message: "用户不存在",
		}, nil
	}

	// 更新用户信息
	userModel.Nickname = req.Nickname
	userModel.Avatar = req.Avatar
	userModel.Phone = req.Phone
	userModel.Address = req.Address

	// 保存更新
	if err = model.Update(mysql.DB, s.ctx, userModel); err != nil {
		return &user.UpdateUserInfoResp{
			Success: false,
			Message: "更新失败: " + err.Error(),
		}, nil
	}

	return &user.UpdateUserInfoResp{
		Success: true,
		Message: "更新成功",
	}, nil
}
