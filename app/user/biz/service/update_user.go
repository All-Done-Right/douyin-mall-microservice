package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	userModel, err := model.GetByID(mysql.DB, s.ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		userModel.Email = req.Email
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		userModel.PasswordHashed = string(hashedPassword)
	}

	if err = model.Update(mysql.DB, s.ctx, userModel); err != nil {
		return nil, err
	}

	return &user.UpdateUserResp{Success: true}, nil
}
