// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type IngestStatus struct {
	// Number of blocks created
	BlocksCreated *int64 `json:"blocksCreated,omitempty"`
	// Number of failures that occurred during ingest
	Failed int64 `json:"failed"`
	// List of failures that occurred during ingest
	Failures []IngestFailure `json:"failures,omitempty"`
	// Number of events ingested
	Ingested int64 `json:"ingested"`
	// Number of bytes processed
	ProcessedBytes int64 `json:"processedBytes"`
	// Length of the WAL
	WalLength *int64 `json:"walLength,omitempty"`
}

func (o *IngestStatus) GetBlocksCreated() *int64 {
	if o == nil {
		return nil
	}
	return o.BlocksCreated
}

func (o *IngestStatus) GetFailed() int64 {
	if o == nil {
		return 0
	}
	return o.Failed
}

func (o *IngestStatus) GetFailures() []IngestFailure {
	if o == nil {
		return nil
	}
	return o.Failures
}

func (o *IngestStatus) GetIngested() int64 {
	if o == nil {
		return 0
	}
	return o.Ingested
}

func (o *IngestStatus) GetProcessedBytes() int64 {
	if o == nil {
		return 0
	}
	return o.ProcessedBytes
}

func (o *IngestStatus) GetWalLength() *int64 {
	if o == nil {
		return nil
	}
	return o.WalLength
}