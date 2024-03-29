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
        "/category": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get categories",
                "operationId": "get-categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Category"
                            }
                        }
                    }
                }
            }
        },
        "/certificate": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Get all bought certificates",
                "operationId": "get-bought-certificate",
                "parameters": [
                    {
                        "description": "input info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Certificate"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Create new certificate",
                "operationId": "create-new-certificate",
                "parameters": [
                    {
                        "description": "certificate info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Certificate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/certificate/link": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Generate link",
                "operationId": "generate-link",
                "parameters": [
                    {
                        "description": "certificate info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Certificate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Accept link",
                "operationId": "accept-link",
                "parameters": [
                    {
                        "description": "input info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Input"
                        }
                    }
                }
            }
        },
        "/certificate/phone-number": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "certificate"
                ],
                "summary": "Share phone number",
                "operationId": "share-number",
                "parameters": [
                    {
                        "description": "input info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Input"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Input"
                        }
                    }
                }
            }
        },
        "/service": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Get services",
                "operationId": "get-services",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Service"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sum": {
                    "type": "string"
                }
            }
        },
        "entity.Certificate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "parameters": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "entity.Service": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sub_services": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "handler.Input": {
            "type": "object",
            "properties": {
                "cert_id": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "droplet.senkevichdev.work:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Backend for MTS Hackathon",
	Description:      "API Server for mobile application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
