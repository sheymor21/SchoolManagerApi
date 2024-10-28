// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "Jose Armando Coronel Vasquez",
            "email": "joseacvz81@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Get JWT by Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Get JWT by Login",
                "parameters": [
                    {
                        "description": "Login User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mark": {
            "get": {
                "description": "Fetches a mark object based on the provided student ID",
                "tags": [
                    "marks"
                ],
                "summary": "Retrieve a mark",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved mark",
                        "schema": {
                            "$ref": "#/definitions/MarksGetRequest"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new mark entry with student and teacher details",
                "tags": [
                    "mark"
                ],
                "summary": "Add a new mark",
                "parameters": [
                    {
                        "description": "Mark Add Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/MarkAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully added mark",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a mark entry from the database using the provided mark ID",
                "tags": [
                    "mark"
                ],
                "summary": "Delete a mark",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mark ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted mark",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/marks": {
            "get": {
                "description": "Finds and returns the marks of a student using their carnet number.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "marks"
                ],
                "summary": "Retrieves student's marks by carnet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student Carnet",
                        "name": "Carnet",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/MarksGetRequest"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Modifies an existing mark entry using the supplied ID and mark details",
                "tags": [
                    "mark"
                ],
                "summary": "Update a mark",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Mark ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Marks Update Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/MarksUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated mark",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/student": {
            "get": {
                "description": "Retrieve a student's information from the database using their carnet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "Get student by carnet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "StudentRol Carnet",
                        "name": "Carnet",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    },
                    "404": {
                        "description": "StudentRol not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing student's information in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "Update a student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "StudentRol Carnet",
                        "name": "Carnet",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Update StudentRol",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Student"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Add a new student to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "Add a student",
                "parameters": [
                    {
                        "description": "Add StudentRol",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/StudentAddRequest"
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
                "description": "Delete a student from the database using their carnet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "student"
                ],
                "summary": "Delete a student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "StudentRol Carnet",
                        "name": "Carnet",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "StudentRol deleted successfully"
                    },
                    "404": {
                        "description": "StudentRol not found",
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
        "/students": {
            "get": {
                "description": "Fetches a list of all students from the database and returns the data as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Retrieves students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Student"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/teacher": {
            "get": {
                "description": "Fetches the information of a teacher from the database using the \"Carnet\" query parameter.",
                "tags": [
                    "teacher"
                ],
                "summary": "Get a teacher's details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Teacher Carnet",
                        "name": "Carnet",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Teacher"
                        }
                    },
                    "404": {
                        "description": "Not Found",
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
                "description": "Updates the information of an existing teacher in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teacher"
                ],
                "summary": "Update an existing teacher",
                "parameters": [
                    {
                        "description": "Updated Teacher",
                        "name": "teacher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Teacher"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Teacher"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Inserts a new teacher record to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teacher"
                ],
                "summary": "Add a new teacher",
                "parameters": [
                    {
                        "description": "New Teacher",
                        "name": "teacher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/TeacherAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing teacher record from the database using the \"Carnet\" query parameter.",
                "tags": [
                    "teacher"
                ],
                "summary": "Delete a teacher",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Teacher Carnet",
                        "name": "Carnet",
                        "in": "query",
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
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/teachers": {
            "get": {
                "description": "Fetch all teacher records from the database and return them as a JSON payload",
                "tags": [
                    "teachers"
                ],
                "summary": "Retrieve all teachers",
                "responses": {
                    "200": {
                        "description": "List of teachers",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/TeacherGetRequest"
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
        "MarkAddRequest": {
            "type": "object",
            "required": [
                "grade",
                "mark",
                "semester",
                "student_carnet",
                "teacher_carnet"
            ],
            "properties": {
                "grade": {
                    "type": "string"
                },
                "mark": {
                    "type": "string"
                },
                "semester": {
                    "type": "string"
                },
                "student_carnet": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 10
                },
                "teacher_carnet": {
                    "type": "string"
                }
            }
        },
        "MarksGetRequest": {
            "type": "object",
            "properties": {
                "grade": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "mark": {
                    "type": "string"
                },
                "semester": {
                    "type": "string"
                },
                "student_carnet": {
                    "type": "string"
                },
                "teacher_carnet": {
                    "type": "string"
                }
            }
        },
        "MarksUpdateRequest": {
            "type": "object",
            "properties": {
                "grade": {
                    "type": "string"
                },
                "mark": {
                    "type": "string"
                },
                "semester": {
                    "type": "string"
                },
                "studentCarnet": {
                    "type": "string"
                },
                "teacherCarnet": {
                    "type": "string"
                }
            }
        },
        "StudentAddRequest": {
            "type": "object",
            "required": [
                "age",
                "carnet",
                "classroom",
                "firstName",
                "lastName"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "carnet": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 10
                },
                "classroom": {
                    "type": "string",
                    "maxLength": 4,
                    "minLength": 2
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "TeacherAddRequest": {
            "type": "object",
            "required": [
                "age",
                "carnet",
                "classroom",
                "firstName",
                "lastName"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "carnet": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 10
                },
                "classroom": {
                    "type": "string",
                    "maxLength": 4,
                    "minLength": 2
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "TeacherGetRequest": {
            "type": "object",
            "required": [
                "age",
                "carnet",
                "classroom",
                "firstName",
                "lastName"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "carnet": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 10
                },
                "classroom": {
                    "type": "string",
                    "maxLength": 4,
                    "minLength": 2
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "UserLoginRequest": {
            "type": "object",
            "required": [
                "carnet",
                "password"
            ],
            "properties": {
                "carnet": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Student": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "carnet": {
                    "type": "string"
                },
                "classroom": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "models.Teacher": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "carnet": {
                    "type": "string"
                },
                "classroom": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "hi , how are you?",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Marks Api",
	Description:      "This is an API for managing marks",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
