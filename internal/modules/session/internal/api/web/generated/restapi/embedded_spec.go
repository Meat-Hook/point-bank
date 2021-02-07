// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Microservice for managing user session.",
    "title": "Session service.",
    "version": "0.1.0"
  },
  "basePath": "/session/api/v1",
  "paths": {
    "/login": {
      "post": {
        "security": [],
        "description": "Login for user.",
        "operationId": "login",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginParam"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            },
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "Session auth."
              }
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/logout": {
      "post": {
        "description": "Logout for user.",
        "operationId": "logout",
        "responses": {
          "204": {
            "$ref": "#/responses/NoContent"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    }
  },
  "definitions": {
    "Email": {
      "type": "string",
      "format": "email",
      "maxLength": 255,
      "minLength": 1
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "x-order": 0
        }
      }
    },
    "LoginParam": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "x-order": 0,
          "$ref": "#/definitions/Email"
        },
        "password": {
          "x-order": 1,
          "$ref": "#/definitions/Password"
        }
      }
    },
    "Password": {
      "type": "string",
      "format": "password",
      "maxLength": 100,
      "minLength": 8
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "username",
        "email"
      ],
      "properties": {
        "email": {
          "x-order": 2,
          "$ref": "#/definitions/Email"
        },
        "id": {
          "x-order": 0,
          "$ref": "#/definitions/UserID"
        },
        "username": {
          "x-order": 1,
          "$ref": "#/definitions/Username"
        }
      }
    },
    "UserID": {
      "type": "integer",
      "format": "int32"
    },
    "Username": {
      "type": "string",
      "maxLength": 30,
      "minLength": 1
    }
  },
  "responses": {
    "GenericError": {
      "description": "Generic error response.",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NoContent": {
      "description": "The server successfully processed the request and is not returning any content."
    }
  },
  "securityDefinitions": {
    "cookieKey": {
      "description": "Session auth inside cookie.",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookieKey": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Microservice for managing user session.",
    "title": "Session service.",
    "version": "0.1.0"
  },
  "basePath": "/session/api/v1",
  "paths": {
    "/login": {
      "post": {
        "security": [],
        "description": "Login for user.",
        "operationId": "login",
        "parameters": [
          {
            "name": "args",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginParam"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            },
            "headers": {
              "Set-Cookie": {
                "type": "string",
                "description": "Session auth."
              }
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/logout": {
      "post": {
        "description": "Logout for user.",
        "operationId": "logout",
        "responses": {
          "204": {
            "description": "The server successfully processed the request and is not returning any content."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Email": {
      "type": "string",
      "format": "email",
      "maxLength": 255,
      "minLength": 1
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "x-order": 0
        }
      }
    },
    "LoginParam": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "x-order": 0,
          "$ref": "#/definitions/Email"
        },
        "password": {
          "x-order": 1,
          "$ref": "#/definitions/Password"
        }
      }
    },
    "Password": {
      "type": "string",
      "format": "password",
      "maxLength": 100,
      "minLength": 8
    },
    "User": {
      "type": "object",
      "required": [
        "id",
        "username",
        "email"
      ],
      "properties": {
        "email": {
          "x-order": 2,
          "$ref": "#/definitions/Email"
        },
        "id": {
          "x-order": 0,
          "$ref": "#/definitions/UserID"
        },
        "username": {
          "x-order": 1,
          "$ref": "#/definitions/Username"
        }
      }
    },
    "UserID": {
      "type": "integer",
      "format": "int32"
    },
    "Username": {
      "type": "string",
      "maxLength": 30,
      "minLength": 1
    }
  },
  "responses": {
    "GenericError": {
      "description": "Generic error response.",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NoContent": {
      "description": "The server successfully processed the request and is not returning any content."
    }
  },
  "securityDefinitions": {
    "cookieKey": {
      "description": "Session auth inside cookie.",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookieKey": []
    }
  ]
}`))
}