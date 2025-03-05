package service

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestGetUserInfo_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewGetUserInfoService(ctx)

	// 准备测试用户
	testUser := &model.User{Email: "info@test.com", PasswordHashed: "hashed_password"}
	model.Create(mysql.DB, ctx, testUser)

	// 测试用例：正常获取用户信息
	req := &user.GetUserInfoReq{UserId: int32(testUser.ID)}
	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	t.Logf("resp: %v", resp)

}
