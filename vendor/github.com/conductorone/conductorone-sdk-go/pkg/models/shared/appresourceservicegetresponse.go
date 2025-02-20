// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// AppResourceServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type AppResourceServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (a AppResourceServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppResourceServiceGetResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppResourceServiceGetResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *AppResourceServiceGetResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// AppResourceServiceGetResponse - The app resource service get response contains the app resource view and array of expanded items indicated by the request's expand mask.
type AppResourceServiceGetResponse struct {
	// The app resource view returns an app resource with paths for items in the expand mask filled in when this response is returned and a request expand mask has "*" or "app_id" or "resource_type_id".
	AppResourceView *AppResourceView `json:"appResourceView,omitempty"`
	// List of serialized related objects.
	Expanded []AppResourceServiceGetResponseExpanded `json:"expanded,omitempty"`
}

func (o *AppResourceServiceGetResponse) GetAppResourceView() *AppResourceView {
	if o == nil {
		return nil
	}
	return o.AppResourceView
}

func (o *AppResourceServiceGetResponse) GetExpanded() []AppResourceServiceGetResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
