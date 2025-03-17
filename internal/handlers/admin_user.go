package handlers

import (
	"context"
	"net/http"
	"errors"
	"fmt"

	db "clwy-api/internal/database"
)

/**
 * 添加用户
 * POST /admin/users
 */
func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	
	var req struct {
		Email    string `json:"email" validate:"required,email"`  // 邮箱必须填写、邮箱格式不正确
		Username string `json:"username" validate:"required,min=2,max=45"`
		Nickname string `json:"nickname" validate:"required,min=2,max=45"`  
		Password string `json:"password" validate:"required,min=6,max=45"`
		Sex      int    `json:"sex" validate:"required,oneof=0 1 2"` // 性别的值必须是，男性：0 女性：1 未选择：2
		Role     int    `json:"role" validate:"required,oneof=0 100"`   // 用户组的值必须是，普通用户：0 管理员：100  
		
		Avatar   string `json:"avatar"`
		Company  string `json:"company"`
		Intro    string `json:"intro"`
	}

	if err := readJSON(w, r, &req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(&req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	if len(req.Avatar) == 0 {
		req.Avatar = "default.png"
	}

	hashedPassword, err := GeneratePasswordHash(req.Password)
	if err!= nil {
		h.internalServerError(w, r, err)
		return
	}

	ctx := context.TODO()

	arg := db.AddUserParams{
		Email:    req.Email,
		Username: req.Username,
		Nickname: req.Nickname,
		Password: string(hashedPassword),
		Sex:      int8(req.Sex),
		Role:     int8(req.Role),
		Avatar:   req.Avatar,

		Company:  req.Company,
		Intro:    req.Intro,
	}
	if err := h.db.AddUser(ctx, arg); err!= nil {
		switch err {
		case db.ErrDuplicateEmail:
			h.badRequestResponse(w, r, errors.New("邮箱已存在, 请直接登录。"))
			return
		case db.ErrDuplicateUsername:
			h.badRequestResponse(w, r, errors.New("用户名已存在。"))
			return
		default:
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonify(w, true, "添加用户成功", nil); err!= nil {
		h.internalServerError(w, r, err)
		return
	}
}



/**
 * 查询当前用户
 * GET /admin/users/:id
 */
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	
	id, err := getInt64FromPathParam(r)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := context.TODO()

	user, err := h.db.GetUser(ctx, id)
	if err!= nil {
		switch err {
		case db.ErrNoRecord:
			h.notFoundResponse(w, r, fmt.Errorf("ID %d 的用户未找到", id))
			return
		default:
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonify(w, true, "查询用户成功", user); err!= nil {
		h.internalServerError(w, r, err)
		return
	}
}


/**
 * 查询用户列表
 * GET /admin/users
 */
 func (h *Handler) AllUsers(w http.ResponseWriter, r *http.Request) {

 }