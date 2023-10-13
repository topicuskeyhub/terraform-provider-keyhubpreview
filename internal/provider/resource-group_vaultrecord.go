// Code generated by "terraform-provider-keyhub-generator"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	keyhub "github.com/topicuskeyhub/sdk-go"
	keyhubreq "github.com/topicuskeyhub/sdk-go/group"
	keyhubmodels "github.com/topicuskeyhub/sdk-go/models"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &groupVaultrecordResource{}
	_ resource.ResourceWithImportState = &groupVaultrecordResource{}
	_ resource.ResourceWithConfigure   = &groupVaultrecordResource{}
)

func NewGroupVaultrecordResource() resource.Resource {
	return &groupVaultrecordResource{}
}

type groupVaultrecordResource struct {
	client *keyhub.KeyHubClient
}

func (r *groupVaultrecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = ProviderName + "_group_vaultrecord"
	tflog.Info(ctx, "Registred resource "+resp.TypeName)
}

func (r *groupVaultrecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: resourceSchemaAttrsGroupVaultVaultRecord(true),
	}
}

func (r *groupVaultrecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*keyhub.KeyHubClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *keyhub.KeyHubClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *groupVaultrecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
	obj, diags := types.ObjectValueFrom(ctx, groupVaultVaultRecordAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSGroupVaultVaultRecord(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Creating Topicus KeyHub group_vaultrecord")
	newWrapper := keyhubmodels.NewVaultVaultRecordLinkableWrapper()
	newWrapper.SetItems([]keyhubmodels.VaultVaultRecordable{newTkh})
	tkhParent, diags := findGroupGroupPrimerByUUID(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wrapper, err := r.client.Group().ByGroupidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().Post(
		ctx, newWrapper, &keyhubreq.ItemVaultRecordRequestBuilderPostRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordRequestBuilderPostQueryParameters{
				Additional: collectAdditional(data.AdditionalObjects),
			},
		})
	tkh, diags := findFirst[keyhubmodels.VaultVaultRecordable](ctx, wrapper, "group_vaultrecord", nil, err)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "group_uuid", types.StringValue(data.GroupUUID.ValueString()))
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Info(ctx, "Created a new Topicus KeyHub group_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupVaultrecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
	tflog.Info(ctx, "Reading group_vaultrecord from Topicus KeyHub")
	tkhParent, diags := findGroupGroupPrimerByUUID(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tkh, err := r.client.Group().ByGroupidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().ByRecordidInt64(getSelfLink(data.Links).ID.ValueInt64()).Get(
		ctx, &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderGetRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderGetQueryParameters{
				Additional: collectAdditional(data.AdditionalObjects),
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "group_uuid", types.StringValue(data.GroupUUID.ValueString()))
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupVaultrecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
	obj, diags := types.ObjectValueFrom(ctx, groupVaultVaultRecordAttrTypesRSRecurse, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTkh, diags := tfObjectToTKHRSGroupVaultVaultRecord(ctx, true, obj)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Updating Topicus KeyHub group_vaultrecord")
	tkhParent, diags := findGroupGroupPrimerByUUID(ctx, data.GroupUUID.ValueStringPointer())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tkh, err := r.client.Group().ByGroupidInt64(*tkhParent.GetLinks()[0].GetId()).Vault().Record().ByRecordidInt64(getSelfLink(data.Links).ID.ValueInt64()).Put(
		ctx, newTkh, &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutRequestConfiguration{
			QueryParameters: &keyhubreq.ItemVaultRecordWithRecordItemRequestBuilderPutQueryParameters{
				Additional: collectAdditional(data.AdditionalObjects),
			},
		})

	if !isHttpStatusCodeOk(ctx, -1, err, &resp.Diagnostics) {
		return
	}

	tf, diags := tkhToTFObjectRSGroupVaultVaultRecord(true, tkh)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tf = setAttributeValue(ctx, tf, "group_uuid", types.StringValue(data.GroupUUID.ValueString()))
	fillDataStructFromTFObjectRSGroupVaultVaultRecord(&data, tf)

	tflog.Info(ctx, "Updated a Topicus KeyHub group_vaultrecord")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *groupVaultrecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data groupVaultVaultRecordDataRS
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = context.WithValue(ctx, keyHubClientKey, r.client)
	tflog.Info(ctx, "Deleting group_vaultrecord from Topicus KeyHub")
	err := r.client.Group().ByGroupidInt64(-1).Vault().Record().ByRecordidInt64(-1).WithUrl(getSelfLink(data.Links).Href.ValueString()).Delete(ctx, nil)
	if !isHttpStatusCodeOk(ctx, 404, err, &resp.Diagnostics) {
		return
	}
	tflog.Info(ctx, "Deleted group_vaultrecord from Topicus KeyHub")
}

func (r *groupVaultrecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("uuid"), req, resp)
}
