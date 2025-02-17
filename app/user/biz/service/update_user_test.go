package service

import (
	"context"
	"testing"

	"github.com/All-Done-Right/douyin-mall-microservice/app/user/biz/dal/mysql"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
)

func TestUpdateUser_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		klog.Error(err.Error())
	}
	mysql.Init()
	ctx := context.Background()
	s := NewUpdateUserService(ctx)
	// init req and assert value

	req := &user.UpdateUserReq{
		UserId:   5,
		Email:    "update_email_test",
		Password: "update_password_test",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
