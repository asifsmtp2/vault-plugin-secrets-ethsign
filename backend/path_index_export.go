package backend

import (
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathExportIndex(b *backend) *framework.Path {
	return &framework.Path{
		Pattern:      "export/accounts/" + framework.GenericNameWithAtRegex("index") + "/" + framework.GenericNameRegex("name"),
		HelpSynopsis: "Export an Ethereum account",
		HelpDescription: `

    POST - return the account by the name with the private key

    `,
		Fields: map[string]*framework.FieldSchema{
			"index": &framework.FieldSchema{Type: framework.TypeInt},
			"name":  &framework.FieldSchema{Type: framework.TypeString},
			"hdpath": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "DerivationPath represents the computer friendly version of a hierarchical",
				Default:     "m/44'/60'/0'/0/",
			},
		},
		ExistenceCheck: b.pathExistenceCheck,
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: b.exportIndexAccount,
		},
	}
}
