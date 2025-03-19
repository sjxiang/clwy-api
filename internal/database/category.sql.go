package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)


type AddCategoryParams struct {
	Name string `json:"name"`
	Rank int8  `json:"rank"`
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
	return withTx(d.db, ctx, func(tx *sql.Tx) error {
	
	/*
	
	1. 外键约束
	2. 删除分类的同时，删除所有关联课程
	3. 只有没有关联课程的分类，才能被删除
	
	
	 */
		if err := countCourse(ctx, tx, categoryId); err != nil {
			return err
		}
	
		if err := deleteCategory(ctx, tx, categoryId); err != nil {
			return err
		}

		return nil 
	})
}

func countCourse(ctx context.Context, tx *sql.Tx, categoryId int64) error {

	stmt := `
		SELECT COUNT(*) FROM courses WHERE category_id =?
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var count int64
	err := tx.QueryRowContext(ctx, stmt, categoryId).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrAlreadyExists
	}

	return nil
}

func deleteCategory(ctx context.Context, tx *sql.Tx, categoryId int64) error {
	
	stmt := `
		DELETE FROM categories
		WHERE id = ?
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	result, err := tx.ExecContext(ctx, stmt, categoryId)
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


func existsCategory(ctx context.Context, tx *sql.Tx, id int64) (bool, error) {
	
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM categories WHERE id = ?)"
	
	err := tx.QueryRowContext(ctx, stmt, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, err
}






// 查询分类详情 (✅)
func (d *DB) GetCategory(ctx context.Context, id int64) (*Category, error) {

	stmt := `
	SELECT 
		c.id AS category_id, 
		c.created_at AS category_created, 
		c.updated_at AS category_updated,
		c.name AS category_name, 
		c.rank,
		co.id AS course_id, 
		co.user_id, 
		co.name AS course_name, 
		co.image, 
		co.recommended,
		co.introductory, 
		co.content, 
		co.likes_count, 
		co.chapters_count, 
		co.created_at AS course_created, 
		co.updated_at AS course_updated
	FROM 
		categories c
	LEFT JOIN 
		courses co ON co.category_id = c.id
	WHERE
		c.id = ?
	ORDER BY 
		c.rank DESC, c.id, co.created_at DESC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	rows, err := d.db.QueryContext(ctx, stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var i Category
	
	for rows.Next() {	

		var co Course
		
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Rank,
			&co.ID,
			&co.UserID,
			&co.Name,
			&co.Image,
			&co.Recommended,
			&co.Introductory,
			&co.Content,
			&co.LikesCount,
			&co.ChaptersCount,
			&co.CreatedAt,
			&co.UpdatedAt,
		); err != nil {
			return nil, err
		}

		i.Coureses = append(i.Coureses, co)
	}

	if err := rows.Err(); err!= nil {
		return nil, err
	}

	return &i, nil
}
