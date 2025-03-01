package database

import "context"


// 添加课程
func (d *DB) CreateCourse(ctx context.Context, course *Course) (int64, error) {
	return 0, nil
}

// 按 `分类编号` 查询所有课程
func (d *DB) GetAllCoursesByCategoryId(ctx context.Context, categoryId int64) ([]Course, error) {
	return nil, nil
}