package handlers

import (
	"context"
	"net/http"
	
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
			h.logger.Infow("公告未找到", "status", true)
			h.notFoundResponse(w, r, err)
			return
		default:
			h.logger.Errorw("查询公告失败", "status", false, "err", err)
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
	currentPage, err := parseQueryParamToInt64(r, "current_page", 1)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

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
		h.logger.Errorw("查询公告列表失败", "status", false, "err", err)
		h.internalServerError(w, r, err)
		return
	}
	
	if err := h.jsonResponse(w, http.StatusOK, result); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}


type createNoticeRequest struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=10000"`
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

	id, err := h.db.CreateNotice(ctx, &arg)
	if err!= nil {
		h.logger.Errorw("创建公告失败", "status", false, "err", err)
		h.internalServerError(w, r, err)
		return
	}

	if err := h.jsonResponse(w, http.StatusCreated, id); err != nil {
		h.internalServerError(w, r, err)
	}
}

/**
 * 删除公告
 * DELETE /admin/notices/:id
 */
func (h *Handler) DeleteNotice(w http.ResponseWriter, r *http.Request) {
	id, err := getInt64FromPathParam(r)
	if err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := context.TODO()
	
	err = h.db.DeleteNotice(ctx, id)
	if err != nil {
		switch err {
		case db.ErrNotFound:
			h.logger.Infow("公告未找到", "status", true)
			h.notFoundResponse(w, r, err)
			return
		default:
			h.logger.Errorw("删除公告失败", "status", false, "err", err)
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonResponse(w, http.StatusOK, "删除公告成功"); err != nil {
		h.internalServerError(w, r, err)
	}
}


type updateNoticeRequest struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
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
			h.logger.Infow("公告未找到", "status", true)
			h.notFoundResponse(w, r, err)
			return
		default:
			h.logger.Errorw("更新公告失败", "status", false, "err", err)
			h.internalServerError(w, r, err)
			return
		}
	}

	if err := h.jsonResponse(w, http.StatusOK, "更新公告成功"); err != nil {
		h.internalServerError(w, r, err)
	}
}

