package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bearl27/todos/back/internal/domain/model"
	"github.com/bearl27/todos/back/internal/infrastructure/persistence/postgres"
	"github.com/bearl27/todos/back/internal/interface/api/handler"
	"github.com/bearl27/todos/back/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate
	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	todoRepo := postgres.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := gin.Default()

	// CORS Setup
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:   []string{"Content-Length"},
	}))

	r.POST("/todos", todoHandler.Create)
	r.GET("/todos", todoHandler.GetAll)
	r.GET("/todos/:id", todoHandler.GetByID)
	r.PUT("/todos/:id", todoHandler.Update)
	r.DELETE("/todos/:id", todoHandler.Delete)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
