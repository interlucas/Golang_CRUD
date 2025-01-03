package routes

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserById/:userId", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.DeleteUser)
	r.DELETE("/deleteUserd/:userId")
}