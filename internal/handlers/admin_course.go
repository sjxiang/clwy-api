package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	db "clwy-api/internal/database"
)

/*
 * 添加课程
 * POST /admin/courses
 */
func (h *Handler) AddCourse(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name          string    `json:"name" validate:"required"`        
		Image         string    `json:"image" validate:"required"`          
		Content       string    `json:"content" validate:"required"`       
		LikesCount    int64     `json:"likes_count" validate:"required"`   
		ChaptersCount int64     `json:"chapters_count" validate:"required"` 

		Recommended   bool      `json:"recommended"`
		Introductory  bool      `json:"introductory"`  
		
		CategoryID    int64     `json:"category_id" validate:"required"`    
		UserID        int64     `json:"user_id" validate:"required"`        
	}

	if err := readJSON(w, r, &req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(&req); err != nil {
		h.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	arg := db.CreateCourseParams{
		Name: req.Name,
		Image: req.Image,
		Content: req.Content,
		Recommended: req.Recommended,
		Introductory: req.Introductory,
		CategoryId: req.CategoryID,
		UserId: req.UserID,
	}
	err := h.db.CreateCourse(ctx, &arg)
	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			h.conflictResponse(w, r, fmt.Errorf("ID为 %d 的分类不存在, 或者ID为 %d的用户不存在", req.CategoryID, req.UserID))
			return
		}
		if errors.Is(err, db.ErrAlreadyExists) {
			h.conflictResponse(w, r, fmt.Errorf("课程名为 %s 已存在", req.Name))
			return
		}
		
		h.internalServerError(w, r, err)
		return
	}

	if err := h.jsonify(w, true, "添加课程成功", nil); err!= nil {
		h.internalServerError(w, r, err)
		return
	}
}


/*
 * 查询课程列表
 * GET /admin/courses?name=云&&recommender=true&&introductory=true&&current_page=1&&page_size=10
 */

func (h *Handler) AllCourses(w http.ResponseWriter, r *http.Request) {
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

	name := readQueryParam(r, "name")
	recommender := readQueryParam(r, "recommender")
	re, err := strconv.ParseBool(recommender)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	introductory := readQueryParam(r, "introductory")
	in, err := strconv.ParseBool(introductory)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}

	ctx := context.TODO()

	arg := db.FindAndCountAllCoursesParams{
		KeyWord: name,
		Limit: pageSize,
		Offset: (currentPage - 1) * pageSize,
		Recommended: re,
		Introductory: in,
	}

	result, err := h.db.FindAndCountAllCourses(ctx, arg)
	if err != nil {
		h.internalServerError(w, r, err)
		return
	}
	if err := h.jsonify(w, true, "查询课程列表成功", result); err != nil {
		h.internalServerError(w, r, err)
		return
	}
}


/*
 * 查询课程详情
 * GET /admin/courses/:id
 */
func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {

}
