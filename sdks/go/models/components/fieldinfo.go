// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type FieldInfo struct {
	// AggInfo captures information about an aggregation
	Agg  *AggInfo `json:"agg,omitempty"`
	Name *string  `json:"name,omitempty"`
	Type *string  `json:"type,omitempty"`
}

func (o *FieldInfo) GetAgg() *AggInfo {
	if o == nil {
		return nil
	}
	return o.Agg
}

func (o *FieldInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FieldInfo) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
