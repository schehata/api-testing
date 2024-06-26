// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type TrimResult struct {
	// numDeleted has been deprecated on 2022-09-14.
	// There is currently no way to tell how much was trimmed via the trim result.
	// If you need to check how much was deleted you can either query the removed time range,
	// or poll the dataset info endpoint.
	NumDeleted int64 `json:"numDeleted"`
}

func (o *TrimResult) GetNumDeleted() int64 {
	if o == nil {
		return 0
	}
	return o.NumDeleted
}
