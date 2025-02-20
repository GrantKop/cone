// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PolicyType - The enum of the policy type.
type PolicyType string

const (
	PolicyTypePolicyTypeUnspecified   PolicyType = "POLICY_TYPE_UNSPECIFIED"
	PolicyTypePolicyTypeGrant         PolicyType = "POLICY_TYPE_GRANT"
	PolicyTypePolicyTypeRevoke        PolicyType = "POLICY_TYPE_REVOKE"
	PolicyTypePolicyTypeCertify       PolicyType = "POLICY_TYPE_CERTIFY"
	PolicyTypePolicyTypeAccessRequest PolicyType = "POLICY_TYPE_ACCESS_REQUEST"
	PolicyTypePolicyTypeProvision     PolicyType = "POLICY_TYPE_PROVISION"
)

func (e PolicyType) ToPointer() *PolicyType {
	return &e
}
func (e *PolicyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POLICY_TYPE_UNSPECIFIED":
		fallthrough
	case "POLICY_TYPE_GRANT":
		fallthrough
	case "POLICY_TYPE_REVOKE":
		fallthrough
	case "POLICY_TYPE_CERTIFY":
		fallthrough
	case "POLICY_TYPE_ACCESS_REQUEST":
		fallthrough
	case "POLICY_TYPE_PROVISION":
		*e = PolicyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyType: %v", v)
	}
}

// The CreatePolicyRequest message is used to create a new policy.
type CreatePolicyRequest struct {
	// The description of the new policy.
	Description *string `json:"description,omitempty"`
	// The display name of the new policy.
	DisplayName *string `json:"displayName,omitempty"`
	// The map of policy type to policy steps. The key is the stringified version of the enum. See other policies for examples.
	PolicySteps map[string]PolicySteps `json:"policySteps,omitempty"`
	// The enum of the policy type.
	PolicyType *PolicyType `json:"policyType,omitempty"`
	// Actions to occur after a policy finishes. As of now this is only valid on a certify policy to remediate a denied certification immediately.
	PostActions []PolicyPostActions `json:"postActions,omitempty"`
	// Allows reassigning tasks to delegates.
	ReassignTasksToDelegates *bool `json:"reassignTasksToDelegates,omitempty"`
	// The rules field.
	Rules []Rule `json:"rules,omitempty"`
}

func (o *CreatePolicyRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreatePolicyRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *CreatePolicyRequest) GetPolicySteps() map[string]PolicySteps {
	if o == nil {
		return nil
	}
	return o.PolicySteps
}

func (o *CreatePolicyRequest) GetPolicyType() *PolicyType {
	if o == nil {
		return nil
	}
	return o.PolicyType
}

func (o *CreatePolicyRequest) GetPostActions() []PolicyPostActions {
	if o == nil {
		return nil
	}
	return o.PostActions
}

func (o *CreatePolicyRequest) GetReassignTasksToDelegates() *bool {
	if o == nil {
		return nil
	}
	return o.ReassignTasksToDelegates
}

func (o *CreatePolicyRequest) GetRules() []Rule {
	if o == nil {
		return nil
	}
	return o.Rules
}
