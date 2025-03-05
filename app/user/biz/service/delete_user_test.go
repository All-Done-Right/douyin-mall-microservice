package service

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/model"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestDeleteUser_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewDeleteUserService(ctx)

	// 准备测试用户
	testUser := &model.User{Email: "delete@test.com", PasswordHashed: "hashed_password"}
	model.Create(mysql.DB, ctx, testUser)

	// 测试用例：正常删除用户
	req := &user.DeleteUserReq{UserId: int32(testUser.ID)}
	resp, err := s.Run(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	t.Logf("resp: %v", resp)
}
