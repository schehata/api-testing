// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type GroupInfo struct {
	Name *string `json:"name,omitempty"`
}

func (o *GroupInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}