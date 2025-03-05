package service

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)

	// 测试用例：正常注册
	req := &user.RegisterReq{
		Email:           "test@test.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	t.Logf("resp: %v", resp)

	// 测试用例：邮箱已被注册
	model.Create(mysql.DB, ctx, &model.User{Email: req.Email, PasswordHashed: "hashed_password"})
	req.Email = "test@test.com"
	resp, err = s.Run(req)
	if err == nil || err.Error() != "邮箱已被注册" {
		t.Fatalf("expected error '邮箱已被注册', got %v", err)
	}

}
