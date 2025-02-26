package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)


func (d *DB) GetNotice(ctx context.Context, id int64) (Notice, error) {
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM notices 
		WHERE id = ?
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	row := d.db.QueryRowContext(ctx, query, id)
	
	var i Notice

	err := row.Scan(
		&i.ID, 
		&i.Title, 
		&i.Content, 
		&i.CreatedAt, 
		&i.UpdatedAt,
	)
	
	if err!= nil {
		switch err {
		case sql.ErrNoRows:
			return Notice{}, ErrNotFound
		default:
			return Notice{}, err
		}
	}

	return i, nil
}


func (d *DB) DeleteNotice(ctx context.Context, id int64) error {
	stmt := `
		DELETE FROM notices 
		WHERE id = ?
	`
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	result, err := d.db.ExecContext(ctx, stmt, id)
	if err!= nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err!= nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}
	
	return nil 	
}


type CreateNoticeParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (d *DB) CreateNotice(ctx context.Context, arg *CreateNoticeParams) (int64, error) {

	stmt := `
		INSERT INTO notices 
			(title, content) 
		VALUES 
			(?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	result, err := d.db.ExecContext(ctx, stmt, arg.Title, arg.Content)
	if err!= nil {
		return 0, err
	}
	
	id , err := result.LastInsertId()
	if err!= nil {
		return 0, err
	}
	
	return id, nil
}

type UpdateNoticeParams struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (d *DB) UpdateNotice(ctx context.Context, arg *UpdateNoticeParams) error {

	stmt := `
		UPDATE notices 
		SET title = ?, content = ?
		WHERE id = ?;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	result, err := d.db.ExecContext(ctx, stmt, arg.Title, arg.Content, arg.ID)
	if err!= nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err!= nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}


type GetNoticesWithPaginationParams struct {
	SearchKey  string `json:"title"`
	Limit      int64  `json:"limit"`
	Offset     int64  `json:"offset"`
}

type PaginationResult struct {
	Data        []Notice `json:"data"`
	Total       int64    `json:"total"`        // 总数据量
	TotalPages  int64    `json:"total_pages"`  // 总页数
	CurrentPage int64    `json:"current_page"` // 当前页码
}


// 获取分页数据及元信息
func (d *DB) GetNoticesWithPagination(ctx context.Context, arg *GetNoticesWithPaginationParams) (*PaginationResult, error) {
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// SQL 1
	query, args := buildDynamicSQL(`
		SELECT id, title, content, created_at, updated_at
		FROM notices
	`, arg)

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Notice

	for rows.Next() {
		var i Notice

		err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err!= nil {
			return nil, err
		}

		items = append(items, i)
	}
	if err = rows.Err(); err!= nil {
		return nil, err
	}


	// SQL 2
	countQuery, countArgs := buildDynamicSQL(`
		SELECT COUNT(*)
		FROM notices
	`, arg)

	var total int64
	if err = d.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total); err != nil {
		return nil, err
	}

	// 计算总页数
	totalPages := total / arg.Limit
	if total%arg.Limit != 0 {
		totalPages++
	}

	return &PaginationResult{
		Data:        items,
		Total:       total,
		TotalPages:  totalPages,
		CurrentPage: arg.Offset / arg.Limit + 1,
	}, nil
}


// 动态构建 SQL
func buildDynamicSQL(basicSQL string, arg *GetNoticesWithPaginationParams) (string, []interface{}) {
	
	var query strings.Builder

	args := make([]interface{}, 0)

	// 构建基本查询
	query.WriteString(basicSQL)

	// 模糊搜索
	if arg.SearchKey != "" {
		query.WriteString("WHERE title LIKE ? ")
		args = append(args, "%" + arg.SearchKey + "%")
	}

	// 排序和分页
	query.WriteString("ORDER BY updated_at DESC ")
	query.WriteString("LIMIT ?, ?")
	args = append(args, arg.Offset, arg.Limit)
	
	fmt.Println(query.String())
	fmt.Println(args)
	return query.String(), args
}
