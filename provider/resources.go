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

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(heroku.Provider())

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
			"heroku_account_feature":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuAccountFeature")},
			"heroku_addon":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuAddon")},
			"heroku_addon_attachment":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuAddonAttachment")},
			"heroku_app":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuApp")},
			"heroku_app_config_association":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuAppConfigAssociation")},
			"heroku_app_feature":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuAppFeature")},
			"heroku_app_release":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuAppRelease")},
			"heroku_app_webhook":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuAppWebhook")},
			"heroku_build":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuBuild")},
			"heroku_cert":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuCert")},
			"heroku_collaborator":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuCollaborator")},
			"heroku_config":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuConfig")},
			"heroku_domain":                            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuDomain")},
			"heroku_drain":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuDrain")},
			"heroku_formation":                         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuFormation")},
			"heroku_pipeline":                          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuPipeline")},
			"heroku_pipeline_config_var":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuPipelineConfigVar")},
			"heroku_pipeline_coupling":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuPipelineCoupling")},
			"heroku_review_app_config":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuReviewAppConfig")},
			"heroku_slug":                              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuSlug")},
			"heroku_space":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuSpace")},
			"heroku_space_app_access":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuSpaceAppAccess")},
			"heroku_space_inbound_ruleset":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuSpaceInboundRuleset")},
			"heroku_space_peering_connection_accepter": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuSpacePeeringConnectionAccepter")},
			"heroku_space_vpn_connection":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuSpaceVpnConnection")},
			"heroku_ssl":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuSsl")},
			"heroku_team_collaborator":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuTeamCollaborator")},
			"heroku_team_member":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HerokuTeamMember")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"heroku_addon":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHerokuAddon")},
			"heroku_app":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHerokuApp")},
			"heroku_pipeline":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHerokuPipeline")},
			"heroku_space":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHerokuSpace")},
			"heroku_space_peering_info": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHerokuSpacePeeringInfo")},
			"heroku_team":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHerokuTeam")},
			"heroku_team_members":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getHerokuTeamMembers")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^16.11.7", // so we can access strongly typed node definitions.
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
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
