package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	QueryTimeoutDuration = time.Second * 5
)


type DB struct {
	db *sql.DB
}


func New(connStr string) (*DB, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	                                                        
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil 
}

func (d *DB) HealthCheck() error {
	return d.db.Ping()
}


/*

FindAll
FindByPk
FindAndCountAll

Destroy



查询
修改
新增
删除


 */

 