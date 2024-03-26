package backend

import (
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathImport(b *backend) *framework.Path {
	return &framework.Path{
		Pattern:      "import/memonic",
		HelpSynopsis: "Import an Ethereum account using memonic",
		HelpDescription: `

    GET - return the account by the name with the private key

    `,
		Fields: map[string]*framework.FieldSchema{
			"name": &framework.FieldSchema{Type: framework.TypeString},
			"memonic": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "string consisting of the memonic words",
				Default:     "",
			},
			"password": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "pharse phase",
				Default:     "",
			},
		},
		ExistenceCheck: b.pathExistenceCheck,
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: b.memonic,
		},
	}
}
