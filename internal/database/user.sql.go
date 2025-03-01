package database

import "context"


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
	return nil, nil
}


// 统计用户性别
func (d *DB) CountUserGenders(ctx context.Context) (map[string]int64, error) {
	return nil, nil
}


type CreateUserParams struct {
	
}

// 添加用户
func (d *DB) AddUser(ctx context.Context, arg *CreateUserParams) error {
	return nil
}