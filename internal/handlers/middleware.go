package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"clwy-api/internal/auth"
	db "clwy-api/internal/database"
)

// 认证, 例 'Authenticate'
func (h *Handler) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// 当前接口需要认证才能访问
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			h.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			h.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
			return
		}


		accessToken := parts[1]
		// 验证
		userId, err := h.authn.ValidateToken(accessToken)
		if err != nil {
			if errors.Is(err, auth.ErrTokenExpiry) {
				h.unauthorizedErrorResponse(w, r, err)
				return
			}

			h.internalServerError(w, r, err)
			return
		}


		ctx := r.Context()

		user, err := h.db.GetUser(ctx, userId)
		if err != nil {
			if errors.Is(err, db.ErrNoRecord) {
				h.unauthorizedErrorResponse(w, r, fmt.Errorf("用户不存在"))
				return
			}

			h.internalServerError(w, r, err)
			return
		}
		

		// 将用户信息放入上下文中
		ctx = context.WithValue(ctx, userKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


