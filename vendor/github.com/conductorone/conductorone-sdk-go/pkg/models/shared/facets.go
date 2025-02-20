// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Facets - Indicates one value of a facet.
type Facets struct {
	// The count of items in this facet.
	Count *int64 `integer:"string" json:"count,omitempty"`
	// The facet being referenced.
	Facets []FacetCategory `json:"facets,omitempty"`
}

func (o *Facets) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *Facets) GetFacets() []FacetCategory {
	if o == nil {
		return nil
	}
	return o.Facets
}
