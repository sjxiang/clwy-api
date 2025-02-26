package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"

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
	}
}

func (a *App) SetupServer() {
	
	logger, err := logger.New("gua")
	if err!= nil {
		log.Fatal(err)
	}

	db, err := db.New(a.Env.DB_CONNECTION_STRING)
	if err!= nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	logger.Infow("database connected")
	
	h := handlers.New(db, logger)
	h.SetupRoutes()

	logger.Infow("server has started", "addr", a.Env.API_SERVER_PORT)
	
	err = h.StartServer()
	if err != nil {
		log.Printf("Error Starting Server:\nMessage:\n%v", err.Error())
	}
}

