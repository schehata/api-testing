// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"axiom/internal/utils"
	"time"
)

type Status struct {
	BlocksExamined    int64     `json:"blocksExamined"`
	CacheStatus       int64     `json:"cacheStatus"`
	ContinuationToken *string   `json:"continuationToken,omitempty"`
	ElapsedTime       int64     `json:"elapsedTime"`
	IsEstimate        *bool     `json:"isEstimate,omitempty"`
	IsPartial         bool      `json:"isPartial"`
	MaxBlockTime      time.Time `json:"maxBlockTime"`
	// Row id of the newest row, as seen server side.
	// May be higher than what the results include if the server scanned more data than included in the results.
	// Can be used to efficiently resume time-sorted non-aggregating queries (ie filtering only).
	MaxCursor    *string   `json:"maxCursor,omitempty"`
	Messages     []Message `json:"messages,omitempty"`
	MinBlockTime time.Time `json:"minBlockTime"`
	// Row id of the oldest row, as seen server side.
	// May be lower than what the results include if the server scanned more data than included in the results.
	// Can be used to efficiently resume time-sorted non-aggregating queries (ie filtering only).
	MinCursor    *string `json:"minCursor,omitempty"`
	NumGroups    int64   `json:"numGroups"`
	RowsExamined int64   `json:"rowsExamined"`
	RowsMatched  int64   `json:"rowsMatched"`
}

func (s Status) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Status) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Status) GetBlocksExamined() int64 {
	if o == nil {
		return 0
	}
	return o.BlocksExamined
}

func (o *Status) GetCacheStatus() int64 {
	if o == nil {
		return 0
	}
	return o.CacheStatus
}

func (o *Status) GetContinuationToken() *string {
	if o == nil {
		return nil
	}
	return o.ContinuationToken
}

func (o *Status) GetElapsedTime() int64 {
	if o == nil {
		return 0
	}
	return o.ElapsedTime
}

func (o *Status) GetIsEstimate() *bool {
	if o == nil {
		return nil
	}
	return o.IsEstimate
}

func (o *Status) GetIsPartial() bool {
	if o == nil {
		return false
	}
	return o.IsPartial
}

func (o *Status) GetMaxBlockTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.MaxBlockTime
}

func (o *Status) GetMaxCursor() *string {
	if o == nil {
		return nil
	}
	return o.MaxCursor
}

func (o *Status) GetMessages() []Message {
	if o == nil {
		return nil
	}
	return o.Messages
}

func (o *Status) GetMinBlockTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.MinBlockTime
}

func (o *Status) GetMinCursor() *string {
	if o == nil {
		return nil
	}
	return o.MinCursor
}

func (o *Status) GetNumGroups() int64 {
	if o == nil {
		return 0
	}
	return o.NumGroups
}

func (o *Status) GetRowsExamined() int64 {
	if o == nil {
		return 0
	}
	return o.RowsExamined
}

func (o *Status) GetRowsMatched() int64 {
	if o == nil {
		return 0
	}
	return o.RowsMatched
}