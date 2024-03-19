package linearise

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestAPIKeyProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, APIKey().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.APIKey: "lin_api_iPUvhFnbLvvW6amwStmoNmXe2oJ47t1MiEXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"LINEARISE__LINEAR_API_KEY": "lin_api_iPUvhFnbLvvW6amwStmoNmXe2oJ47t1MiEXAMPLE",
				},
			},
		},
	})
}

func TestAPIKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, APIKey().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"LINEARISE__LINEAR_API_KEY": "lin_api_iPUvhFnbLvvW6amwStmoNmXe2oJ47t1MiEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.APIKey: "lin_api_iPUvhFnbLvvW6amwStmoNmXe2oJ47t1MiEXAMPLE",
					},
				},
			},
		},
	})
}
