// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package developer

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func newDPGClient() *DPGClient {
	pl := runtime.NewPipeline("developer", "1.0.0", runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewDPGClient(pl)
}

// GetRawModel - Read a JSON from a GET
func TestGetRawModel(t *testing.T) {
	client := newDPGClient()
	resp, err := client.GetModel(context.Background(), "raw", nil)
	require.NoError(t, err)
	require.Equal(t, resp.Received, "raw") // should get raw JSON response
}

// GetHandwrittenModel - Read a model from a GET
func TestGetHandwrittenModel(t *testing.T) {
	client := newDPGClient()
	resp, err := client.GetModel(context.Background(), "model", nil)
	require.NoError(t, err)
	require.NotNil(t, resp.Received, "model")
}

// PostRawModel - Post a JSON {"hello": "world!"}
func TestPostRawModel(t *testing.T) {
	client := newDPGClient()
	_, err := client.PostModel(context.Background(), "raw", Input{Hello: to.Ptr("world")}, nil) // should post raw JSON
	require.NoError(t, err)
}

// PostHandwrittenModel - Pass a model that will serialize as {"hello": "world!"}
func TestPostHandwrittenModel(t *testing.T) {
	client := newDPGClient()
	_, err := client.PostModel(context.Background(), "model", Input{Hello: to.Ptr("world")}, nil)
	require.NoError(t, err)
}

// GetRawPages - Read the second page
func TestGetRawPages(t *testing.T) {
	client := newDPGClient()
	pager := client.NewGetPagesPager("raw", nil)
	result := []*Product{} // should get raw JSON result
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		result = append(result, page.Values...)
	}
	require.Equal(t, result[len(result)-1].Received, "raw")
}

// GetHandwrittenModelPages - Read the second page
func TestGetHandwrittenModelPages(t *testing.T) {
	client := newDPGClient()
	pager := client.NewGetPagesPager("model", nil)
	result := []*Product{}
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		result = append(result, page.Values...)
	}
	require.Equal(t, result[len(result)-1].Received, "model")
}

// RawLRO - Read a polling result as JSON
func TestRawLRO(t *testing.T) {
	client := newDPGClient()
	poller, err := client.BeginLro(context.Background(), "raw", nil)
	require.NoError(t, err)
	result, err := poller.PollUntilDone(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, result.Received, "raw") // should get raw JSON response
}

// HandwrittenModelLRO - HandwrittenModelLRO
func TestHandwrittenModelLRO(t *testing.T) {
	client := newDPGClient()
	poller, err := client.BeginLro(context.Background(), "model", nil)
	require.NoError(t, err)
	result, err := poller.PollUntilDone(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, result.Received, "model")
}
