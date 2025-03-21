package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"

	"clwy-api/internal/auth"
	db "clwy-api/internal/database"
)

type Handler struct {
	router *chi.Mux
	logger *zap.SugaredLogger
	db     *db.DB
	authn   auth.Authenticator
}

func New(db *db.DB, logger *zap.SugaredLogger, authn auth.Authenticator) *Handler {
	return &Handler{
		router: chi.NewRouter(),
		db: db,
		logger: logger,
		authn: authn,
	}
}

func (h *Handler) SetupRoutes() {
	
	h.router.Use(middleware.RequestID)
	h.router.Use(middleware.RealIP)
	h.router.Use(middleware.Logger)
	h.router.Use(middleware.Recoverer)
	
	h.router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{("http://localhost:10005")},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		},
	))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	h.router.Use(middleware.Timeout(60 * time.Second))


	// 后台路由
	h.router.Route("/admin", func(r chi.Router) {
		r.Get("/health", h.HealthCheck)

		r.Route("/notices", func(r chi.Router) {
			
			r.Use(h.AuthTokenMiddleware)

			// 公告列表
			r.Get("/", h.AllNotices)  
			// 公告详情
			r.Get("/{id}", h.GetNotice) 
			// 创建公告
			r.Post("/", h.CreateNotice)
			// 删除公告
			r.Delete("/{id}", h.DeleteNotice)	
			// 更新公告
			r.Put("/{id}", h.UpdateNotice)
			
		})

		r.Route("/categories", func(r chi.Router) {
			
			// 分类列表
			r.Get("/", h.AllCategories)  
			// 创建分类
			r.Post("/", h.AddCategory)
			// 删除分类
			r.Delete("/{id}", h.DeleteCategory)	
			// 查询分类详情
			r.Get("/{id}", h.GetCategory)
		})

		r.Route("/courses", func(r chi.Router) {
			// 添加课程
			r.Post("/", h.AddCourse)
			// 删除课程

			// 课程列表
			r.Get("/", h.AllCourses)

			// 
		})

		r.Route("/chapters", func(r chi.Router) {
			r.Get("/", h.AllChapters)
		})

		r.Route("/users", func(r chi.Router) {
			// 添加用户
			r.Post("/", h.AddUser)
			// 查询用户列表
			r.Get("/", h.AllUsers)
			// 查询用户
			r.Get("/{id}", h.GetUser)
		})

		r.Route("/settings", func(r chi.Router) {
			// 编辑系统设置
			r.Put("/", h.UpdateSetting)
			// 查找当前系统设置
			r.Get("/", h.GetSetting)	
		})

		r.Route("/echarts", func(r chi.Router) {
			r.Get("/sex", h.CountGenders)
			r.Get("/user", h.CountUser)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign_in", h.SignIn)
		})

	
	})

	// public route
	
}

func (h *Handler) StartServer() error {
	err := http.ListenAndServe(":10005", h.router)
	if err!= nil {
		return err
	}
	return nil
}


