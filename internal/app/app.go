package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"clwy-api/internal/auth"
	db "clwy-api/internal/database"
	"clwy-api/internal/handlers"
	"clwy-api/internal/logger"
)

type App struct {
	Env *Env
}

type Env struct {
	API_SERVER_PORT string

	DB_CONNECTION_STRING string

	AUTH_JWT_SECRET_KEY string
	AUTH_JWT_ISSUER string
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	a.SetupEnv()
	a.SetupServer()
}

func (a *App) SetupEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	a.Env = &Env{
		API_SERVER_PORT: os.Getenv("API_SERVER_PORT"),

		DB_CONNECTION_STRING: os.Getenv("DB_CONNECTION_STRING"),

		AUTH_JWT_SECRET_KEY: os.Getenv("AUTH_JWT_SECRET_KEY"),
		AUTH_JWT_ISSUER: os.Getenv("AUTH_JWT_ISSUER"),	
	}
}

func (a *App) SetupServer() {
	
	fmt.Println(a.Env)
	
	logger, err := logger.New("gua")
	if err!= nil {
		log.Fatal(err)
	}

	jwtAuthenticator := auth.NewJWTAuthenticator(
		a.Env.AUTH_JWT_SECRET_KEY,
		a.Env.AUTH_JWT_ISSUER,
	)

	db, err := db.New(a.Env.DB_CONNECTION_STRING)
	if err!= nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	logger.Infow("database connected")
	
	h := handlers.New(db, logger, jwtAuthenticator)
	h.SetupRoutes()

	logger.Infow("server has started", "addr", a.Env.API_SERVER_PORT)
	
	err = h.StartServer()
	if err != nil {
		log.Printf("Error Starting Server:\nMessage:\n%v", err.Error())
	}
}

