package api

import (
	"fmt"
	"net/http"

	utils "github.com/bbitere/gameapi.git/pkg/utils"
	gin "github.com/gin-gonic/gin"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)


type InputAuthenticate struct{

	Token 			string 		`json:"sessionUID"`
}
type OutputAuthenticate struct{

	UserID			string 		`json:"userID"`
	Wallet  		string  	`json:"wallet"`
	Currency  		string  	`json:"currency"`
	Error  			int    		`json:"error"`
	Description 	string  	`json:"description"`

}
// @Summary Authenticate user
// @ID Authenticate
// @Produce json
// @Param data body InputAuthenticate true "Auth input data"
// @Success 200 {object} OutputAuthenticate
// @Failure 400 {object} OutputAuthenticate
// @Router /authenticate [post]
func Authenticate(c *gin.Context){

	var inputData = InputAuthenticate{};
	//var outData   = OutputAuthenticate{};

	utils.Log_log("Authenticate", "", 1, fmt.Sprint(inputData), "" )

	var err = c.BindJSON(&inputData);
	if( err != nil){

		utils.Log_log("Authenticate", "", 1, fmt.Sprint(err), "Cannot read Token" )
		c.IndentedJSON( http.StatusOK, err)
	}
	var outData = 	OutputAuthenticate{
					Currency : "RON",
					UserID: "244",
					Wallet: "1000",
					Error: 0,
					Description: "",
				};


	c.IndentedJSON(http.StatusOK, outData)
}

