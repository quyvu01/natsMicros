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
        "/api/masterData/getProvinces": {
            "get": {
                "description": "Get list of provinces",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "masterData"
                ],
                "summary": "List Provinces",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Searching for province name",
                        "name": "SearchKey",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Sizing",
                        "name": "PageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Index",
                        "name": "PageIndex",
                        "in": "query"
                    }
                ],
                "responses": {}
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
