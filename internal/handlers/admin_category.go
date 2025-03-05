package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	db "clwy-api/internal/database"
)


 
type addCategoryRequest struct {
    Name  string `json:"name" validate:"required,min=2,max=45"`  // 必须填写, 不能为空, 长度在 2~45 之间
    Rank  int `json:"rank" validate:"required,gt=0"`  // 必须是正整数
}

/**
 * 新增分类
 * POST /admin/categories/:id
 */
func (h *Handler) AddCategory(w http.ResponseWriter, r *http.Request) {
	
	var req addCategoryRequest
	
	if err := readJSON(w, r, &req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(&req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	arg := db.AddCategoryParams{
		Name: req.Name,
		Rank: req.Rank,
	}	
	err := h.db.AddCategory(ctx, &arg)
	if err != nil {
		switch err {
		case db.ErrAlreadyExists:
			h.badRequestResponse(w, r, errors.New("名称已存在, 请选择其它名称"))
			return
		default:
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonify(w, true, "查询分类列表成功", nil); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}


/**
 * 查询分类列表
 * GET /admin/categories
 */
 func (h *Handler) AllCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	items, err := h.db.GetAllCategories(ctx)
	if err!= nil {
		h.internalServerError(w, r, err)
		return
	}
	
	if err := h.jsonify(w, true, "查询分类列表成功", items); err!= nil {
		h.internalServerError(w, r, err)
		return
	}
 }
 

/**
 * 删除分类
 * DELETE /admin/categories/:id
 */
 func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	
	id, err := getInt64FromPathParam(r)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := context.TODO()
	
	err = h.db.DeleteCategory(ctx, id)
	if err != nil {
		switch err {
		case db.ErrNotFound:
			h.notFoundResponse(w, r, fmt.Errorf("ID %d 的分类未找到", id))
			return
		default:
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonify(w, true, "删除分类成功", nil); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}

