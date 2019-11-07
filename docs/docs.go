// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-10-17 17:47:57.887651112 +0900 KST m=+0.063075053

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "SYM-GW 애플리케이션의 상태를 체크합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "SYM-GW 애플리케이션의 상태를 체크합니다."
            }
        },
        "/v1/bootnode/closeNodes": {
            "get": {
                "description": "BootNode API 요청하여 Node 정보를 제공합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BootNode API 요청하여 가장 근접한 Node 정보를 제공합니다.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonRpcResponse"
                        }
                    }
                }
            }
        },
        "/v1/bootnode/nodes": {
            "get": {
                "description": "BootNode API 요청하여 Node 정보를 제공합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BootNode API 요청하여 Node 정보를 제공합니다.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonRpcResponse"
                        }
                    }
                }
            }
        },
        "/v1/rpc/node/{number}": {
            "post": {
                "description": "WorkNode로 HTTP JSON-RPC 요청을 Proxy 합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "WorkNode로 WebSocket JSON-RPC 요청을 Proxy 합니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node Number",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Json-Rpc Request",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.JsonRpcRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonRpcResponse"
                        }
                    }
                }
            }
        },
        "/v1/rpc/node/{number}/ws": {
            "post": {
                "description": "WorkNode로 WebSocket JSON-RPC 요청을 Proxy 합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "WorkNode로 WebSocket JSON-RPC 요청을 Proxy 합니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Node Number",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Json-Rpc Request",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.JsonRpcRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.JsonRpcResponse"
                        }
                    }
                }
            }
        },
        "/v1/rpc/nodes": {
            "get": {
                "description": "Proxy를 제공할 URL 정보가 있는 WorkNode 리스트를 요청합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Static WorkNodes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/config.WorkNode"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.JsonRpcRequest": {
            "type": "object",
            "required": [
                "id",
                "jsonrpc",
                "method",
                "params"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "jsonrpc": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "params": {
                    "type": "array",
                    "items": {
                        "type": "\u0026{%!s(token.Pos=466) %!s(*ast.FieldList=\u0026{475 [] 476}) %!s(bool=false)}"
                    }
                }
            }
        },
        "common.JsonRpcResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object"
                },
                "id": {
                    "type": "integer"
                },
                "jsonrpc": {
                    "type": "string"
                },
                "result": {
                    "type": "object"
                }
            }
        },
        "config.WorkNode": {
            "type": "object",
            "properties": {
                "httpUrl": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "symId": {
                    "type": "string"
                },
                "wsUrl": {
                    "type": "string"
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
	Version:     "1.0.0",
	Host:        "testnet-gateway.symverse.com",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "SYM-GW API Docs",
	Description: "Gsym API 문서입니다.",
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
