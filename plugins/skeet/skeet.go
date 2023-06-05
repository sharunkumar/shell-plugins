package skeet

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func BlueskyCLI() schema.Executable {
	return schema.Executable{
		Name:    "Skeet",
		Runs:    []string{"skeet"},
		DocsURL: sdk.URL("https://github.com/sharunkumar/skeet"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.AppPassword,
			},
		},
	}
}
