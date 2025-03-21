package handlers

import (
	"context"
	"net/http"
)


/*
 * 统计_用户性别
 * GET /admin/echarts/sex
 */
func (h *Handler) CountGenders(w http.ResponseWriter, r *http.Request) {
	result, err := h.db.CountUserGenders(context.TODO())
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}
	if err := h.jsonify(w, true, "统计用户性别成功", result); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}


/*
 * 统计_每个月的注册用户数量
 * GET /admin/echarts/user
 */
func (h *Handler) CountUser(w http.ResponseWriter, r *http.Request) {
	result, err := h.db.CountMonthlyUserRegistrations(context.TODO())
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	if err := h.jsonify(w, true, "统计每个月的注册用户数量成功", result); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}





