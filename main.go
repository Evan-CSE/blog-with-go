package main

import (
	"blog-api/model"
	"blog-api/repository"
	"blog-api/router"
	"blog-api/service"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func goDotEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file")
	}

	return os.Getenv(key)
}

func main() {
	dsn := "host=" + goDotEnv("DB_HOST") + " user=" + goDotEnv("DB_USER") + " password=" + goDotEnv("DB_PASSWORD") + " dbname=" + goDotEnv("DB_NAME") + " port=" + goDotEnv("DB_PORT") + " sslmode=disable TimeZone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if isProd := goDotEnv("IS_PROD"); isProd != "true" {
		db.AutoMigrate(&model.Blog{})
	}

	if err != nil {
		panic("failed to connect database")
	}

	repo := repository.NewBlogRepository(db)
	service := service.NewBlogService(repo)
	router := router.SetupRouter(service)
	router.Run()
}
