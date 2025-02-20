// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// ConnectorServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type ConnectorServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (c ConnectorServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectorServiceGetResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectorServiceGetResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *ConnectorServiceGetResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The ConnectorServiceGetResponse message contains the connectorView, and an expand mask.
type ConnectorServiceGetResponse struct {
	// The ConnectorView object provides a connector response object, as well as JSONPATHs to related objects provided by expanders.
	ConnectorView *ConnectorView `json:"connectorView,omitempty"`
	// The array of expanded items indicated by the request.
	Expanded []ConnectorServiceGetResponseExpanded `json:"expanded,omitempty"`
}

func (o *ConnectorServiceGetResponse) GetConnectorView() *ConnectorView {
	if o == nil {
		return nil
	}
	return o.ConnectorView
}

func (o *ConnectorServiceGetResponse) GetExpanded() []ConnectorServiceGetResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
