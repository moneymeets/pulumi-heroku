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

package heroku

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/moneymeets/pulumi-heroku/provider/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/heroku/terraform-provider-heroku/v4/heroku"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "heroku"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := heroku.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "heroku",
		Description: "A Pulumi package for creating and managing heroku cloud resources.",
		Keywords:    []string{"pulumi", "heroku"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/moneymeets/pulumi-heroku",
		Config: map[string]*tfbridge.SchemaInfo{
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"heroku_account_feature":                   {Tok: makeResource(mainMod, "HerokuAccountFeature")},
			"heroku_addon":                             {Tok: makeResource(mainMod, "HerokuAddon")},
			"heroku_addon_attachment":                  {Tok: makeResource(mainMod, "HerokuAddonAttachment")},
			"heroku_app":                               {Tok: makeResource(mainMod, "HerokuApp")},
			"heroku_app_config_association":            {Tok: makeResource(mainMod, "HerokuAppConfigAssociation")},
			"heroku_app_feature":                       {Tok: makeResource(mainMod, "HerokuAppFeature")},
			"heroku_app_release":                       {Tok: makeResource(mainMod, "HerokuAppRelease")},
			"heroku_app_webhook":                       {Tok: makeResource(mainMod, "HerokuAppWebhook")},
			"heroku_build":                             {Tok: makeResource(mainMod, "HerokuBuild")},
			"heroku_cert":                              {Tok: makeResource(mainMod, "HerokuCert")},
			"heroku_config":                            {Tok: makeResource(mainMod, "HerokuConfig")},
			"heroku_domain":                            {Tok: makeResource(mainMod, "HerokuDomain")},
			"heroku_drain":                             {Tok: makeResource(mainMod, "HerokuDrain")},
			"heroku_formation":                         {Tok: makeResource(mainMod, "HerokuFormation")},
			"heroku_pipeline":                          {Tok: makeResource(mainMod, "HerokuPipeline")},
			"heroku_pipeline_config_var":               {Tok: makeResource(mainMod, "HerokuPipelineConfigVar")},
			"heroku_pipeline_coupling":                 {Tok: makeResource(mainMod, "HerokuPipelineCoupling")},
			"heroku_slug":                              {Tok: makeResource(mainMod, "HerokuSlug")},
			"heroku_space":                             {Tok: makeResource(mainMod, "HerokuSpace")},
			"heroku_space_app_access":                  {Tok: makeResource(mainMod, "HerokuSpaceAppAccess")},
			"heroku_space_inbound_ruleset":             {Tok: makeResource(mainMod, "HerokuSpaceInboundRuleset")},
			"heroku_space_peering_connection_accepter": {Tok: makeResource(mainMod, "HerokuSpacePeeringConnectionAccepter")},
			"heroku_space_vpn_connection":              {Tok: makeResource(mainMod, "HerokuSpaceVpnConnection")},
			"heroku_team_collaborator":                 {Tok: makeResource(mainMod, "HerokuTeamCollaborator")},
			"heroku_team_member":                       {Tok: makeResource(mainMod, "HerokuTeamMember")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"heroku_addon":              {Tok: makeDataSource(mainMod, "getHerokuAddon")},
			"heroku_app":                {Tok: makeDataSource(mainMod, "getHerokuApp")},
			"heroku_space":              {Tok: makeDataSource(mainMod, "getHerokuSpace")},
			"heroku_space_peering_info": {Tok: makeDataSource(mainMod, "getHerokuSpacePeeringInfo")},
			"heroku_team":               {Tok: makeDataSource(mainMod, "getHerokuTeam")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
