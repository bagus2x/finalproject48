// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://github.com/bagus2x",
            "email": "tubagus.sflh@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "list of all categories.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "get all categories.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CategoryResponse"
                            }
                        }
                    }
                }
            }
        },
        "/category": {
            "post": {
                "description": "Creating a new category.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "create a new category.",
                "parameters": [
                    {
                        "description": "the body for creating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddCategoryRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.CategoryResponse"
                        }
                    }
                }
            }
        },
        "/category/{categoryId}": {
            "get": {
                "description": "finding a category by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "get a category by id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category id",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CategoryResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "updating a category.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "update single category.",
                "parameters": [
                    {
                        "description": "the body for updating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateCategoryRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Category id",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CategoryResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "deleting a new category.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "delete a category.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Category id",
                        "name": "categoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/comment": {
            "post": {
                "description": "Creating a new comment.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "create a new comment.",
                "parameters": [
                    {
                        "description": "the body for creating comment",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddCommentRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comment/{commentId}": {
            "get": {
                "description": "finding a comment by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "get a comment by id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment id",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "put": {
                "description": "updating a comment.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "update single comment.",
                "parameters": [
                    {
                        "description": "the body for updating comment",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateCommentRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Comment id",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "delete": {
                "description": "deleting a comment by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "delete single comment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Comment id",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "Creating a new post.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "create a new post.",
                "parameters": [
                    {
                        "description": "the body for creating post",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreatePostRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.PostResponse"
                        }
                    }
                }
            }
        },
        "/post/{postId}": {
            "get": {
                "description": "finding a post by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "get a category by id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post id",
                        "name": "postId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PostResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "updating a post.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "update single post.",
                "parameters": [
                    {
                        "description": "the body for updating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdatePostRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Post id",
                        "name": "postId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PostResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "deleting a post.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "delete single post.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Post id",
                        "name": "postId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/post/{postId}/comments": {
            "get": {
                "description": "finding  comment by post id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "get comments by post id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post id",
                        "name": "postId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Comment"
                            }
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "description": "list of all posts.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "get all posts.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PostResponse"
                            }
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "Update authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update a user.",
                "parameters": [
                    {
                        "description": "the body for updating user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SignUpResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete a user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/signin": {
            "post": {
                "description": "Gaining access to the protected API. Put the token from response in the Authorization header to access protected API.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "log in with registered email.",
                "parameters": [
                    {
                        "description": "the body to sign in",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SignInResponse"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Creating a new account to authenticate a user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create a new account.",
                "parameters": [
                    {
                        "description": "the body for sign up",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SignUpResponse"
                        }
                    }
                }
            }
        },
        "/user/{userId}": {
            "get": {
                "description": "Finding user by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get a user by id.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "list of all users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get all users.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddCategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "models.AddCommentRequest": {
            "type": "object",
            "required": [
                "body",
                "postId"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "postId": {
                    "type": "integer"
                }
            }
        },
        "models.CategoryResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Comment": {
            "type": "object",
            "properties": {
                "authorId": {
                    "type": "integer"
                },
                "body": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "postId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.CreatePostRequest": {
            "type": "object",
            "required": [
                "body",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "maxLength": 8142
                },
                "categoryIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "thumbnail": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "models.PostResponse": {
            "type": "object",
            "required": [
                "body",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "maxLength": 8142
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CategoryResponse"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "thumbnail": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 256
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.SignInRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 5
                }
            }
        },
        "models.SignInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 5
                }
            }
        },
        "models.SignUpResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.UpdateCategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "models.UpdateCommentRequest": {
            "type": "object",
            "required": [
                "body"
            ],
            "properties": {
                "body": {
                    "type": "string"
                }
            }
        },
        "models.UpdatePostRequest": {
            "type": "object",
            "required": [
                "body",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "maxLength": 8142
                },
                "categoryIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "thumbnail": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "models.UpdateUserRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "updatedAt": {
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
