// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Arsham Roshannejad",
            "url": "arsham.cloudarshamdev2001@gmail.com",
            "email": "arshamdev2001@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://www.mit.edu/~amini/LICENSE.md"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "userRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_entities.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.LoginSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.UserNotFound"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.WrongPassword"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/auth/refresh-token": {
            "post": {
                "description": "Refresh token endpoint to get new access token with refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "description": "Token data",
                        "name": "tokenRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_entities.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RefreshSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "userRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_entities.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.UserCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/response.UserAlreadyExists"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/question": {
            "get": {
                "description": "Get all questions handled by teacher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Get All Questions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.questionResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new question",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Create Question",
                "parameters": [
                    {
                        "description": "Question data",
                        "name": "questionRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_entities.QuestionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.QuestionCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
                        "schema": {
                            "$ref": "#/definitions/response.QuestionAlreadyExists"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/question/:id": {
            "get": {
                "description": "Get single question with id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Get Question",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.QuestionData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.QuestionNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete question with id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questions"
                ],
                "summary": "Delete Question",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/template": {
            "get": {
                "description": "Get all templates",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Templates"
                ],
                "summary": "Get Templates",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_models.Templates"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new template for exam",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Templates"
                ],
                "summary": "Create Template",
                "parameters": [
                    {
                        "description": "Template data",
                        "name": "templateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_entities.TemplateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.TemplateCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
                        "schema": {
                            "$ref": "#/definitions/response.TemplateExists"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/template/{id}": {
            "delete": {
                "description": "Delete a template with id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Templates"
                ],
                "summary": "Delete Template",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Template ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get a user data and properties",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.UserNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a user based on jwt and current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "userRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_entities.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user based on jwt and current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete User",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        },
        "/user/change-password": {
            "post": {
                "description": "Change user password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Password Change",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "userRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_iarsham_teacher-tool-api_internal_entities.PassChangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.PasswordChanged"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.InternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_iarsham_teacher-tool-api_internal_entities.PassChangeRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8,
                    "example": "1qaz2wsx"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8,
                    "example": "1qaz2wsx"
                }
            }
        },
        "github_com_iarsham_teacher-tool-api_internal_entities.QuestionRequest": {
            "type": "object",
            "required": [
                "grade",
                "lesson",
                "level",
                "title"
            ],
            "properties": {
                "file": {
                    "$ref": "#/definitions/multipart.FileHeader"
                },
                "grade": {
                    "type": "integer",
                    "enum": [
                        0,
                        1,
                        2,
                        3,
                        4,
                        5
                    ],
                    "example": 5
                },
                "lesson": {
                    "type": "integer",
                    "example": 1
                },
                "level": {
                    "type": "integer",
                    "enum": [
                        0,
                        1,
                        2
                    ],
                    "example": 2
                },
                "title": {
                    "type": "string",
                    "example": "Math Stuff"
                }
            }
        },
        "github_com_iarsham_teacher-tool-api_internal_entities.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
                }
            }
        },
        "github_com_iarsham_teacher-tool-api_internal_entities.TemplateRequest": {
            "type": "object",
            "properties": {
                "file": {
                    "$ref": "#/definitions/multipart.FileHeader"
                }
            }
        },
        "github_com_iarsham_teacher-tool-api_internal_entities.UpdateUserRequest": {
            "type": "object",
            "required": [
                "phone"
            ],
            "properties": {
                "phone": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 10,
                    "example": "+9891154326250"
                },
                "role": {
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ],
                    "example": 1
                }
            }
        },
        "github_com_iarsham_teacher-tool-api_internal_entities.UserRequest": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8,
                    "example": "1qaz2wsx"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 10,
                    "example": "+9891154326250"
                }
            }
        },
        "github_com_iarsham_teacher-tool-api_internal_models.Templates": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-29T03:09:00+03:30"
                },
                "file": {
                    "type": "string",
                    "example": "domain.com/media/file.docx"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "multipart.FileHeader": {
            "type": "object",
            "properties": {
                "filename": {
                    "type": "string"
                },
                "header": {
                    "$ref": "#/definitions/textproto.MIMEHeader"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "response.BadRequest": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "bad request"
                }
            }
        },
        "response.InternalServerError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        },
        "response.LoginSuccess": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
                },
                "refreshToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
                }
            }
        },
        "response.PasswordChanged": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "password changed successfully"
                }
            }
        },
        "response.QuestionAlreadyExists": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "question already exists"
                }
            }
        },
        "response.QuestionCreated": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "question created"
                }
            }
        },
        "response.QuestionData": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2024-01-29T03:09:00+03:30"
                },
                "file": {
                    "type": "string",
                    "example": "domain.com/media/image.png"
                },
                "grade": {
                    "type": "string",
                    "example": "fifth"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lesson": {
                    "type": "string",
                    "example": "12"
                },
                "level": {
                    "type": "string",
                    "example": "advanced"
                },
                "title": {
                    "type": "string",
                    "example": "Math statistics"
                },
                "used": {
                    "type": "integer",
                    "example": 195
                },
                "userID": {
                    "type": "integer",
                    "example": 10
                },
                "views": {
                    "type": "integer",
                    "example": 666
                }
            }
        },
        "response.QuestionNotFound": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "question not found"
                }
            }
        },
        "response.RefreshSuccess": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
                }
            }
        },
        "response.TemplateCreated": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "template created"
                }
            }
        },
        "response.TemplateExists": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "template already exists"
                }
            }
        },
        "response.UserAlreadyExists": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "user already exists"
                }
            }
        },
        "response.UserCreated": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "string",
                    "example": "user created"
                }
            }
        },
        "response.UserData": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-29T03:09:00+03:30"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "phone": {
                    "type": "string",
                    "example": "+989029266610"
                },
                "role": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "response.UserNotFound": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "user not found"
                }
            }
        },
        "response.WrongPassword": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "wrong password"
                }
            }
        },
        "response.questionResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2024-01-29T03:09:00+03:30"
                },
                "file": {
                    "type": "string",
                    "example": "domain.com/media/image.png"
                },
                "grade": {
                    "type": "string",
                    "example": "fifth"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lesson": {
                    "type": "string",
                    "example": "12"
                },
                "level": {
                    "type": "string",
                    "example": "advanced"
                },
                "title": {
                    "type": "string",
                    "example": "Math statistics"
                },
                "used": {
                    "type": "integer",
                    "example": 195
                },
                "userID": {
                    "type": "integer",
                    "example": 10
                },
                "views": {
                    "type": "integer",
                    "example": 666
                }
            }
        },
        "textproto.MIMEHeader": {
            "type": "object",
            "additionalProperties": {
                "type": "array",
                "items": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Teacher-Tools-API",
	Description:      "API for Teacher Tools application that provides various endpoints for managing data.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
