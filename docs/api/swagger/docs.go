// Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/me": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "Get current user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.CreateUserParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/short-uri": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uri"
                ],
                "summary": "Update short uri",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UpdateURIParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.URL"
                        }
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
                    "uri"
                ],
                "summary": "Create new short uri",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.CreateURIParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.URL"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Create new token",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.LoginParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.LoginResponse"
                        }
                    }
                }
            }
        },
        "/urls": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "Get current user urls",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.URL"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.CreateURIParams": {
            "type": "object",
            "required": [
                "original_url"
            ],
            "properties": {
                "original_url": {
                    "type": "string",
                    "minLength": 5
                },
                "short_uri": {
                    "type": "string"
                }
            }
        },
        "entity.CreateUserParams": {
            "type": "object",
            "required": [
                "full_name",
                "password",
                "username"
            ],
            "properties": {
                "full_name": {
                    "type": "string",
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "entity.LoginParams": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "entity.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_expires_at": {
                    "type": "string"
                }
            }
        },
        "entity.URL": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "original_url": {
                    "type": "string"
                },
                "requested_count": {
                    "type": "integer"
                },
                "short_uri": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "entity.UpdateURIParams": {
            "type": "object",
            "required": [
                "new_short_uri",
                "old_short_uri"
            ],
            "properties": {
                "new_short_uri": {
                    "type": "string"
                },
                "old_short_uri": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
