// Code generated by go-swagger; DO NOT EDIT.

package handler

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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API to enrich and manage information about individuals",
    "title": "Enrichment Service API",
    "version": "1.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/people": {
      "get": {
        "summary": "Get people with filters and pagination",
        "parameters": [
          {
            "type": "integer",
            "default": 1,
            "description": "Page number for pagination",
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 10,
            "description": "Number of items per page",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "List of people",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Person"
              }
            }
          }
        }
      },
      "post": {
        "summary": "Add a new person",
        "parameters": [
          {
            "description": "Person information",
            "name": "person",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PersonInput"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Person added successfully",
            "schema": {
              "$ref": "#/definitions/Person"
            }
          }
        }
      }
    },
    "/people/{id}": {
      "get": {
        "summary": "Get person by ID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the person to retrieve",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Person information",
            "schema": {
              "$ref": "#/definitions/Person"
            }
          }
        }
      },
      "put": {
        "summary": "Update person by ID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the person to update",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated person information",
            "name": "person",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PersonInput"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Person updated successfully",
            "schema": {
              "$ref": "#/definitions/Person"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete person by ID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the person to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Person deleted successfully"
          }
        }
      }
    }
  },
  "definitions": {
    "Person": {
      "type": "object",
      "properties": {
        "age": {
          "type": "integer"
        },
        "gender": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "nationality": {
          "type": "string"
        },
        "patronymic": {
          "type": "string"
        },
        "surname": {
          "type": "string"
        }
      }
    },
    "PersonInput": {
      "type": "object",
      "properties": {
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
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API to enrich and manage information about individuals",
    "title": "Enrichment Service API",
    "version": "1.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/people": {
      "get": {
        "summary": "Get people with filters and pagination",
        "parameters": [
          {
            "type": "integer",
            "default": 1,
            "description": "Page number for pagination",
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 10,
            "description": "Number of items per page",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "List of people",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Person"
              }
            }
          }
        }
      },
      "post": {
        "summary": "Add a new person",
        "parameters": [
          {
            "description": "Person information",
            "name": "person",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PersonInput"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Person added successfully",
            "schema": {
              "$ref": "#/definitions/Person"
            }
          }
        }
      }
    },
    "/people/{id}": {
      "get": {
        "summary": "Get person by ID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the person to retrieve",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Person information",
            "schema": {
              "$ref": "#/definitions/Person"
            }
          }
        }
      },
      "put": {
        "summary": "Update person by ID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the person to update",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated person information",
            "name": "person",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PersonInput"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Person updated successfully",
            "schema": {
              "$ref": "#/definitions/Person"
            }
          }
        }
      },
      "delete": {
        "summary": "Delete person by ID",
        "parameters": [
          {
            "type": "integer",
            "description": "ID of the person to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Person deleted successfully"
          }
        }
      }
    }
  },
  "definitions": {
    "Person": {
      "type": "object",
      "properties": {
        "age": {
          "type": "integer"
        },
        "gender": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "nationality": {
          "type": "string"
        },
        "patronymic": {
          "type": "string"
        },
        "surname": {
          "type": "string"
        }
      }
    },
    "PersonInput": {
      "type": "object",
      "properties": {
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
}`))
}
