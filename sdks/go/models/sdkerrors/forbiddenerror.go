// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// ForbiddenError - Forbidden
type ForbiddenError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

var _ error = &ForbiddenError{}

func (e *ForbiddenError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}