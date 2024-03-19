package linearise

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "linearise",
		Platform: schema.PlatformInfo{
			Name:     "Linear",
			Homepage: sdk.URL("https://linear.app"),
		},
		Credentials: []schema.CredentialType{
			APIKey(),
		},
		Executables: []schema.Executable{
			LineariseCLI(),
			CyclomaticCLI(),
		},
	}
}
