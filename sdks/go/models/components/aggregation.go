// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Argument struct {
}

type Op string

const (
	OpCount       Op = "count"
	OpDistinct    Op = "distinct"
	OpSum         Op = "sum"
	OpAvg         Op = "avg"
	OpMin         Op = "min"
	OpMax         Op = "max"
	OpTopk        Op = "topk"
	OpPercentiles Op = "percentiles"
	OpHistogram   Op = "histogram"
	OpStdev       Op = "stdev"
	OpVariance    Op = "variance"
	OpArgmin      Op = "argmin"
	OpArgmax      Op = "argmax"
	OpMakeset     Op = "makeset"
	OpRate        Op = "rate"
	OpMakelist    Op = "makelist"
)

func (e Op) ToPointer() *Op {
	return &e
}

func (e *Op) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "count":
		fallthrough
	case "distinct":
		fallthrough
	case "sum":
		fallthrough
	case "avg":
		fallthrough
	case "min":
		fallthrough
	case "max":
		fallthrough
	case "topk":
		fallthrough
	case "percentiles":
		fallthrough
	case "histogram":
		fallthrough
	case "stdev":
		fallthrough
	case "variance":
		fallthrough
	case "argmin":
		fallthrough
	case "argmax":
		fallthrough
	case "makeset":
		fallthrough
	case "rate":
		fallthrough
	case "makelist":
		*e = Op(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Op: %v", v)
	}
}

type Aggregation struct {
	Alias    *string   `json:"alias,omitempty"`
	Argument *Argument `json:"argument,omitempty"`
	Field    string    `json:"field"`
	Op       Op        `json:"op"`
}

func (o *Aggregation) GetAlias() *string {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *Aggregation) GetArgument() *Argument {
	if o == nil {
		return nil
	}
	return o.Argument
}

func (o *Aggregation) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *Aggregation) GetOp() Op {
	if o == nil {
		return Op("")
	}
	return o.Op
}