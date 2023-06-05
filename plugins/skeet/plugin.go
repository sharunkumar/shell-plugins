package skeet

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "skeet",
		Platform: schema.PlatformInfo{
			Name:     "Bluesky",
			Homepage: sdk.URL("https://bsky.app"),
		},
		Credentials: []schema.CredentialType{
			AppPassword(),
		},
		Executables: []schema.Executable{
			BlueskyCLI(),
		},
	}
}
