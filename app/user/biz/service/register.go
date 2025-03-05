package service

import (
	"context"
	"errors"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	//fmt.Printf("email: %s, password: %s, confirm_password: %s\n", req.Email, req.Password, req.ConfirmPassword)
	existingUser, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	//fmt.Printf("existingUser: %v  error: %v \n", existingUser, err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 邮箱不存在，可以继续注册
			err = nil
		} else {
			// 其他错误，返回错误信息
			return nil, err
		}
	} else if existingUser != nil {
		// 邮箱已存在
		return nil, errors.New("邮箱已被注册")
	}
	// 继续注册流程...
	if req.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	if req.Password == "" || req.ConfirmPassword == "" {
		return nil, errors.New("密码不能为空")
	}
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("两次输入的密码不一致")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		return
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
