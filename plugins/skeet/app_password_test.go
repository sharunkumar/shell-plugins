package skeet

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestAppPasswordProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, AppPassword().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.: "kazhj4oy2wjfexample",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"SKEET": "kazhj4oy2wjfexample",
				},
			},
		},
	})
}

func TestAppPasswordImporter(t *testing.T) {
	plugintest.TestImporter(t, AppPassword().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"SKEET": "kazhj4oy2wjfexample",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.: "kazhj4oy2wjfexample",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in skeet/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "kazhj4oy2wjfexample",
			// 		},
			// 	},
			},
		},
	})
}
