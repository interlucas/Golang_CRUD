package controller

import (
	"github.com/interlucas/Golang_CRUD/src/configuration/rest_err"
	"github.com/interlucas/Golang_CRUD/src/controller/model"
	"github.com/gin-gonic/gin"
	"fmt"
)
	func CreateUser(c *gin.Context){		
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest): err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error))
		c.JSON(restErr.Code, restErr)
		return
			}
	fmt.Println(userRequest)
}