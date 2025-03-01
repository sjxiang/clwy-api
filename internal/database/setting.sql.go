package database

import (
	"context"
	"database/sql"
)


type UpdateSettingParams struct {
	Name          string   
	ICP           string   
	Copyright     string   
}

// 更改系统设置
func (d *DB) UpdateSetting(ctx context.Context, arg UpdateSettingParams) error {
	stmt := `
		UPDATE settings 
		SET name =?, icp =?, coptyright =?
		WHERE id = 1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	
	result, err := d.db.ExecContext(ctx, stmt, arg.Name, arg.ICP, arg.Copyright)
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


// 查找系统设置, 写死了
func (d *DB) GetSetting(ctx context.Context) (Setting, error) {
	stmt := `
		SELECT name, icp, coptyright 
		FROM settings 
		WHERE id = 1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	row := d.db.QueryRowContext(ctx, stmt)
	
	var i Setting

	err := row.Scan(
		&i.Name,
		&i.ICP,
		&i.Copyright,
	)
	
	if err!= nil {
		switch err {
		case sql.ErrNoRows:
			return Setting{}, ErrNotFound
		default:
			return Setting{}, err
		}
	}

	return i, nil
}


