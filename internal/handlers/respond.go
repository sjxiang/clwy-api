package handlers

import (
	"net/http"
)

// 500 
// 服务器开小差啦, 稍后再来试一试 
// 数据库繁忙, 请稍后再试
func (h *Handler) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}


// 400
// 无效的请求
func (h *Handler) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

// 409
// 编辑冲突
func (h *Handler) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Errorf("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusConflict, err.Error())
}

// 404
// 资源未找到
func (h *Handler) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusNotFound, "not found")
}

// 429
// 速率限制
func (h *Handler) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	h.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}


// 403
// 携带了凭证, 但权限不足（授权）
func (h *Handler) notPermittedResponse(w http.ResponseWriter, r *http.Request) {
	h.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")

	writeJSONError(w, http.StatusUnauthorized, "you do not have permission to perform this action")
}


// 401
// 认证错误
// 未携带凭证 authentication required 
// 凭证格式错误
// 凭证过期 invalid or expired authentication token
// 密码错误 invalid credentials
func (h *Handler) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Warnf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}


func (h *Handler) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logger.Warnf("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

