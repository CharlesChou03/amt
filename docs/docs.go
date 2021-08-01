// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/tutor/{tutor}": {
            "get": {
                "description": "get tutor information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tutor Information"
                ],
                "summary": "get tutor information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "tutor",
                        "name": "tutor",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/models.GetTutorRes"
                        }
                    },
                    "204": {
                        "description": "no content"
                    },
                    "400": {
                        "description": "bad request"
                    }
                }
            }
        },
        "/health": {
            "get": {
                "summary": "health checker API",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "summary": "service version API",
                "responses": {
                    "200": {
                        "description": "0.0.1",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.GetTutorRes": {
            "type": "object",
            "properties": {
                "headline": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price_info": {
                    "$ref": "#/definitions/models.PriceInfo"
                },
                "slug": {
                    "type": "string"
                },
                "teaching_languages": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "models.PriceInfo": {
            "type": "object",
            "properties": {
                "normal": {
                    "type": "number"
                },
                "trial": {
                    "type": "number"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
