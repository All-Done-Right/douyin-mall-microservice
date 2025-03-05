package service

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)

	// 准备测试用户
	password := "123456"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	model.Create(mysql.DB, ctx, &model.User{Email: "test1@test.com", PasswordHashed: string(hashedPassword)})

	// 测试用例：正常登录
	req := &user.LoginReq{
		Email:    "test1@test.com",
		Password: password,
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	t.Logf("resp: %v", resp)

	// 测试用例：邮箱不存在
	req.Email = "nonexistent@test.com"
	resp, err = s.Run(req)
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	// 测试用例：密码错误
	req.Email = "test1@test.com"
	req.Password = "wrongpassword"
	resp, err = s.Run(req)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}
