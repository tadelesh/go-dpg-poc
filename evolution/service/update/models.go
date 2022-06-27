//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package update

// ParamsClientDeleteParametersOptions contains the optional parameters for the ParamsClient.DeleteParameters method.
type ParamsClientDeleteParametersOptions struct {
	// placeholder for future optional parameters
}

// ParamsClientGetNewOperationOptions contains the optional parameters for the ParamsClient.GetNewOperation method.
type ParamsClientGetNewOperationOptions struct {
	// placeholder for future optional parameters
}

// ParamsClientGetOptionalOptions contains the optional parameters for the ParamsClient.GetOptional method.
type ParamsClientGetOptionalOptions struct {
	// I'm a new input optional parameter
	NewParameter *string
	// I am an optional parameter
	OptionalParam *string
}

// ParamsClientGetRequiredOptions contains the optional parameters for the ParamsClient.GetRequired method.
type ParamsClientGetRequiredOptions struct {
	// I'm a new input optional parameter
	NewParameter *string
}

// ParamsClientHeadNoParamsOptions contains the optional parameters for the ParamsClient.HeadNoParams method.
type ParamsClientHeadNoParamsOptions struct {
	// I'm a new input optional parameter
	NewParameter *string
}

// ParamsClientPostParametersOptions contains the optional parameters for the ParamsClient.PostParameters method.
type ParamsClientPostParametersOptions struct {
	// placeholder for future optional parameters
}

// ParamsClientPostParametersWithJSONOptions contains the optional parameters for the ParamsClient.PostParametersWithJSON
// method.
type ParamsClientPostParametersWithJSONOptions struct {
	// placeholder for future optional parameters
}

// ParamsClientPutRequiredOptionalOptions contains the optional parameters for the ParamsClient.PutRequiredOptional method.
type ParamsClientPutRequiredOptionalOptions struct {
	// I'm a new input optional parameter
	NewParameter *string
	// I am an optional parameter
	OptionalParam *string
}

type PostInput struct {
	// REQUIRED
	URL *string `json:"url,omitempty"`
}

