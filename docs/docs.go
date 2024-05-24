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
        "/v1/item": {
            "get": {
                "description": "list items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "List item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "회사 이름 검색 키워드",
                        "name": "company",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                0,
                                1,
                                2,
                                3,
                                4
                            ],
                            "type": "integer"
                        },
                        "collectionFormat": "multi",
                        "description": "회사 규모 (0:스타트업, 1:중소기업, 2:중견기업, 3:대기업, 4:외국계)",
                        "name": "company_size",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "multi",
                        "description": "관련 직무 DB ID 배열",
                        "name": "job_tags",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "multi",
                        "description": "관련 스킬 DB ID 배열",
                        "name": "skill_tags",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "페이지",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_internal_controller_rest_dto_item_dto.FindAllItemRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/tag/job": {
            "get": {
                "description": "list job tags",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "List Job Tag",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_team-nerd-planet_api-server_internal_controller_rest_dto_tag_dto.FindAllJobTagRes"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/tag/skill": {
            "get": {
                "description": "list skill tags",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "List Skill Tag",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_team-nerd-planet_api-server_internal_controller_rest_dto_tag_dto.FindAllSkillTagRes"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_infra_router_util.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_team-nerd-planet_api-server_infra_router_util.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_team-nerd-planet_api-server_internal_controller_rest_dto_item_dto.FindAllItemRes": {
            "type": "object",
            "properties": {
                "company_size": {
                    "description": "회사 규모",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_team-nerd-planet_api-server_internal_entity.CompanySizeType"
                        }
                    ]
                },
                "feed_link": {
                    "description": "회사 URL",
                    "type": "string"
                },
                "feed_name": {
                    "description": "회사 이름",
                    "type": "string"
                },
                "feed_title": {
                    "description": "회사 Feed 제목",
                    "type": "string"
                },
                "item_description": {
                    "description": "글 설명",
                    "type": "string"
                },
                "item_id": {
                    "description": "글 DB ID",
                    "type": "integer"
                },
                "item_link": {
                    "description": "글 URL",
                    "type": "string"
                },
                "item_published": {
                    "description": "개시 시간",
                    "type": "string"
                },
                "item_title": {
                    "description": "글 제목",
                    "type": "string"
                },
                "job_tags_id_arr": {
                    "description": "관련 직무 DB ID 배열",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "skill_tags_id_arr": {
                    "description": "관련 스킬 DB ID 배열",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "github_com_team-nerd-planet_api-server_internal_controller_rest_dto_tag_dto.FindAllJobTagRes": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "직무 DB ID",
                    "type": "integer"
                },
                "name": {
                    "description": "직무 이름",
                    "type": "string"
                }
            }
        },
        "github_com_team-nerd-planet_api-server_internal_controller_rest_dto_tag_dto.FindAllSkillTagRes": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "스킬 DB ID",
                    "type": "integer"
                },
                "name": {
                    "description": "스킬 이름",
                    "type": "string"
                }
            }
        },
        "github_com_team-nerd-planet_api-server_internal_entity.CompanySizeType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-comments": {
                "FOREIGN": "외국계",
                "LARGE": "대기업",
                "MEDIUM": "중견기업",
                "SMALL": "중소기업",
                "STARTUP": "스타트업"
            },
            "x-enum-varnames": [
                "STARTUP",
                "SMALL",
                "MEDIUM",
                "LARGE",
                "FOREIGN"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "nerd-planet-api-server",
	Description:      "nerd-planet-api-server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
