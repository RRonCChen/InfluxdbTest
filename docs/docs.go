// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/influx/v1/hello": {
            "get": {
                "tags": [
                    "Hello"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/influx/v1/temperature": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Temperature"
                ],
                "parameters": [
                    {
                        "description": "location deviceId temperature",
                        "name": "temperatureRecord",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TemperatureRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/influx/v1/temperature/avg/{location}/{deviceId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Temperature"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "deviceId",
                        "name": "deviceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "location",
                        "name": "location",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.TemperatureRecord": {
            "type": "object",
            "properties": {
                "deviceId": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "temperature": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Ron influxdb swagger demo",
	Description:      "influxdb practice",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
