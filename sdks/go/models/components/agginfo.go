// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Args struct {
}

// AggInfo captures information about an aggregation
type AggInfo struct {
	// Args specifies any non-field arguments for the aggregation. Fx. [10] for topk(players, 10).
	Args []Args `json:"args,omitempty"`
	// Fields specifies the names of the fields this aggregation is computed on.
	// Fx ["players"] for topk(players, 10)
	Fields []string `json:"fields,omitempty"`
	// Name is the system name of the aggregation, which is the string form of aggregation.Type.
	// If the aggregation is aliased, the alias is stored in the parent FieldInfo
	Name *string `json:"name,omitempty"`
}

func (o *AggInfo) GetArgs() []Args {
	if o == nil {
		return nil
	}
	return o.Args
}

func (o *AggInfo) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *AggInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
