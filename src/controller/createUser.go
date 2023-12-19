package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/azevedoedison/golang_crud/src/configuration/validation"
	"github.com/azevedoedison/golang_crud/src/controller/model/request"
	"github.com/azevedoedison/golang_crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	log.Println("Init CreateUser Controller")
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		log.Printf("Error tring to marshal object, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}
