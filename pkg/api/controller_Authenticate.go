package api

//utils "github.com/bbitere/gameapi.git/pkg/utils"


type AuthenticateInput struct{

	Token 			string 		`json:"token"`
	Currency 		string 		`json:"currency"`
	Language 		string 		`json:"language"`
}
type AuthenticateResponse struct{

	AuthSessionID		string 		`json:"AuthSessionID"`
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
					AuthSessionID: "244-5442352-234234",
					Wallet: "1000",
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;	
}

