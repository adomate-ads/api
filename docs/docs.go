// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Adomate API Support",
            "url": "https://adomate.com/support",
            "email": "support@adomate.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageResponse"
                        }
                    }
                }
            }
        },
        "/billing": {
            "get": {
                "description": "Gets a slice of all bills.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Get all Bills",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Billing"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new bill.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Create Bill",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/billing/:id": {
            "get": {
                "description": "Gets all information about a single bill.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Gets a Bill",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Billing"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a bill.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Delete Bill",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update information about a bill.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Update Bill",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/models.Billing"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/billing/company/:id": {
            "get": {
                "description": "Gets a slice of all the bills for a specific company.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Billing"
                ],
                "summary": "Get all Bills for a Company",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Billing"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ErrorResponse": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.MessageResponse": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Billing": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 900.25
                },
                "comments": {
                    "type": "string",
                    "example": "Something about the invoice..."
                },
                "company": {
                    "$ref": "#/definitions/models.Company"
                },
                "created_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "due_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "issued_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "name": {
                    "type": "integer"
                },
                "status": {
                    "description": "Available options: paid, unpaid, pending",
                    "type": "string",
                    "example": "paid"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                }
            }
        },
        "models.Company": {
            "type": "object",
            "properties": {
                "budget": {
                    "type": "integer",
                    "example": 1000
                },
                "created_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "domain": {
                    "type": "string",
                    "example": "raajpatel.dev"
                },
                "email": {
                    "type": "string",
                    "example": "the@raajpatel.dev"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "industry": {
                    "$ref": "#/definitions/models.Industry"
                },
                "name": {
                    "type": "string",
                    "example": "Google LLC"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                }
            }
        },
        "models.Industry": {
            "type": "object",
            "properties": {
                "Industry": {
                    "type": "string",
                    "example": "Health Care"
                },
                "created_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "updated_at": {
                    "type": "string",
                    "example": "2020-01-01T00:00:00Z"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Adomate API",
	Description:      "Adomate Monolithic API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
