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
        "/api/v1/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Authorization",
                "parameters": [
                    {
                        "description": "Auth data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.AuthDataLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.JwtPair"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Update Access Token",
                "parameters": [
                    {
                        "description": "Data Refresh Token",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.AuthDataRefresh"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.JwtPair"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Registration",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.AuthDataRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.JwtPair"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Data Of My Account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AuthDataLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.AuthDataRefresh": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "auth.AuthDataRegister": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.JwtPair": {
            "type": "object",
            "properties": {
                "accessExpiresIn": {
                    "type": "string"
                },
                "accessToken": {
                    "type": "string"
                },
                "refreshExpiresIn": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "lastSeen": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/enum.Role"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "enum.Role": {
            "type": "string",
            "enum": [
                "USER",
                "ADMIN"
            ],
            "x-enum-varnames": [
                "User",
                "Admin"
            ]
        },
        "helper.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "serverCode": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Bearer Token authortization",
            "type": "apiKey",
            "name": "Authorization",
            "in": "Header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "<SHIKARU TEMPLATE>",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
