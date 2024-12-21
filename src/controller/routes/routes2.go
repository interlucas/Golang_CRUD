package routes2

import (
	"github.com/gin-gonic/gin"
	"github.com/interlucas/Golang_CRUD.git/src/controller"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
