package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)


type CreateCourseParams struct {
	CategoryId   int64
	UserId       int64

	Name         string
	Image        string
	Content      string
	
	Recommended  bool
	Introductory bool
}


// 添加课程
func (d *DB) CreateCourse(ctx context.Context, arg *CreateCourseParams) error {

	return withTx(d.db, ctx, func(tx *sql.Tx) error {
		
		// 检查分类是否存在
		exists, err := existsCategory(ctx, tx, arg.CategoryId)
		if err != nil {
			return err
		}
		if !exists {
			return ErrNoRecord
		}

		// 检查用户是否存在
		ok, err := existsUser(ctx, tx, arg.UserId)
		if err != nil {
			return err
		}
		if !ok {
			return ErrNoRecord
		}

		// 插入数据
		return insertCourse(ctx, tx, arg)
	})    	
}


func insertCourse(ctx context.Context, tx *sql.Tx, arg *CreateCourseParams) error {
	stmt := `
		INSERT INTO courses
			(category_id, name, user_id, image, recommended, introductory, content, likes_count, chapters_count, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	if _, err := tx.ExecContext(ctx, stmt, 
		arg.CategoryId, 
		arg.Name, 
		arg.UserId, 
		arg.Image, 
		arg.Recommended, 
		arg.Introductory, 
		arg.Content, 
		0, 
		0, 
		time.Now(), 
		nil); err != nil {
		var mysqlError *mysql.MySQLError

		if errors.As(err, &mysqlError) {
			has := strings.Contains(mysqlError.Message, "users.idx_name")
			
			if has && mysqlError.Number == 1062 {
				return ErrAlreadyExists
			}
		}

		return err
	}
	
	return nil
}


type FindAndCountAllCoursesParams struct {
	KeyWord       string
	Recommended   bool
	Introductory  bool
	Limit         int64
	Offset        int64
}

type CoursesPaginationResult struct {
	Courses     []Course `json:"courses"`
	Total       int64    `json:"total"`        // 总数据量
	TotalPages  int64    `json:"total_pages"`  // 总页数
	CurrentPage int64    `json:"current_page"` // 当前页码
}

// 查询课程列表 (✅ 重写一个结构体 ForExport)
func (d *DB) FindAndCountAllCourses(ctx context.Context, arg FindAndCountAllCoursesParams) (*CoursesPaginationResult, error) {

	stmt := `
	SELECT
		co.id, 
	    co.name, 
		co.image, 
		co.recommended, 
		co.introductory, 
		co.content, 
		co.likes_count, 
		co.chapters_count,
    	co.category_id, 
		co.user_id,
		co.created_at AS course_created, 
		co.updated_at AS course_updated,
		u.username,
    	ca.name AS category_name
	FROM 
		courses co
	LEFT JOIN 
		users u ON co.user_id = u.id
	LEFT JOIN 
		categories ca ON co.category_id = ca.id
	WHERE 
		co.name LIKE ? AND  -- 单个参数占位符
		co.recommended = ? AND
		co.introductory = ?
	LIMIT 
		?, ?
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	likeParam := "%" + arg.KeyWord + "%"  
	rows, err := d.db.QueryContext(ctx, stmt, likeParam, arg.Recommended, arg.Introductory, arg.Offset, arg.Limit)
	if err!= nil {
		return nil, err
	}
	defer rows.Close()

	var items []Course
	
	for rows.Next() {	
		var i Course

		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Image,
			&i.Recommended,
			&i.Introductory,
			&i.Content,
			&i.LikesCount,
			&i.ChaptersCount,
			&i.CategoryID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Author,
			&i.CategoryName,	
		); err != nil {
			return nil, err
		}
	
		items = append(items, i)
	}

	if err := rows.Err(); err!= nil {
		return nil, err
	}

	// TODO: 统计总数
	query := `
	SELECT 
		COUNT(*) 
	FROM 
		courses c
	LEFT JOIN 
		users u ON c.user_id = u.id
	LEFT JOIN 
		categories ca ON c.category_id = ca.id
	WHERE 
		c.name LIKE ? AND
		c.recommended = ? AND
		c.introductory = ?
	`

	var total int64
	if err = d.db.QueryRowContext(ctx, query, likeParam, arg.Recommended, arg.Introductory).Scan(&total); err != nil {
		return nil, err
	}

	// TODO: 计算总页数
	totalPages := total / arg.Limit
	if total%arg.Limit != 0 {
		totalPages++
	}

	// TODO: 计算当前页码
	currentPage := arg.Offset / arg.Limit + 1

	return &CoursesPaginationResult{
		Courses: items,
		Total: total,
		TotalPages: totalPages,
		CurrentPage: currentPage,
	}, nil
}



