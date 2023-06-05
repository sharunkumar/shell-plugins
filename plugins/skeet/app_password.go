package skeet

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func AppPassword() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.AppPassword,
		DocsURL:       sdk.URL("https://skeet.com/docs/app_password"),            // TODO: Replace with actual URL
		ManagementURL: sdk.URL("https://console.skeet.com/user/security/tokens"), // TODO: Replace with actual URL
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Username,
				MarkdownDescription: " used to authenticate to Bluesky.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 19,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                fieldname.Password,
				MarkdownDescription: " used to authenticate to Bluesky.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 19,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
			// TryBlueskyConfigFile(),
		)}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"SKEET_USERNAME": fieldname.Username,
	"SKEET_PASSWORD": fieldname.Password,
}

// // TODO: Check if the platform stores the App Password in a local config file, and if so,
// // implement the function below to add support for importing it.
// func TryBlueskyConfigFile() sdk.Importer {
// 	return importer.TryFile("~/path/to/config/file.yml", func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
// 		// var config Config
// 		// if err := contents.ToYAML(&config); err != nil {
// 		// 	out.AddError(err)
// 		// 	return
// 		// }

// 		// if config. == "" {
// 		// 	return
// 		// }

// 		// out.AddCandidate(sdk.ImportCandidate{
// 		// 	Fields: map[sdk.FieldName]string{
// 		// 		fieldname.: config.,
// 		// 	},
// 		// })
// 	})
// }

// // TODO: Implement the config file schema
// // type Config struct {
// //	 string
// // }
