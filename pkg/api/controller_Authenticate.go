package api

//utils "github.com/bbitere/gameapi.git/pkg/utils"


type AuthenticateInput struct{

	Token 			string 		`json:"sessionUID"`
}
type AuthenticateResponse struct{

	UserID			string 		`json:"userID"`
	Wallet  		string  	`json:"wallet"`
	Currency  		string  	`json:"currency"`
	Error  			int    		`json:"error"`
	Description 	string  	`json:"description"`
}
type AuthenticateResponseErr = AuthenticateResponse
// @Summary Authenticate user
// @ID Controller_Authenticate
// @Produce json
// @Param data body AuthenticateInput true "Auth input data"
// @Success 200 {object} AuthenticateResponse
// @Failure 400 {object} AuthenticateResponseErr
// @Router /authenticate [post]
func Controller_Authenticate(inp *AuthenticateInput) (*AuthenticateResponse, *AuthenticateResponseErr, error){

	var outData = 	AuthenticateResponse{
					Currency : "RON",
					UserID: "244",
					Wallet: "1000",
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;	
}

