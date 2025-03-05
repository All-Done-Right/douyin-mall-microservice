package service

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestUpdateUserInfo_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewUpdateUserInfoService(ctx)

	// 准备测试用户
	testUser := &model.User{Email: "update@test.com", PasswordHashed: "hashed_password"}
	model.Create(mysql.DB, ctx, testUser)

	// 测试用例：正常更新用户信息
	req := &user.UpdateUserInfoReq{
		UserId:   int32(testUser.ID),
		Nickname: "New Nickname",
		Avatar:   "new_avatar_url",
		Phone:    "1234567890",
		Address:  "New Address",
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	t.Logf("resp: %v", resp)

	// 测试用例：更新不存在的用户信息
	req.UserId = 99999 // 假设这个ID不存在
	resp, err = s.Run(req)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}
