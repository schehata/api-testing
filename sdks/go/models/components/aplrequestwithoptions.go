// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Variables struct {
}

type APLRequestWithOptions struct {
	Apl           string        `json:"apl"`
	Cursor        *string       `json:"cursor,omitempty"`
	EndTime       *string       `json:"endTime,omitempty"`
	IncludeCursor *bool         `json:"includeCursor,omitempty"`
	QueryOptions  *QueryOptions `json:"queryOptions,omitempty"`
	// start and end time for the query, these must be specified as RFC3339 strings
	// or using relative time expressions (e.g. now-1h, now-1d, now-1w, etc)
	StartTime *string `json:"startTime,omitempty"`
	// Variables is an optional set of additional variables that are inserted into the APL
	Variables map[string]Variables `json:"variables,omitempty"`
}

func (o *APLRequestWithOptions) GetApl() string {
	if o == nil {
		return ""
	}
	return o.Apl
}

func (o *APLRequestWithOptions) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *APLRequestWithOptions) GetEndTime() *string {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *APLRequestWithOptions) GetIncludeCursor() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeCursor
}

func (o *APLRequestWithOptions) GetQueryOptions() *QueryOptions {
	if o == nil {
		return nil
	}
	return o.QueryOptions
}

func (o *APLRequestWithOptions) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *APLRequestWithOptions) GetVariables() map[string]Variables {
	if o == nil {
		return nil
	}
	return o.Variables
}
