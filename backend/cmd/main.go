package main

import (
	"log"
	"user-center/internal/config"
	"user-center/internal/repository"
	"user-center/internal/router"
	"user-center/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := service.NewUserController(userService)

	r := gin.Default()
	router.RegisterUserRoutes(r, userController)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("服务启动成功，监听端口: %d", cfg.Server.Port)
	r.Run()
}
