// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type QueryOptions struct {
	Against            *string `json:"against,omitempty"`
	AgainstStart       *string `json:"againstStart,omitempty"`
	AgainstTimestamp   *string `json:"againstTimestamp,omitempty"`
	AggChartOpts       *string `json:"aggChartOpts,omitempty"`
	CaseSensitive      *string `json:"caseSensitive,omitempty"`
	ContainsTimeFilter *string `json:"containsTimeFilter,omitempty"`
	Datasets           *string `json:"datasets,omitempty"`
	DisplayNull        *string `json:"displayNull,omitempty"`
	EditorContent      *string `json:"editorContent,omitempty"`
	EndColumn          *string `json:"endColumn,omitempty"`
	EndLineNumber      *string `json:"endLineNumber,omitempty"`
	EndTime            *string `json:"endTime,omitempty"`
	IntegrationsFilter *string `json:"integrationsFilter,omitempty"`
	OpenIntervals      *string `json:"openIntervals,omitempty"`
	QuickRange         *string `json:"quickRange,omitempty"`
	Resolution         *string `json:"resolution,omitempty"`
	ShownColumns       *string `json:"shownColumns,omitempty"`
	StartColumn        *string `json:"startColumn,omitempty"`
	StartLineNumber    *string `json:"startLineNumber,omitempty"`
	StartTime          *string `json:"startTime,omitempty"`
	TimeSeriesVariant  *string `json:"timeSeriesVariant,omitempty"`
	TimeSeriesView     *string `json:"timeSeriesView,omitempty"`
}

func (o *QueryOptions) GetAgainst() *string {
	if o == nil {
		return nil
	}
	return o.Against
}

func (o *QueryOptions) GetAgainstStart() *string {
	if o == nil {
		return nil
	}
	return o.AgainstStart
}

func (o *QueryOptions) GetAgainstTimestamp() *string {
	if o == nil {
		return nil
	}
	return o.AgainstTimestamp
}

func (o *QueryOptions) GetAggChartOpts() *string {
	if o == nil {
		return nil
	}
	return o.AggChartOpts
}

func (o *QueryOptions) GetCaseSensitive() *string {
	if o == nil {
		return nil
	}
	return o.CaseSensitive
}

func (o *QueryOptions) GetContainsTimeFilter() *string {
	if o == nil {
		return nil
	}
	return o.ContainsTimeFilter
}

func (o *QueryOptions) GetDatasets() *string {
	if o == nil {
		return nil
	}
	return o.Datasets
}

func (o *QueryOptions) GetDisplayNull() *string {
	if o == nil {
		return nil
	}
	return o.DisplayNull
}

func (o *QueryOptions) GetEditorContent() *string {
	if o == nil {
		return nil
	}
	return o.EditorContent
}

func (o *QueryOptions) GetEndColumn() *string {
	if o == nil {
		return nil
	}
	return o.EndColumn
}

func (o *QueryOptions) GetEndLineNumber() *string {
	if o == nil {
		return nil
	}
	return o.EndLineNumber
}

func (o *QueryOptions) GetEndTime() *string {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *QueryOptions) GetIntegrationsFilter() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationsFilter
}

func (o *QueryOptions) GetOpenIntervals() *string {
	if o == nil {
		return nil
	}
	return o.OpenIntervals
}

func (o *QueryOptions) GetQuickRange() *string {
	if o == nil {
		return nil
	}
	return o.QuickRange
}

func (o *QueryOptions) GetResolution() *string {
	if o == nil {
		return nil
	}
	return o.Resolution
}

func (o *QueryOptions) GetShownColumns() *string {
	if o == nil {
		return nil
	}
	return o.ShownColumns
}

func (o *QueryOptions) GetStartColumn() *string {
	if o == nil {
		return nil
	}
	return o.StartColumn
}

func (o *QueryOptions) GetStartLineNumber() *string {
	if o == nil {
		return nil
	}
	return o.StartLineNumber
}

func (o *QueryOptions) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *QueryOptions) GetTimeSeriesVariant() *string {
	if o == nil {
		return nil
	}
	return o.TimeSeriesVariant
}

func (o *QueryOptions) GetTimeSeriesView() *string {
	if o == nil {
		return nil
	}
	return o.TimeSeriesView
}