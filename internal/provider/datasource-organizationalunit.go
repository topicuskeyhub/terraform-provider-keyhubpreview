// Code generated by "terraform-provider-keyhub-generator"; DO NOT EDIT.
// Copyright (c) Topicus Security B.V.
// SPDX-License-Identifier: APSL-2.0

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
	keyhubreq "github.com/topicuskeyhub/sdk-go/organizationalunit"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &organizationalunitDataSource{}
	_ datasource.DataSourceWithConfigure = &organizationalunitDataSource{}
)

func NewOrganizationalunitDataSource() datasource.DataSource {
	return &organizationalunitDataSource{}
}

type organizationalunitDataSource struct {
	providerData *KeyHubProviderData
}

func (d *organizationalunitDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = ProviderName + "_organizationalunit"
	log.Printf("Registered data source %s", resp.TypeName)
}

func (d *organizationalunitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: dataSourceSchemaAttrsOrganizationOrganizationalUnit(true),
	}
}

func (d *organizationalunitDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*KeyHubProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *keyhub.KeyHubClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.providerData = providerData
}

func (d *organizationalunitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data organizationOrganizationalUnitDataDS
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading organizationalunit from Topicus KeyHub by UUID")
	listValue, _ := data.Additional.ToListValue(ctx)
	additional, _ := tfToSlice(listValue, func(val attr.Value, diags *diag.Diagnostics) string {
		return val.(basetypes.StringValue).ValueString()
	})
	uuid := data.UUID.ValueString()

	d.providerData.Mutex.RLock()
	defer d.providerData.Mutex.RUnlock()
	wrapper, err := d.providerData.Client.Organizationalunit().Get(ctx, &keyhubreq.OrganizationalunitRequestBuilderGetRequestConfiguration{
		QueryParameters: &keyhubreq.OrganizationalunitRequestBuilderGetQueryParameters{
			Uuid:       []string{uuid},
			Additional: additional,
		},
	})

	tkh, diags := findFirst[keyhubmodels.OrganizationOrganizationalUnitable](ctx, wrapper, "organizationalunit", &uuid, false, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectDSOrganizationOrganizationalUnit(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	fillDataStructFromTFObjectDSOrganizationOrganizationalUnit(&data, tf)
	data.Additional = listValue

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
