// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	vaultsdk "github.com/hashicorp/vault/sdk/helper/roottoken"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = RootTokenDecodeFunction{}
)

func NewRootTokenDecodeFunction() function.Function {
	return RootTokenDecodeFunction{}
}

type RootTokenDecodeFunction struct {
}

func (r RootTokenDecodeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "token"
}

func (r RootTokenDecodeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Decodes a Vault root token",
		MarkdownDescription: "This function decodes a Vault root token and returns the decoded value.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "token",
				MarkdownDescription: "The Vault root token to decode.",
			},
			function.StringParameter{
				Name:                "one_time_password",
				MarkdownDescription: "The one-time password associated with the token.",
			},
			function.Int64Parameter{
				Name:                "one_time_password_length",
				MarkdownDescription: "The length of the one-time password.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r RootTokenDecodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var token, otp string
	var otpLen int

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &token, &otp, &otpLen))

	if resp.Error != nil {
		return
	}

	// Decode root token
	result, err := vaultsdk.DecodeToken(token, otp, otpLen)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
