// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"axiom/models/components"
)

type CreateDatasetResponse struct {
	HTTPMeta components.HTTPMetadata
	// Created
	DatasetSpec *components.DatasetSpec
}

func (o *CreateDatasetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateDatasetResponse) GetDatasetSpec() *components.DatasetSpec {
	if o == nil {
		return nil
	}
	return o.DatasetSpec
}