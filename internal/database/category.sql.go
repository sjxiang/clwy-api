package database

import "context"


// 添加分类
func (d *DB) AddCategory(ctx context.Context) error {
	return nil 
}

// 查询所有分类
func (d *DB) GetAllCategories(ctx context.Context) ([]*Category, error) {
	return nil, nil
}

// 删除分类
func (d *DB) DeleteCategory(ctx context.Context) error {
	return nil 
}