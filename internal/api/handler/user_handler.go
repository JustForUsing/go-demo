package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"item-manager-new/internal/api/dto"
	"item-manager-new/internal/api/request"
	"item-manager-new/internal/api/response"
	"item-manager-new/internal/auth"
	"item-manager-new/internal/errors/business"
	"item-manager-new/internal/repos"
	"item-manager-new/internal/services"
	"strings"
)

type UserHandler struct {
	sessionManager *auth.Manager
}

func NewUserHandler(sessionManager *auth.Manager) *UserHandler {
	return &UserHandler{
		sessionManager: sessionManager,
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var req request.Login
	resp := response.NewBuilder(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest("请求体格式错误")
		return
	}

	username := strings.TrimSpace(req.Username)
	email := strings.TrimSpace(req.Email)
	if username == "" && email == "" {
		resp.BadRequest("用户名或邮箱必需提供其一")
		return
	}

	if req.Password == "" {
		resp.BadRequest("用户名/邮箱或密码不能为空")
		return
	}

	user, err := services.Instance.User.Auth(username, email, req.Password)
	if err != nil {
		if errors.Is(err, business.ErrInvalidCredential) {
			resp.Unauthorized("用户名/邮箱或密码错误")
			return
		}
		resp.Unauthorized(fmt.Sprintf("认证失败: %v", err))
		return
	}

	err = u.sessionManager.SetSession(c, user.ID)
	if err != nil {
		resp.InternalServerError("创建会话失败")
		return
	}

	roleList, _ := repos.Instance.Role.ListUserRoles(user.ID)
	_ = services.Instance.Audit.Record(user, "用户 %s 登录", user.Username)
	resp.OK(dto.ToUserDTO(user, dto.MapRoles(roleList)))
}
