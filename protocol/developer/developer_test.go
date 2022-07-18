// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package developer

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/tidwall/gjson"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
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
	resp, err := client.GetModelRaw(context.Background(), "raw", nil)
	require.NoError(t, err)
	result := gjson.ParseBytes(resp)
	require.Equal(t, result.Get("received").String(), "raw")
}

// GetHandwrittenModel - Read a model from a GET
func TestGetHandwrittenModel(t *testing.T) {
	client := newDPGClient()
	resp, err := client.GetModelRaw(context.Background(), "model", nil)
	require.NoError(t, err)
	value := &Product{}
	err = json.Unmarshal(resp, value)
	require.NoError(t, err)
	require.NotNil(t, value.Received, "model")
}

// PostRawModel - Post a JSON {"hello": "world!"}
func TestPostRawModel(t *testing.T) {
	client := newDPGClient()
	payload := `{"hello": "world!"}`
	options := RequestOptions{
		Body:        streaming.NopCloser(strings.NewReader(payload)),
		ContentType: "application/json",
	}
	_, err := client.PostModelRaw(context.Background(), "raw", &options)
	require.NoError(t, err)
}

// PostHandwrittenModel - Pass a model that will serialize as {"hello": "world!"}
func TestPostHandwrittenModel(t *testing.T) {
	client := newDPGClient()
	payload, err := json.Marshal(&Input{Hello: to.Ptr("world!")})
	require.NoError(t, err)
	options := RequestOptions{
		Body:        streaming.NopCloser(bytes.NewReader(payload)),
		ContentType: "application/json",
	}
	_, err = client.PostModelRaw(context.Background(), "model", &options)
	require.NoError(t, err)
}

// GetRawPages - Read the second page
func TestGetRawPages(t *testing.T) {
	client := newDPGClient()
	pager := client.NewGetPagesPagerRaw("raw", nil)
	result := []gjson.Result{}
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageResult := gjson.ParseBytes(page)
		result = append(result, pageResult.Get("values").Array()...)
	}
	require.Equal(t, result[len(result)-1].Get("received").String(), "raw")
}

// GetHandwrittenModelPages - Read the second page
func TestGetHandwrittenModelPages(t *testing.T) {
	client := newDPGClient()
	pager := client.NewGetPagesPagerRaw("model", nil)
	result := []*Product{}
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		pageValue := &ProductResult{}
		err = json.Unmarshal(page, pageValue)
		require.NoError(t, err)
		result = append(result, pageValue.Values...)
	}
	require.Equal(t, *result[len(result)-1].Received, ProductReceived("model"))
}

// RawLRO - Read a polling result as JSON
func TestRawLRO(t *testing.T) {
	client := newDPGClient()
	poller, err := client.BeginLroRaw(context.Background(), "raw", nil)
	require.NoError(t, err)
	resp, err := poller.PollUntilDone(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, resp.(map[string]interface{})["received"].(string), "raw")
}

// HandwrittenModelLRO - HandwrittenModelLRO
func TestHandwrittenModelLRO(t *testing.T) {
	client := newDPGClient()
	poller, err := client.BeginLro(context.Background(), "model", nil)
	require.NoError(t, err)
	result, err := poller.PollUntilDone(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, *result.Received, ProductReceived("model"))
}

func TestGlassBreaker(t *testing.T) {
	pl := runtime.NewPipeline("developer", "1.0.0", runtime.PipelineOptions{}, &azcore.ClientOptions{})
	client := NewGlassBreakerClient(pl)
	resp, err := client.Do(context.Background(), "/servicedriven/glassbreaker", http.MethodGet, nil)
	require.NoError(t, err)
	value := &map[string]interface{}{}
	err = json.Unmarshal(resp, value)
	require.NoError(t, err)
	require.Equal(t, (*value)["message"].(string), "An object was successfully returned")
}
