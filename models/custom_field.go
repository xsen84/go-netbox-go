// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"net/url"

	"github.com/xsen84/go-netbox-go/common"
)

// CustomFieldType represents the type label/value pair for a custom field type
type CustomFieldType struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// CustomFieldChoice represents a single choice in a select-type custom field
type CustomFieldChoice struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// NestedChoiceSet is the nested reference to a choice set within a custom field response
type NestedChoiceSet struct {
	ID           int    `json:"id"`
	URL          string `json:"url,omitempty"`
	Display      string `json:"display,omitempty"`
	ChoicesCount int    `json:"choices_count,omitempty"`
}

// CustomField represents a Netbox custom field definition
type CustomField struct {
	ID            int                 `json:"id"`
	URL           string              `json:"url,omitempty"`
	Display       string              `json:"display,omitempty"`
	ContentTypes  []string            `json:"content_types"`
	Type          CustomFieldType     `json:"type"`
	ObjectType    string              `json:"object_type,omitempty"`
	DataType      string              `json:"data_type,omitempty"`
	Name          string              `json:"name"`
	Label         string              `json:"label"`
	GroupName     string              `json:"group_name,omitempty"`
	Description   string              `json:"description,omitempty"`
	Required      bool                `json:"required"`
	SearchWeight  int                 `json:"search_weight,omitempty"`
	FilterLogic   CustomFieldType     `json:"filter_logic,omitempty"`
	UIVisibility  CustomFieldType     `json:"ui_visibility,omitempty"`
	Weight        int                 `json:"weight,omitempty"`
	Default       interface{}         `json:"default,omitempty"`
	ValidationMin int                 `json:"validation_minimum,omitempty"`
	ValidationMax int                 `json:"validation_maximum,omitempty"`
	Choices       []string            `json:"choices,omitempty"`
	ChoiceSet     *NestedChoiceSet    `json:"choice_set,omitempty"`
}

// ListCustomFieldsRequest holds query options for listing custom fields
type ListCustomFieldsRequest struct {
	common.ListParams
	ContentType  string
	ContentTypes string
	Type         string
}

// ListCustomFieldsResponse is the paginated response for custom fields
type ListCustomFieldsResponse struct {
	common.ReturnValues
	Results []CustomField `json:"results"`
}

// SetCustomFieldsParams adds the request-specific filters to the URL query
func (p *ListCustomFieldsRequest) SetCustomFieldsParams(values *url.Values) {
	p.SetListParams(values)
	if p.ContentType != "" {
		values.Set("content_type", p.ContentType)
	}
	if p.ContentTypes != "" {
		values.Set("content_types", p.ContentTypes)
	}
	if p.Type != "" {
		values.Set("type", p.Type)
	}
}
