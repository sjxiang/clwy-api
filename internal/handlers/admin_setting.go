package handlers

import (
	"context"
	"errors"
	"net/http"

	db "clwy-api/internal/database"
)

/**
 * 查询当前系统设置
 * GET /admin/settings
 */
func (h *Handler) GetSetting(w http.ResponseWriter, r *http.Request) {
	
	ctx := context.TODO()

	setting, err := h.db.GetSetting(ctx)
	if err != nil {
		switch err {
		case db.ErrNotFound:
			h.notFoundResponse(w, r, errors.New("初始系统设置未找到, 请允许种子文件。"))
			return
		default:
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonResponse(w, http.StatusOK, setting); err!= nil {
		h.internalServerError(w, r, err)
		return
	}
}



type UpdateSettingRequest struct {
	Name          string    `json:"name" validate:"required,min=2,max=45"`
	ICP           string    `json:"icp" validate:"required"`
	Copyright     string    `json:"copyright" validate:"required"`	
}

/**
 * 更新系统设置
 * PUT /admin/settings
 */
func (h *Handler) UpdateSetting(w http.ResponseWriter, r *http.Request) {

	var req UpdateSettingRequest

	if err := readJSON(w, r, &req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(&req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := context.TODO()
	arg := db.UpdateSettingParams{
		Name: req.Name,
		ICP: req.ICP,
		Copyright: req.Copyright,
	}
	
	if err := h.db.UpdateSetting(ctx, arg); err != nil {
		switch err {
		case db.ErrNotFound:
			h.notFoundResponse(w, r, errors.New("初始系统设置未找到, 请允许种子文件。"))
			return
		default:
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonResponse(w, http.StatusOK, nil); err!= nil {
		h.internalServerError(w, r, err)
		return
	}
}