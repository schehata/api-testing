// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"axiom/models/components"
)

type GetCurrentUserResponse struct {
	HTTPMeta components.HTTPMetadata
	// User
	User *components.User
}

func (o *GetCurrentUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCurrentUserResponse) GetUser() *components.User {
	if o == nil {
		return nil
	}
	return o.User
}
