package service

import (
	"context"

	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
)

type LogoutService struct {
	ctx context.Context
}

func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

func (s *LogoutService) Run(req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	return &user.LogoutResp{Success: true}, nil
}
