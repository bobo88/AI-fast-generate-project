package router

import (
	"backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userController *controller.UserController) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/create", userController.Create)
		userGroup.POST("/update", userController.Update)
		userGroup.POST("/delete", userController.Delete)
		userGroup.GET("/detail", userController.Detail)
		userGroup.GET("/list", userController.List)

		userGroup.POST("/update-password", userController.UpdatePassword)
		userGroup.POST("/assign-roles", userController.AssignRoles)
	}
}
