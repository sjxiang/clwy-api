package handlers

import (
	"context"
	"net/http"
	"fmt"
	"strconv"

	"github.com/go-chi/chi/v5"
	db "clwy-api/internal/database"
)

/**
 * 查询公告详情
 * GET /admin/notices/:id
 */
func (h *Handler) GetNotice(w http.ResponseWriter, r *http.Request) {
	
	id, err := getInt64FromPathParam(r)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := context.TODO()

	notice, err := h.db.GetNotice(ctx, id)
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

	if err := h.jsonResponse(w, http.StatusOK, notice); err != nil {
		h.internalServerError(w, r, err)	
		return
	}
}

/**
 * 分页模糊搜索查询公告
 * GET /admin/notices?title=你好&&current_page=1&&page_size=10
 */
func (h *Handler) AllNotices(w http.ResponseWriter, r *http.Request) {

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

	title := readQueryParam(r, "title")

	ctx := context.TODO()

	arg := db.GetNoticesWithPaginationParams{
		SearchKey: title,
		Limit:     pageSize,
		Offset:    (currentPage - 1) * pageSize,
	}
	
	result, err := h.db.GetNoticesWithPagination(ctx, &arg)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}
	
	if err := h.jsonResponse(w, http.StatusOK, result); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}



type createNoticeRequest struct {
    Title   string `json:"title" validate:"required,min=2,max=45"`  // 标题长度 2~45, 标题不能为空
    Content string `json:"content" validate:"required,max=10000"`  // 内容长度 <=10000
}

/**
 * 创建公告
 * POST /admin/notices
 */
func (h *Handler) CreateNotice(w http.ResponseWriter, r *http.Request) {
	
	var req createNoticeRequest
	
	if err := readJSON(w, r, &req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	// 参数校验	
	if err := Validate.Struct(&req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := context.TODO()

	arg := db.CreateNoticeParams{
		Title:   req.Title,
		Content: req.Content,
	}

	err := h.db.CreateNotice(ctx, &arg)
	if err!= nil {
		h.internalServerError(w, r, err)
		return
	}

	if err := h.jsonResponse(w, http.StatusCreated, nil); err != nil {
		h.internalServerError(w, r, err)
	}
}

/**
 * 删除公告
 * DELETE /admin/notices/:id
 */
func (h *Handler) DeleteNotice(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	ctx := context.TODO()
	
	err = h.db.DeleteNotice(ctx, id)
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

	if err := h.jsonResponse(w, http.StatusOK, nil); err != nil {
		h.internalServerError(w, r, err)
	}
}


type updateNoticeRequest struct {
	Title   string `json:"title" validate:"required,min=2,max=45"`  
    Content string `json:"content" validate:"required,max=10000"`  
}

/*
 * 更新公告
 * PUT /admin/notices/:id
 */
func (h *Handler) UpdateNotice(w http.ResponseWriter, r *http.Request) {
	
	id, err := getInt64FromPathParam(r)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
  
	var req updateNoticeRequest
	
	if err := readJSON(w, r, &req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(&req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := context.TODO()
	
	arg := db.UpdateNoticeParams{
		ID: id,
		Title:   req.Title,
		Content: req.Content,
	}
	if err = h.db.UpdateNotice(ctx, &arg); err != nil {
		switch err {
		case db.ErrNotFound:
			h.notFoundResponse(w, r, fmt.Errorf("ID %d 的分类未找到", id))
			return
		default:
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonResponse(w, http.StatusOK, nil); err != nil {
		h.internalServerError(w, r, err)
	}
}

