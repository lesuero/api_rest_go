package controllers


import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"../circuitbraker"
)

func GetResultFromAPI(c *gin.Context){

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

  //arreglar int

	/*userIdInt, err := strconv.Atoi(userId)
	if err == nil {
		fmt.Println(userIdInt)
	}*/

	if circuitbraker.CircuitResult.State != "Closed"{
		c.JSON(http.StatusInternalServerError, "Servidor no disponible, intente nuevamente más tarde")
		return
	}
	response, apiErr := services.GetResult(userId)


	if apiErr != nil {
		c.JSON(apiErr.Status,apiErr)
		circuitbraker.CircuitResult.AddError()
		return
	}
	c.JSON(http.StatusOK,response)
}
