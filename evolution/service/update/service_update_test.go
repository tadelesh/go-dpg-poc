// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package update

import (
	"context"
	"os"
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
	resp, err := client.GetRequired(context.Background(), "requiredParam", &ParamsClientGetRequiredOptions{NewParameter: to.Ptr("newOptionalParam")})
	require.NoError(t, err)
	require.NotNil(t, resp.Interface)
}

// DPGAddOptionalInput_NoParams - Initially has no query parameters. After evolution, a new optional query parameter is added
func TestDPGAddOptionalInputNoParams(t *testing.T) {
	client := newParamsClient()
	_, err := client.HeadNoParams(context.Background(), &ParamsClientHeadNoParamsOptions{NewParameter: to.Ptr("newOptionalParam")})
	require.NoError(t, err)
}

// DPGAddOptionalInput_RequiredOptionalParam - Initially has one required query parameter and one optional query parameter.  After evolution, a new optional query parameter is added
func TestDPGAddOptionalInputRequiredOptionalParam(t *testing.T) {
	client := newParamsClient()
	resp, err := client.PutRequiredOptional(context.Background(), "requiredParam", &ParamsClientPutRequiredOptionalOptions{OptionalParam: to.Ptr("optionalParam"), NewParameter: to.Ptr("newOptionalParam")})
	require.NoError(t, err)
	require.NotNil(t, resp.Interface)
}

// DPGAddOptionalInput_OptionalParam - Initially has one optional query parameter. After evolution, a new optional query parameter is added
func TestDPGAddOptionalInputOptionalParam(t *testing.T) {
	client := newParamsClient()
	resp, err := client.GetOptional(context.Background(), &ParamsClientGetOptionalOptions{OptionalParam: to.Ptr("optionalParam"), NewParameter: to.Ptr("newOptionalParam")})
	require.NoError(t, err)
	require.NotNil(t, resp.Interface)
}

// DPGAddNewOperation - Initially the path exists but there is no delete method. After evolution this is a new method in a known path
func TestDPGAddNewOperation(t *testing.T) {
	client := newParamsClient()
	_, err := client.DeleteParameters(context.Background(), nil)
	require.NoError(t, err)
}

// DPGAddNewPath - Initiallty neither path or method exist for this operation. After evolution, this is a new method in a new path
func TestDPGAddNewPath(t *testing.T) {
	client := newParamsClient()
	resp, err := client.GetNewOperation(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Interface)
}

// DPGNewBodyType - Initial version is JSON only, and now supports also JPEG
func TestDPGNewBodyType(t *testing.T) {
	// breaking
	client := newParamsClient()
	_, err := client.PostParametersWithJSON(context.Background(), PostInput{URL: to.Ptr("http://example.org/myimage.jpeg")}, nil)
	require.NoError(t, err)

	file, err := os.OpenFile("test.jpg", os.O_RDONLY, os.ModeAppend)
	require.NoError(t, err)
	_, err = client.PostParameters(context.Background(), file, nil)
	require.NoError(t, err)
}

// DPGGlassBreaker - Not support
