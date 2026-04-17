// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"fmt"

	"github.com/xsen84/go-netbox-go/common"
)

// CustomFieldChoiceSetChoice represents a single choice within a choice set.
// NetBox returns choices as 2-element arrays: [["value","label"], ...],
// so we need custom JSON unmarshaling.
type CustomFieldChoiceSetChoice struct {
	Value   string
	Display string
}

// UnmarshalJSON handles NetBox's tuple format: ["value", "label"]
func (c *CustomFieldChoiceSetChoice) UnmarshalJSON(data []byte) error {
	var tuple []string
	if err := json.Unmarshal(data, &tuple); err != nil {
		return fmt.Errorf("choice must be a 2-element string array: %w", err)
	}
	if len(tuple) < 2 {
		return fmt.Errorf("choice array must have at least 2 elements, got %d", len(tuple))
	}
	c.Value = tuple[0]
	c.Display = tuple[1]
	return nil
}

// MarshalJSON writes back as a 2-element array for consistency
func (c CustomFieldChoiceSetChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{c.Value, c.Display})
}

// CustomFieldChoiceSet represents a Netbox custom field choice set
type CustomFieldChoiceSet struct {
	ID               int                          `json:"id"`
	URL              string                       `json:"url,omitempty"`
	Display          string                       `json:"display,omitempty"`
	Name             string                       `json:"name"`
	Description      string                       `json:"description,omitempty"`
	ExtraChoices     []CustomFieldChoiceSetChoice  `json:"extra_choices"`
	ChoicesCount     int                           `json:"choices_count,omitempty"`
}

// ListCustomFieldChoiceSetsRequest holds query options for listing choice sets
type ListCustomFieldChoiceSetsRequest struct {
	common.ListParams
}

// ListCustomFieldChoiceSetsResponse is the paginated response for choice sets
type ListCustomFieldChoiceSetsResponse struct {
	common.ReturnValues
	Results []CustomFieldChoiceSet `json:"results"`
}
