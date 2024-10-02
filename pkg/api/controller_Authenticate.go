package api

import (
	gin "github.com/gin-gonic/gin"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)

// @Summary Authenticate user
// @ID Authenticate
// @Produce json
// @Param data body InputAuthData true "Auth input data"
// @Success 200 {object} OutputAuthData
// @Failure 400 {object} OutputAuthData
// @Router /authenticate [post]
func Authenticate(c *gin.Context){

	
}