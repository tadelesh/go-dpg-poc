// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package initial

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newParamsClient() *ParamsClient {
	pl := runtime.NewPipeline("developer", "1.0.0", runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewParamsClient(pl)
}

// DPGAddOptionalInput - Initially only has one required Query Parameter. After evolution, a new optional query parameter is added
func TestDPGAddOptionalInput(t *testing.T) {
	client := newParamsClient()
	resp, err := client.GetRequired(context.Background(), "requiredParam", nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Interface)
}

// DPGAddOptionalInput_NoParams - Initially has no query parameters. After evolution, a new optional query parameter is added
func TestDPGAddOptionalInputNoParams(t *testing.T) {
	client := newParamsClient()
	_, err := client.HeadNoParams(context.Background(), nil)
	require.NoError(t, err)
}

// DPGAddOptionalInput_RequiredOptionalParam - Initially has one required query parameter and one optional query parameter.  After evolution, a new optional query parameter is added
func TestDPGAddOptionalInputRequiredOptionalParam(t *testing.T) {
	client := newParamsClient()
	resp, err := client.PutRequiredOptional(context.Background(), "requiredParam", nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Interface)
}

// DPGAddOptionalInput_OptionalParam - Initially has one optional query parameter. After evolution, a new optional query parameter is added
func TestDPGAddOptionalInputOptionalParam(t *testing.T) {
	client := newParamsClient()
	resp, err := client.GetOptional(context.Background(), &ParamsClientGetOptionalOptions{OptionalParam: to.Ptr("optionalParam")})
	require.NoError(t, err)
	require.NotNil(t, resp.Interface)
}

// DPGNewBodyType - Initial version is JSON only, and now supports also JPEG
func TestDPGNewBodyType(t *testing.T) {
	client := newParamsClient()
	_, err := client.PostParameters(context.Background(), PostInput{URL: to.Ptr("http://example.org/myimage.jpeg")}, nil)
	require.NoError(t, err)
}
