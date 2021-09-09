package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"yn/todo/internal/handler"
	"yn/todo/internal/repository"
	"yn/todo/internal/service"
	"yn/todo/pkg/database"
	"yn/todo/pkg/server"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := database.NewPostgresDB(database.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "pg",
		DBName:   "pg_todo_go",
		SSLMode:  "disable",
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	_repository := repository.NewRepository(db)
	_service := service.NewService(_repository)
	_handler := handler.NewHandler(_service)
	_server := new(server.Server)

	go func() {
		if err := _server.Run("8080", _handler.InitRoutes()); err != nil {
			logrus.Fatalf(
				"error occured while running http server: #{err.Error()}",
			)
		}
	}()
	logrus.Print("Server started.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Server shutting down.")
	if err := _server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
