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
        "/user": {
            "post": {
                "description": "Creates a new user in the database with a hashed password using the Argon2id algorithm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.CreateUserParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user_id: ID of the created user",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Authenticates a user by comparing the submitted password to the hashed password stored in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.CreateUserParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Logged in user information",
                        "schema": {
                            "$ref": "#/definitions/db.CreateUserParams"
                        }
                    },
                    "400": {
                        "description": "Invalid login credentials",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/password": {
            "post": {
                "description": "Hashes and updates the password for an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update password",
                "parameters": [
                    {
                        "description": "Password Update Request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.ChangePasswordParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password updated successfully"
                    },
                    "400": {
                        "description": "Invalid input or password criteria not met",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error during password update",
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
        "/users/{id}": {
            "delete": {
                "description": "Deletes a user by ID from the database, requires admin privileges",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully deleted"
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.ChangePasswordParams": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "password_hash": {
                    "type": "string"
                }
            }
        },
        "db.CreateUserParams": {
            "type": "object",
            "properties": {
                "password_hash": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
