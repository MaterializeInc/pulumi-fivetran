// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fivetran

import (
	"fmt"
	"unicode"

	"github.com/benesch/terraform-provider-fivetran/fivetran"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

const (
	mainPkg = "fivetran"
	mainMod = "index"
)

// makeMember manufactures a type token for the package and the given module and
// type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and
// type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and
// resource name.  It automatically uses the main package and names the file by
// simply lower casing the data source's first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and
// resource name.  It automatically uses the main package and names the file by
// simply lower casing the resource's first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the
// provider.
func Provider(version string) tfbridge.ProviderInfo {
	p := shimv2.NewProvider(fivetran.Provider())

	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "fivetran",
		Description:       "A Pulumi package for creating and managing Fivetran resources.",
		Keywords:          []string{"pulumi", "fivetran"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/benesch/pulumi-fivetran",
		Repository:        "https://github.com/benesch/pulumi-fivetran",
		PluginDownloadURL: fmt.Sprintf("https://github.com/benesch/pulumi-fivetran/releases/download/v%s/", version),

		Resources: map[string]*tfbridge.ResourceInfo{
			"fivetran_user":        {Tok: makeResource(mainMod, "User")},
			"fivetran_group":       {Tok: makeResource(mainMod, "Group")},
			"fivetran_destination": {Tok: makeResource(mainMod, "Destination")},
			"fivetran_connector": {
				Tok: makeResource(mainMod, "Connector"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"config": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"last_synced_changes__utc_": {
									Name: "last_synced_changes_utc",
								},
							},
						},
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"fivetran_user":                {Tok: makeDataSource(mainMod, "getUser")},
			"fivetran_users":               {Tok: makeDataSource(mainMod, "getUsers")},
			"fivetran_group":               {Tok: makeDataSource(mainMod, "getGroup")},
			"fivetran_groups":              {Tok: makeDataSource(mainMod, "getGroups")},
			"fivetran_group_connectors":    {Tok: makeDataSource(mainMod, "getGroupConnectors")},
			"fivetran_group_users":         {Tok: makeDataSource(mainMod, "getGroupUsers")},
			"fivetran_destination":         {Tok: makeDataSource(mainMod, "getDestination")},
			"fivetran_connectors_metadata": {Tok: makeDataSource(mainMod, "getConnectorsMetadata")},
			"fivetran_connector": {
				Tok: makeDataSource(mainMod, "getConnector"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"config": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"last_synced_changes__utc_": {
									Name: "last_synced_changes_utc",
								},
							},
						},
					},
				},
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
