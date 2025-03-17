package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)


/*
统计每个月的注册用户数量

返回值：
- 一个字典，其中键是月份，值是该月份的用户注册数量

SELECT DATE_FORMAT(`created_at`, '%Y-%m') AS `month`,
COUNT(*) AS `count`
FROM `users`
GROUP BY `month`
ORDER BY `month` ASC;

{
	"2023-01": 2,
	"2023-02": 1,
	"2023-03": 1,
	"2023-09": 1,
}

*/     

func (d *DB) CountMonthlyUserRegistrations(ctx context.Context) (map[string]int64, error) {
	query := `
		SELECT 
			DATE_FORMAT(created_at, '%Y-%m') AS month,
			COUNT(*) AS count
		FROM 
			users
		GROUP BY 
			month
		ORDER BY 
			month ASC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int64)
	
	for rows.Next() {
		var (
			month string
			count int64
		)
	
		err := rows.Scan(&month, &count)
		if err != nil {
			return nil, err
		}
	
		result[month] = count
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}


// 统计用户性别
func (d *DB) CountUserGenders(ctx context.Context) (map[string]int64, error) {
	query := `
		SELECT 
			sex as gender,
			COUNT(*) AS count
		FROM 
			users
		GROUP BY 
			gender
		ORDER BY 
			gender ASC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int64)

	for rows.Next() {
		var (
			gender string
			count int64
		)

		err := rows.Scan(&gender, &count)
		if err != nil {
			return nil, err
		}

		result[gender] = count
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}


type AddUserParams struct {
	Email    string 
    Username string 
    Nickname string 
    Password string 
    Avatar   string 

    Sex      int8   
	Role     int8  
    Company  string 
    Intro    string 
}

// 添加用户
func (d *DB) AddUser(ctx context.Context, arg AddUserParams) error {
	stmt := `
		INSERT INTO users 
			(email, username, nickname, password, avatar, sex, role, company, intro) 
		VALUES 
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()


	if _, err := d.db.ExecContext(ctx, stmt, 
		arg.Email, 
		arg.Username, 
		arg.Nickname, 
		arg.Password, 
		arg.Avatar, 
		arg.Sex, 
		arg.Role, 
		arg.Company, 
		arg.Intro); err != nil {

		var mysqlError *mysql.MySQLError

		if errors.As(err, &mysqlError) {
			switch {
			case mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users.idx_email"):
				return ErrDuplicateEmail
			case mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users.idx_username"):
				return ErrDuplicateUsername
			}
		}

		return err
	}
	
	return nil
}


type saveUserParams struct {
	
}
func (d *DB) SaveUser(ctx context.Context, arg saveUserParams) error {
	return nil
}


// 查询用户
func (d *DB) GetUser(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT 
			email, username, nickname, password, avatar, sex, company, intro, role, created_at, updated_at
		FROM 
			users
		WHERE 
			id = ?
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	row := d.db.QueryRowContext(ctx, query, id)
	
	var item User

	err := row.Scan(
		&item.Email,
		&item.Username,
		&item.Nickname,
		&item.Password,
		&item.Avatar,
		&item.Sex,
		&item.Company,
		&item.Intro,
		&item.Role,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		
		return nil, err
	}

	return &item, nil
}

func (d *DB) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT 
			email, username, nickname, password, avatar, sex, company, intro, role, created_at, updated_at 
		FROM 
			users
		WHERE 
			email = ?
		LIMIT 
			1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	row := d.db.QueryRowContext(ctx, query, email)
	
	var item User

	err := row.Scan(
		&item.Email,
		&item.Username,
		&item.Nickname,
		&item.Password,
		&item.Avatar,
		&item.Sex,
		&item.Company,
		&item.Intro,
		&item.Role,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		
		return nil, err
	}

	return &item, nil	
}

func (d *DB) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	query := `
		SELECT 
			email, username, nickname, password, avatar, sex, company, intro, role, created_at, updated_at 
		FROM 
			users
		WHERE 
			username = ?
		LIMIT
			1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	row := d.db.QueryRowContext(ctx, query, username)
	
	var item User

	err := row.Scan(
		&item.Email,
		&item.Username,
		&item.Nickname,
		&item.Password,
		&item.Avatar,
		&item.Sex,
		&item.Company,
		&item.Intro,
		&item.Role,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		
		return nil, err
	}

	return &item, nil
}


func (d *DB) AllUsers(ctx context.Context) ([]User, error) {
	return nil, nil
}