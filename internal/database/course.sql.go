package database

import (
	"time"
	"context"
	"errors"
	"database/sql"
)


type CreateCourseParams struct {
	CategoryId   int64
	Name         string
	UserId       int64
	Image        string
	Recommended  bool
	Introductory bool
	Content      string	
}


// 添加课程
func (d *DB) CreateCourse(ctx context.Context, arg *CreateCourseParams) error {
	stmt := `
		INSERT INTO courses
			(category_id, name, user_id, image, recommended, introductory, content, likes_count, chapters_count, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	result, err := d.db.ExecContext(ctx, stmt, 
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
		nil)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err!= nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrAlreadyExists
	}

	return nil
}




// 按 `分类编号` 查询所有课程 (✅)
func (d *DB) GetAllCoursesByCategoryId(ctx context.Context, categoryId int64) ([]Course, error) {

	stmt := `
		SELECT id, category_id, user_id, name, image, recommended, introductory, content, likes_count, chapters_count, created_at, updated_at
		FROM courses
		WHERE category_id = ?
		ORDER BY id ASC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	rows, err := d.db.QueryContext(ctx, stmt, categoryId)
	if err!= nil {
		return nil, err
	}
	defer rows.Close()

	var items []Course
	
	for rows.Next() {	
		var i Course
	
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.UserID,
			&i.Name,
			&i.Image,
			&i.Recommended,
			&i.Introductory,
			&i.Content,
			&i.LikesCount,
			&i.ChaptersCount,
			&i.CreatedAt,
			&i.UpdatedAt,
		); 	err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ErrNotFound
			default:
				return nil, err
			}
		}
	
		items = append(items, i)
	}

	if err := rows.Err(); err!= nil {
		return nil, err
	}

	return items, nil
}


