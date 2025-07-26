package main

import (
	"context"
	"github.com/emomovg/todo-app/configs"
	"github.com/emomovg/todo-app/internal/repository"
	"github.com/emomovg/todo-app/internal/routes"
	"github.com/emomovg/todo-app/internal/services"
	"github.com/emomovg/todo-app/pkg/db"
	"github.com/emomovg/todo-app/pkg/handler"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}

	if err := configs.Load(); err != nil {
		log.Fatalf("error initialization configs %s", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	postres, err := db.New(ctx, db.PostgresConfig{
		Host:          viper.GetString("db.host"),
		Port:          viper.GetInt("db.port"),
		Username:      viper.GetString("db.username"),
		Password:      os.Getenv("DB_PASSWORD"),
		DBName:        viper.GetString("db.DBName"),
		SSLMode:       viper.GetString("db.SSLMode"),
		DBMaxCons:     viper.GetInt("db.DBMaxCons"),
		DBMinCons:     viper.GetInt("db.DBMinCons"),
		DBMaxLifetime: viper.GetString("db.DBMaxLifetime"),
		DBMaxIdleTime: viper.GetString("db.DBMaxIdTime"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer postres.Close()

	repo := repository.NewRepository(postres)
	service := services.NewService(repo.IUserRepository)
	router := routes.NewRouter(service)
	handler := new(handler.Handler)

	if err := handler.Run(viper.GetString("port"), router.InitRoutes()); err != nil {
		log.Fatalf("error occuped while running http server: %s", err.Error())
	}
}
