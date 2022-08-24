package developer

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
)

type RequestOptions struct {
	Header      http.Header
	QueryParam  map[string]string // type alias and provide set add method
	Body        io.ReadSeekCloser
	ContentType string // find some way to infer from body
}

func (options *RequestOptions) AddQueryParam(name, value string) {
	if options.QueryParam == nil {
		options.QueryParam = map[string]string{name: value}
	} else {
		options.QueryParam[name] = value
	}
}

func NewJSONStringRequest(body string) RequestOptions {
	return RequestOptions{
		Body:        streaming.NopCloser(strings.NewReader(body)),
		ContentType: "application/json",
	}
}

func NewFormDataRequest(data url.Values) RequestOptions {
	return RequestOptions{
		Body:        streaming.NopCloser(strings.NewReader(data.Encode())),
		ContentType: "application/x-www-form-urlencoded",
	}
}

// TODO: add more ctor for RequestOptions, plain text, files, multiform etc.

type LRORequestOptions struct {
	RequestOptions
	ResumeToken string
}
