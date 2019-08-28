package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
)
const (
	paramUserID = "userId"
)
func GetUserFromAPI(c *gin.Context){



	userId := c.Param(paramUserID)

	/*
	id,err := strconv.Atoi(userId)
	if err != nil{
		apiErr := &utils.Error{
			Message: err.Error()
			Status :=http.StatusBadRequest,
		}
		c.JSON(apiErr.Status,apiErr)
	}
	*/


	response, apiErr := services.GetUser(userId)
	if apiErr != nil {
		c.JSON(apiErr.Status,apiErr)
		return
	}
	c.JSON(http.StatusOK,response)
}