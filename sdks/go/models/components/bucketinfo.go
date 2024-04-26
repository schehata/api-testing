// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// Size - An integer or float representing the fixed bucket size.
// When the bucket field is _time this value is in nanoseconds.
type Size struct {
}

// BucketInfo - The standard mode of operation for Kirby is to create buckets on the _time column,
type BucketInfo struct {
	// Field specifies the field used to create buckets on. Normally this would be _time.
	Field *string `json:"field,omitempty"`
	// An integer or float representing the fixed bucket size.
	// When the bucket field is _time this value is in nanoseconds.
	Size *Size `json:"size,omitempty"`
}

func (o *BucketInfo) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

func (o *BucketInfo) GetSize() *Size {
	if o == nil {
		return nil
	}
	return o.Size
}