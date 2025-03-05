// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"

	auth "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/hertz_gen/frontend/auth"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/infra/rpc"
	frontendutils "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/utils"
	rpcuser "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp string, userId int32, err error) {
	// 调用 RPC 登录接口
	res, err := rpc.UserClient.Login(h.Context, &rpcuser.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil {
		return "", 0, err
	}

	// 登录成功，保存会话
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", res.UserId)
	err = session.Save()
	if err != nil {
		return "", 0, err // 返回错误而不是 panic
	}

	// 设置重定向地址
	redirect := "/"
	if frontendutils.ValidateNext(req.Next) {
		redirect = req.Next
	}

	return redirect, res.UserId, nil
}

// 自定义错误类型，用于区分“用户不存在”
type UserNotFoundError struct {
	Email string
}

func (e *UserNotFoundError) Error() string {
	return "用户不存在"
}
