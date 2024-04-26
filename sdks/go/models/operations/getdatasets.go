// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"axiom/models/components"
)

type GetDatasetsResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	DatasetSpecs []components.DatasetSpec
}

func (o *GetDatasetsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetDatasetsResponse) GetDatasetSpecs() []components.DatasetSpec {
	if o == nil {
		return nil
	}
	return o.DatasetSpecs
}