package jumpiest

import (
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func JumpCloudCLI() schema.Executable {
	return schema.Executable{
		Name: "Estima JumpCloud utility",
		Runs: []string{"jumpiest"},
		// DocsURL: sdk.URL("https://jumpiest.com/docs/cli"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.APIKey,
			},
		},
	}
}
