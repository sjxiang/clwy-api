package database

import "time"


// 公告
type Notice struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content,omitempty"`
}


// 分类
type Category struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Rank      int64     `json:"rank"`
}


// 课程
type Course struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CategoryID    int64     `json:"category_id"`
	UserID        int64     `json:"user_id"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Recommended   bool      `json:"recommended"`
	Introductory  bool      `json:"introductory"`
	Content       string    `json:"content"`
	LikesCount    int64     `json:"likes_count"`
	ChaptersCount int64     `json:"chapters_count"`	
}


// 章节
type Chapter struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CourseID      int64     `json:"course_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Video         string    `json:"video"`
	Rank          int64     `json:"rank"`
}


// 点赞
type Like struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CourseID      int64     `json:"course_id"`
	UserID        int64     `json:"user_id"`
}


// 用户
type User struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	Nickname      string    `json:"nickname"`
	Password      string    `json:"password"`
	Avatar        string    `json:"avatar"`
	Sex           int8      `json:"sex"`
	Company       string    `json:"company"`
	Introduce     string    `json:"introduce"`
	Role          int8      `json:"role"`
}


// 系统设置
type Setting struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	ICP           string    `json:"value"`
	Copyright     string    `json:"copyright"`	
}