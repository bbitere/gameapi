package api

type PlayeNameInput struct {
	PlayerName string `json:"playerName"`
	IsHidden   bool   `json:"isHidden"`
	Currency   string `json:"currency"`
	Token      string `json:"sessionUID"`
}
type PlayeNameResponse struct {
	PlayerToken string `json:"playerToken"`
	Error       int    `json:"error"`
	Description string `json:"description"`
}
type PlayeNameResponseErr = PlayeNameResponse

// @Summary PlayeName user
// @ID Controller_PlayeName
// @Produce json
// @Param data body PlayeNameInput true "Auth input data"
// @Success 200 {object} PlayeNameResponse
// @Failure 400 {object} PlayeNameResponseErr
// @Router /PlayeName [post]
func Controller_PlayeName(inp *PlayeNameInput) (*PlayeNameResponse, *PlayeNameResponseErr, error) {

	var token, err = API.GameLogic.addPlayer(inp.PlayerName, inp.IsHidden, inp.Currency, true)
	if err != nil {

		var outData = PlayeNameResponse{
			Error:       1,
			Description: err.Error(),
			PlayerToken: token,
		}

		return &outData, nil, err
	} else {

		var outData = PlayeNameResponse{
			Error:       0,
			Description: "",
			PlayerToken: token,
		}

		return &outData, nil, nil
	}
}
