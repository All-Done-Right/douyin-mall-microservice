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

package auth

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/biz/utils"
	auth "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/hertz_gen/frontend/auth"
	common "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/hertz_gen/frontend/common"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/infra/rpc"
	authcenter "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

// Register .
// @router /auth/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewRegisterService(ctx, c).Run(&req)
	if err != nil {
		// 提取 desc 部分
		errStr := err.Error()
		desc := errStr
		if strings.Contains(errStr, "desc = ") {
			parts := strings.SplitAfter(errStr, "desc = ")
			if len(parts) > 1 {
				desc = strings.TrimSuffix(parts[1], " [biz error]")
			}
		}
		// 使用提取的 desc 返回错误
		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New(desc))
		return
	}
	c.Redirect(consts.StatusFound, []byte("/"))
}

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	redirect, userid, err := service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		// 默认错误信息
		errorMsg := "登录失败，请稍后重试"
		errorType := ""
		// 提取业务错误
		errStr := err.Error()
		if strings.Contains(errStr, "密码不正确") {
			errorMsg = "密码不正确，请检查后重试"
			errorType = "password"
		} else if strings.Contains(errStr, "用户不存在") {
			errorMsg = "邮箱未注册，请先注册"
			errorType = "email"
		}
		// 渲染 sign-in 页面，传递错误信息和类型
		c.HTML(consts.StatusOK, "sign-in", hertzUtils.H{
			"error":     errorMsg,
			"errorType": errorType,
			"email":     req.Email, // 保留用户输入的邮箱
		})
		return
	}

	// 获取 token
	tokenReq := &authcenter.DeliverTokenReq{UserId: userid}
	tokenResp, err := rpc.AuthClient.DeliverTokenByRPC(ctx, tokenReq)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.SetCookie("jwt_token", tokenResp.Token, 86400, "", "", 0, true, false)

	// 登录成功，重定向
	c.Redirect(consts.StatusFound, []byte(redirect))
}

// Logout .
// @router /auth/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewLogoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	redirect := "/"

	c.Redirect(consts.StatusFound, []byte(redirect))
}

// GetUserProfile .
// @router /user/profile [GET]
func GetUserProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.UserProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 从session中获取用户ID
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.HTML(consts.StatusOK, "user-profile.tmpl", hertzUtils.H{
			"title":    "个人信息",
			"Message":  "用户未登录",
			"UserInfo": nil,
		})
		return
	}

	// 设置用户ID到请求中，处理不同类型的转换
	var userIDInt32 int32
	switch v := userID.(type) {
	case int32:
		userIDInt32 = v
	case int:
		userIDInt32 = int32(v)
	case int64:
		userIDInt32 = int32(v)
	case float64:
		userIDInt32 = int32(v)
	default:
		// 如果是其他类型，尝试转换为字符串再解析
		if idStr, ok := userID.(string); ok {
			if idInt, err := strconv.Atoi(idStr); err == nil {
				userIDInt32 = int32(idInt)
			}
		}
	}

	req.UserId = userIDInt32

	resp, err := service.NewGetUserProfileService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// 默认渲染HTML页面，除非明确要求JSON格式
	if c.Query("format") == "json" {
		utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
		return
	}

	// 渲染HTML页面，确保传递user_id给模板
	c.HTML(consts.StatusOK, "user-profile.tmpl", hertzUtils.H{
		"title":      "个人信息",
		"UserInfo":   resp,
		"isLoggedIn": true,
		"user_id":    userID, // 使用原始的userID，而不是转换后的userIDInt32
	})
}

// UpdateUserProfile .
// @router /user/profile/update [POST]
func UpdateUserProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.UpdateUserProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUpdateUserProfileService(ctx, c).Run(&req)

	if err != nil {
		// 获取用户ID以确保在错误页面上仍然显示用户菜单
		session := sessions.Default(c)
		userID := session.Get("user_id")

		c.HTML(consts.StatusOK, "user-profile.tmpl", hertzUtils.H{
			"title":   "个人信息",
			"Message": fmt.Sprintf("更新失败: %v", err),
			"user_id": userID,
		})
		return
	}

	// 如果更新成功且请求来自表单，则重定向到个人信息页面
	if resp.Success && c.Request.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		c.Redirect(consts.StatusFound, []byte("/user/profile"))
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteUser .
// @router /user/delete [GET]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var req auth.DeleteUserRequest

	// 从session中获取用户ID
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.HTML(consts.StatusOK, "user-profile.tmpl", hertzUtils.H{
			"title":    "个人信息",
			"Message":  "用户未登录",
			"UserInfo": nil,
		})
		return
	}

	// 处理不同类型的转换
	var userIDInt32 int32
	switch v := userID.(type) {
	case int32:
		userIDInt32 = v
	case int:
		userIDInt32 = int32(v)
	case int64:
		userIDInt32 = int32(v)
	case float64:
		userIDInt32 = int32(v)
	default:
		// 如果是其他类型，尝试转换为字符串再解析
		if idStr, ok := userID.(string); ok {
			if idInt, err := strconv.Atoi(idStr); err == nil {
				userIDInt32 = int32(idInt)
			}
		}
	}

	req.UserId = userIDInt32

	resp, err := service.NewDeleteUserService(ctx, c).Run(&req)

	if err != nil {
		c.HTML(consts.StatusOK, "user-profile.tmpl", hertzUtils.H{
			"title":   "个人信息",
			"Message": fmt.Sprintf("删除失败: %v", err),
			"user_id": userID, // 确保传递user_id
		})
		return
	}

	// 如果删除成功，清除Session并重定向到首页
	if resp.Success {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.SetCookie("jwt_token", "", -1, "/", "", protocol.CookieSameSiteDisabled, false, true)
		c.Redirect(consts.StatusFound, []byte("/"))
		return
	}

	// 如果删除失败，返回到个人信息页面并显示错误
	c.HTML(consts.StatusOK, "user-profile.tmpl", hertzUtils.H{
		"title":   "个人信息",
		"Message": resp.Message,
		"user_id": userID, // 确保传递user_id
	})
}
