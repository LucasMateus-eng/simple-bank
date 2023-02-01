// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Supports",
            "email": "lucas.falecomigo@gmail.com"
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
        "/health": {
            "get": {
                "description": "Get API availability - if it's running",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Get API availability",
                "responses": {
                    "200": {
                        "description": "API is available.",
                        "schema": {
                            "$ref": "#/definitions/router.AppStatus"
                        }
                    }
                }
            }
        },
        "/wallet": {
            "put": {
                "description": "Update a wallet in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Update a wallet",
                "parameters": [
                    {
                        "description": "Wallet DTO for update",
                        "name": "wallet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aggregate.WalletForUpdateAPI"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Wallet successfully updated.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseOKWithData"
                        }
                    },
                    "400": {
                        "description": "Invalid payload.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    },
                    "500": {
                        "description": "Failed to update wallet.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new wallet in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Create a wallet",
                "parameters": [
                    {
                        "description": "Wallet DTO for create",
                        "name": "wallet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aggregate.WalletAPI"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Wallet successfully created.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseOKWithData"
                        }
                    },
                    "400": {
                        "description": "Invalid payload.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    },
                    "500": {
                        "description": "Failed to create wallet.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    }
                }
            }
        },
        "/wallet/transfer": {
            "put": {
                "description": "Transfer money between two wallets in the database (transaction)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Transfer money",
                "parameters": [
                    {
                        "description": "Wallet DTO for transfer",
                        "name": "wallet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/valueobject.TransferAPI"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Transfer performed successfully.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseOKWithData"
                        }
                    },
                    "400": {
                        "description": "Invalid payload.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    },
                    "500": {
                        "description": "Failed to perform transfer.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    }
                }
            }
        },
        "/wallet/{wallet_id}": {
            "get": {
                "description": "Get a wallet based on the wallet id from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet identifier",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful get wallet.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseOKWithData"
                        }
                    },
                    "400": {
                        "description": "Error parsing the id parameter.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    },
                    "404": {
                        "description": "Error: wallet can't find.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    },
                    "500": {
                        "description": "Error getting wallet.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a wallet based on the wallet id from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Delete a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "wallet identifier",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Wallet successfully deleted.",
                        "schema": {
                            "$ref": "#/definitions/formatter.Response"
                        }
                    },
                    "400": {
                        "description": "Error parsing the id parameter.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    },
                    "500": {
                        "description": "Failed to delete wallet.",
                        "schema": {
                            "$ref": "#/definitions/formatter.ResponseErrorWithData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aggregate.WalletAPI": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/entity.AccountAPI"
                },
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.EntryAPI"
                    }
                },
                "person": {
                    "$ref": "#/definitions/entity.PersonAPI"
                },
                "transfers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/valueobject.TransferAPI"
                    }
                }
            }
        },
        "aggregate.WalletForUpdateAPI": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/entity.AccountAPI"
                },
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.EntryAPI"
                    }
                },
                "person": {
                    "$ref": "#/definitions/entity.PersonAPI"
                },
                "transfers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/valueobject.TransferAPI"
                    }
                }
            }
        },
        "entity.AccountAPI": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "entity.EntryAPI": {
            "type": "object",
            "properties": {
                "account_uuid": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "entity.PersonAPI": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_a_shopkeeper": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "personal_id": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "formatter.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "formatter.ResponseErrorWithData": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "formatter.ResponseOKWithData": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "router.AppStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "UP"
                }
            }
        },
        "valueobject.TransferAPI": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "from_wallet_uuid": {
                    "type": "string"
                },
                "to_wallet_uuid": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Simple-bank API",
	Description:      "This is a simple-bank management application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
