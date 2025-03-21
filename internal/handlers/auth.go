package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	db "clwy-api/internal/database"
)

/**
 * 管理员登录
 * POST /admin/auth/sign_in
 */

type SignInRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}


func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	
	var req SignInRequest
	// 1. 序列化用户发送的邮箱、账号、密码
	if err := readJSON(w, r, &req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	// 2. 基础验证
	if err := Validate.Struct(&req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	// 3. 接收 login, 而不接收 email 和 username

	// 4. 通过 login 查询数据库, 判断用户存在
	user, err := h.db.GetUserByLogin(ctx, req.Login)
	if err!= nil {
		if errors.Is(err, db.ErrNoRecord) {
			h.notFoundResponse(w, r, fmt.Errorf("用户不存在"))
			return
		}

		h.internalServerError(w, r, err)
		return
	}

	// 5. 验证密码是否正确
	equal, err := ComparePasswordAndHash(req.Password, []byte(user.Password))
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}
	if !equal {
		h.notFoundResponse(w, r, fmt.Errorf("密码错误"))
		return
	}

	// 6. 验证不是管理员
	if user.Role != 100 {
		// 您没有权限登录管理员后台
		h.notPermittedResponse(w, r)
		return
	}

	// 7. 生成 jwt token
	token, err := h.authn.GenerateToken(user.ID, time.Hour)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	// 8. 响应
	if err := h.jsonify(w, true, "登录成功", "Bearer "+token); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}