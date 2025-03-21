package handlers

import (
	"fmt"
	"net/http"

	db "clwy-api/internal/database"
)

type ctxKey string


const (
	userKey ctxKey = "user"
)

// 从请求中获取当前用户信息

func getUserFromContext(r *http.Request) (*db.User, error) {
	v, ok := r.Context().Value(userKey).(*db.User)

	if !ok {
		return nil, fmt.Errorf("user not found")  // 类型断言失败
	}

	return v, nil 
}
