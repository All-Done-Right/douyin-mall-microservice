package service

import (
	"context"
	"strconv"

	auth "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/hertz_gen/frontend/auth"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/infra/rpc"
	user "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type UpdateUserProfileService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserProfileService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserProfileService {
	return &UpdateUserProfileService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserProfileService) Run(req *auth.UpdateUserProfileRequest) (resp *auth.UpdateUserProfileResponse, err error) {
	resp = new(auth.UpdateUserProfileResponse)

	// 获取用户ID，如果未提供则从Session中获取
	userID := req.UserId
	if userID == 0 {
		session := sessions.Default(h.RequestContext)
		userIDInterface := session.Get("user_id")
		if userIDInterface == nil {
			resp.Success = false
			resp.Message = "用户未登录"
			return resp, nil
		}

		// 处理不同类型的转换
		switch v := userIDInterface.(type) {
		case int32:
			userID = v
		case int:
			userID = int32(v)
		case int64:
			userID = int32(v)
		case float64:
			userID = int32(v)
		default:
			// 如果是其他类型，尝试转换为字符串再解析
			if idStr, ok := userIDInterface.(string); ok {
				if idInt, err := strconv.Atoi(idStr); err == nil {
					userID = int32(idInt)
				}
			}
		}
	}

	// 调用用户服务更新用户信息
	r, err := rpc.UserClient.UpdateUserInfo(h.Context, &user.UpdateUserInfoReq{
		UserId:   userID,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Phone:    req.Phone,
		Address:  req.Address,
	})
	if err != nil {
		resp.Success = false
		resp.Message = "更新用户信息失败: " + err.Error()
		return resp, nil
	}

	resp.Success = r.Success
	resp.Message = r.Message

	return resp, nil
}
