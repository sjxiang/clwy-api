package database

import (
	"context"
	"database/sql"
)


type GetAllChaptersParams struct {
	CourseId   int64  `json:"course_id"`
	Limit      int64  `json:"limit"`
	Offset     int64  `json:"offset"`  // skip
}

type ChaptersPaginationResult struct {
	Chapters    []Chapter`json:"chapters"`
	Total       int64    `json:"total"`        // 总数据量
	TotalPages  int64    `json:"total_pages"`  // 总页数
	CurrentPage int64    `json:"current_page"` // 当前页码
}


// 查询章节列表
func (d *DB) GetAllChapters(ctx context.Context, arg GetAllChaptersParams) (*ChaptersPaginationResult, error) {
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	stmt := `
		SELECT
			c.id,
			c.title,
			c.content,
			c.video, 
			c.rank,
			c.course_id,
			co.name,
			c.created_at,
			c.updated_at
		FROM
			chapters c
		LEFT JOIN
			courses co ON c.course_id = co.id
		WHERE
			c.course_id = ?
		LIMIT
			?,?
	`

	rows, err := d.db.QueryContext(ctx, stmt, arg.CourseId, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Chapter
	
	for rows.Next() {	
		var i Chapter

		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Video,
			&i.Rank,
			&i.CourseID,
			&i.CourseName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
	
		items = append(items, i)
	}

	if err := rows.Err(); err!= nil {
		return nil, err
	}


	// 统计数据
	countStmt := `
		SELECT
			COUNT(*)
		FROM
			chapters c
		LEFT JOIN
			courses co ON c.course_id = co.id
		WHERE
			course_id = ?
	`
	var total int64
	err = d.db.QueryRowContext(ctx, countStmt, arg.CourseId).Scan(&total)
	if err!= nil {
		return nil, err
	}

	// 计算总页数
	totalPages := total / arg.Limit
	if total % arg.Limit != 0 {
		totalPages++
	}

	// TODO: 计算当前页码
	currentPage := arg.Offset / arg.Limit + 1

	return &ChaptersPaginationResult{
		Chapters: items,
		Total: total,
		TotalPages: totalPages,
		CurrentPage: currentPage,
	}, nil 
}


// 添加章节
func (d *DB) AddChapter(ctx context.Context, chapter *Chapter) error {
	return nil
}


func DeleteAllChaptersByCourseId(ctx context.Context, tx *sql.Tx, courseId int64) error {
	return nil
}