package jumpiest

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
				fieldname.APIKey:       "qgdg7dv481hkbi484h2ymgj0669a0bnugexample",
				fieldname.Organization: "607dfac69ab3417a8example",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"JUMPCLOUD_API_KEY": "qgdg7dv481hkbi484h2ymgj0669a0bnugexample",
					"JUMPCLOUD_ORG_ID":  "607dfac69ab3417a8example",
				},
			},
		},
	})
}

func TestAPIKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, APIKey().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"JUMPCLOUD_API_KEY": "qgdg7dv481hkbi484h2ymgj0669a0bnugexample",
				"JUMPCLOUD_ORG_ID":  "607dfac69ab3417a8example",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.APIKey:       "qgdg7dv481hkbi484h2ymgj0669a0bnugexample",
						fieldname.Organization: "607dfac69ab3417a8example",
					},
				},
			},
		},
	})
}
