// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/PlayeName": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "PlayeName user",
                "operationId": "Controller_PlayeName",
                "parameters": [
                    {
                        "description": "Auth input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PlayeNameInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.PlayeNameResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.PlayeNameResponseErr"
                        }
                    }
                }
            }
        },
        "/Player": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Play user",
                "operationId": "Controller_PlayerList",
                "parameters": [
                    {
                        "description": "Play input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PlayerListInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.PlayerListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.PlayerListResponseErr"
                        }
                    }
                }
            }
        },
        "/authenticate": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Authenticate user",
                "operationId": "Controller_Authenticate",
                "parameters": [
                    {
                        "description": "Auth input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthenticateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.AuthenticateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.AuthenticateResponseErr"
                        }
                    }
                }
            }
        },
        "/cashout": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "do cashot for user",
                "operationId": "Controller_Cashout",
                "parameters": [
                    {
                        "description": "Play input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CashoutInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.CashoutResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.CashoutResponseErr"
                        }
                    }
                }
            }
        },
        "/gameupdate": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "update the entire game for all user",
                "operationId": "Controller_GameUpdate",
                "parameters": [
                    {
                        "description": "Play input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GameUpdateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GameUpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.GameUpdateResponseErr"
                        }
                    }
                }
            }
        },
        "/history": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Play user",
                "operationId": "Controller_HistoryList",
                "parameters": [
                    {
                        "description": "Play input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.HistoryListInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.HistoryListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.HistoryListResponseErr"
                        }
                    }
                }
            }
        },
        "/play": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Play user",
                "operationId": "Controller_Play",
                "parameters": [
                    {
                        "description": "Play input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PlayInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.PlayResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.PlayResponseErr"
                        }
                    }
                }
            }
        },
        "/playautobet": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Play user",
                "operationId": "Controller_PlayAutobet",
                "parameters": [
                    {
                        "description": "Play input data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PlayAutobetInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.PlayAutobetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.PlayAutobetResponseErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AuthenticateInput": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api.AuthenticateResponse": {
            "type": "object",
            "properties": {
                "AuthSessionID": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.AuthenticateResponseErr": {
            "type": "object",
            "properties": {
                "AuthSessionID": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.CashoutInput": {
            "type": "object",
            "properties": {
                "StopOnLoss": {
                    "type": "string"
                },
                "bet": {
                    "type": "string"
                },
                "multiplBashout": {
                    "type": "string"
                },
                "numOfAutobet": {
                    "type": "integer"
                },
                "playerToken": {
                    "type": "string"
                },
                "sessionUID": {
                    "type": "string"
                },
                "stopOnProfit": {
                    "type": "string"
                }
            }
        },
        "api.CashoutResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                }
            }
        },
        "api.CashoutResponseErr": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.GameUpdateInput": {
            "type": "object",
            "properties": {
                "playerToken": {
                    "type": "string"
                },
                "sessionUID": {
                    "type": "string"
                }
            }
        },
        "api.GameUpdateResponse": {
            "type": "object",
            "properties": {
                "announce": {
                    "$ref": "#/definitions/api.TypeCrashArg_Annonce"
                },
                "betStatus": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "gameState": {
                    "type": "string"
                },
                "msg": {
                    "$ref": "#/definitions/api.TypeCrashArg_MsgWinOrLose"
                },
                "playerlist": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.PlayerItem"
                    }
                },
                "playing": {
                    "$ref": "#/definitions/api.TypeCrashArg_Playing"
                }
            }
        },
        "api.GameUpdateResponseErr": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.HistoryItem": {
            "type": "object",
            "properties": {
                "bet": {
                    "type": "string"
                },
                "betCurrency": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "details": {
                    "type": "string"
                },
                "multiplicator": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "win": {
                    "type": "string"
                }
            }
        },
        "api.HistoryListInput": {
            "type": "object",
            "properties": {
                "sessionUID": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "startOffset": {
                    "type": "integer"
                }
            }
        },
        "api.HistoryListResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.HistoryItem"
                    }
                }
            }
        },
        "api.HistoryListResponseErr": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.PlayAutobetInput": {
            "type": "object",
            "properties": {
                "StopOnLoss": {
                    "type": "string"
                },
                "bet": {
                    "type": "string"
                },
                "multiplBashout": {
                    "type": "string"
                },
                "numOfAutobet": {
                    "type": "integer"
                },
                "playerToken": {
                    "type": "string"
                },
                "sessionUID": {
                    "type": "string"
                },
                "stopOnProfit": {
                    "type": "string"
                }
            }
        },
        "api.PlayAutobetResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                }
            }
        },
        "api.PlayAutobetResponseErr": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.PlayInput": {
            "type": "object",
            "properties": {
                "bet": {
                    "type": "string"
                },
                "commandID": {
                    "type": "string"
                },
                "multiplBashout": {
                    "type": "string"
                },
                "playerToken": {
                    "type": "string"
                },
                "sessionUID": {
                    "type": "string"
                }
            }
        },
        "api.PlayResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.PlayResponseErr": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.PlayeNameInput": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "isHidden": {
                    "type": "boolean"
                },
                "playerName": {
                    "type": "string"
                },
                "sessionUID": {
                    "type": "string"
                }
            }
        },
        "api.PlayeNameResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "playerToken": {
                    "type": "string"
                }
            }
        },
        "api.PlayeNameResponseErr": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "playerToken": {
                    "type": "string"
                }
            }
        },
        "api.PlayerItem": {
            "type": "object",
            "properties": {
                "bet": {
                    "type": "string"
                },
                "cashout": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "flg": {
                    "type": "string"
                },
                "isCashedout": {
                    "type": "boolean"
                },
                "playername": {
                    "type": "string"
                },
                "playertoken": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "api.PlayerListInput": {
            "type": "object",
            "properties": {
                "sessionUID": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "startOffset": {
                    "type": "integer"
                }
            }
        },
        "api.PlayerListResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.PlayerItem"
                    }
                }
            }
        },
        "api.PlayerListResponseErr": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "integer"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "api.TypeCrashArg_Annonce": {
            "type": "object",
            "properties": {
                "gameWillStart": {
                    "type": "number"
                }
            }
        },
        "api.TypeCrashArg_MsgWinOrLose": {
            "type": "object",
            "properties": {
                "isWin": {
                    "type": "string"
                },
                "multiplicator": {
                    "type": "string"
                },
                "timeline": {
                    "type": "number"
                },
                "walletAfter": {
                    "type": "string"
                }
            }
        },
        "api.TypeCrashArg_Playing": {
            "type": "object",
            "properties": {
                "multiplicator": {
                    "type": "string"
                },
                "timeline": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "InteractiveService",
	Description:      "This is the backend for the Interactive service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
