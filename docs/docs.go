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
        "/car": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "GetCars with filter and pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string",
                        "name": "regNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string",
                        "name": "mark",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string",
                        "name": "model",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "2002",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string",
                        "name": "patronymic",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Car"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "UpdateCars By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "1",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "for update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "CreateCar By RegNum",
                "parameters": [
                    {
                        "description": "Creating",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "DeleteCars By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "1",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CreateRequest": {
            "type": "object",
            "properties": {
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.Car": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.Owner"
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.Owner": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "goTest",
	Description:      "API Server for goTest Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
