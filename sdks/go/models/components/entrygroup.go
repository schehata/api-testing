// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Group struct {
}

type EntryGroup struct {
	Aggregations []EntryGroupAgg  `json:"aggregations,omitempty"`
	Group        map[string]Group `json:"group"`
	ID           int64            `json:"id"`
}

func (o *EntryGroup) GetAggregations() []EntryGroupAgg {
	if o == nil {
		return nil
	}
	return o.Aggregations
}

func (o *EntryGroup) GetGroup() map[string]Group {
	if o == nil {
		return map[string]Group{}
	}
	return o.Group
}

func (o *EntryGroup) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}
