package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (*user.LoginResp, error) {
	// 检查用户是否存在
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("邮箱或密码不能为空")
	}
	existingUser, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("查询用户时出错: %v", err.Error())
	}
	if existingUser == nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码不正确")
	}

	// 登录成功，返回响应
	return &user.LoginResp{UserId: int32(existingUser.ID)}, nil
}
