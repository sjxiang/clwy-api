package database

import (
	"context"
	"database/sql"
)

// 根据 `课程编号` 查询所有章节
func (d *DB) GetAllChaptersByCourseId(ctx context.Context, courseId int64) ([]Chapter, error) {
	return nil, nil 
}


// 添加章节
func (d *DB) AddChapter(ctx context.Context, chapter *Chapter) error {
	return nil
}


func DeleteAllChaptersByCourseId(ctx context.Context, tx *sql.Tx, courseId int64) error {
	return nil
}