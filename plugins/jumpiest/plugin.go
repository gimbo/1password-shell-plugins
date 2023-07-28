package jumpiest

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "jumpiest",
		Platform: schema.PlatformInfo{
			Name:     "JumpCloud",
			Homepage: sdk.URL("https://docs.jumpcloud.com/api/index.html"),
		},
		Credentials: []schema.CredentialType{
			APIKey(),
		},
		Executables: []schema.Executable{
			JumpCloudCLI(),
		},
	}
}
