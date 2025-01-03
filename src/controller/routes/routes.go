package routes

import(
	"github.com/interlucas/Golang_CRUD/src/controller"
	"github.com/gin-gonic/gin"

)
func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserById/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUserd/:userId")
}