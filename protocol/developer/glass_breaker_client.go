//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package developer

import (
	"context"
	"io"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type GlassBreakerClient struct {
	pl runtime.Pipeline
}

func NewGlassBreakerClient(pl runtime.Pipeline) *GlassBreakerClient {
	client := &GlassBreakerClient{
		pl: pl,
	}
	return client
}

func (client *GlassBreakerClient) SendRequest(ctx context.Context, urlPath, httpMethod string, query map[string]string, header http.Header, body io.ReadSeekCloser) ([]byte, error) {
	req, err := runtime.NewRequest(ctx, httpMethod, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if query != nil && len(query) > 0 {
		reqQP := req.Raw().URL.Query()
		for k, v := range query {
			reqQP.Set(k, v)
		}
		req.Raw().URL.RawQuery = reqQP.Encode()
	}
	if header != nil && len(header) > 0 {
		req.Raw().Header = header
	}
	if body != nil {
		req.SetBody(body, req.Raw().Header.Get("Content-Type"))
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return runtime.Payload(resp)
}
