package database

import (
	"context"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)



// return new_category.to_dict()
type AddCategoryParams struct {
	Name string `json:"name"`
	Rank int  `json:"rank"`
}

// 添加分类
func (d *DB) AddCategory(ctx context.Context, arg *AddCategoryParams) error {
	stmt := "INSERT INTO categories (`name`, `rank`) VALUES (?, ?)"

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()


	if _, err := d.db.ExecContext(ctx, stmt, arg.Name, arg.Rank); err != nil {
		var mysqlError *mysql.MySQLError

		if errors.As(err, &mysqlError) {
			has := strings.Contains(mysqlError.Message, "categories.idx_name")
			
			if has && mysqlError.Number == 1062{
				return ErrAlreadyExists
			}
		}

		return err
	}
	
	return nil
}


// 查询所有分类
func (d *DB) GetAllCategories(ctx context.Context) ([]Category, error) {
	stmt := "SELECT `id`, `name`, `rank` FROM categories ORDER BY `rank` ASC, `id` ASC"
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, stmt)
	if err!= nil {
		return nil, err
	}
	defer rows.Close()
	
	var items []Category
	
	for rows.Next() {
		var i Category
		
		err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Rank,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	if err:= rows.Err(); err!= nil {		
		return nil, err
	}

	return items, nil
}



// 删除分类 (✅ 事务)
func (d *DB) DeleteCategory(ctx context.Context, categoryId int64) error {
	
	stmt := `
		DELETE FROM categories
		WHERE id = ?
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	result, err := d.db.ExecContext(ctx, stmt, categoryId)
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

