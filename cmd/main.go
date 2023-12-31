package main

import (
	"os"

	"github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/handlers"
	"github.com/hel7/Atark-backend/pkg/repository"
	"github.com/hel7/Atark-backend/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Ініціалізація конфігурації
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Завантаження змінних середовища з файлу .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	// Ініціалізація підключення до бази даних MySQL
	db, err := repository.NewMysqlDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize db %s", err.Error())
	}

	repos := repository.NewRepository(db, repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   viper.GetString("db.dbname"),
	})

	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	// Ініціалізація сервера і запуск його на певному порту
	srv := new(farmsage.Server)
	logrus.Printf("Starting the server on port %s", viper.GetString("server.port"))

	if err := srv.Run(viper.GetString("server.port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error running server: %s", err.Error())
	}
}

// Функція для ініціалізації конфігурації
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.SetDefault("server.port", "8000")
	return nil
}
