package handlers

import (
	"context"
	"fmt"
	"net/http"

	db "clwy-api/internal/database"
)

/*
 * 查询章节列表
 * GET /admin/chapters?course_id=1&&current_page=1&&page_size=10
 */
func (h *Handler) AllChapters(w http.ResponseWriter, r *http.Request) {
	// 参数校验
	// 当前是第几页, 如果不传则默认为第1页
	currentPage, err := parseQueryParamToInt64(r, "current_page", 1)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	// 每页显示多少条数据, 如果不传则默认为10条
	pageSize, err := parseQueryParamToInt64(r, "page_size", 10)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	courseId, err := parseQueryParamToInt64(r, "course_id", 0)
	if err != nil {
		h.badRequestResponse(w, r, fmt.Errorf("获取章节列表失败, 课程ID不能为空"))
		return
	}

	ctx := context.TODO()
	arg := db.GetAllChaptersParams{
		CourseId: courseId,
		Offset: (currentPage - 1) * pageSize,
		Limit: pageSize,
	}

	chapters, err := h.db.GetAllChapters(ctx, arg)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	if err := h.jsonify(w, true, "查询章节列表成功", chapters); err!= nil {
		h.internalServerError(w, r, err)
		return
	}
}	
