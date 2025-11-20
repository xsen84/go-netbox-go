// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"github.com/go-openapi/strfmt"

	"github.com/xsen84/go-netbox-go/common"
)

// NestedCircuitProvider represents a nested circuit provider reference
type NestedCircuitProvider struct {
	ID          int        `json:"id"`
	URL         strfmt.URI `json:"url,omitempty"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Display     string     `json:"display,omitempty"`
	Description string     `json:"description,omitempty"`
}

// CircuitProvider represents a circuit provider in Netbox
type CircuitProvider struct {
	ID           int             `json:"id"`
	URL          strfmt.URI      `json:"url,omitempty"`
	Name         string          `json:"name"`
	Slug         string          `json:"slug"`
	Display      string          `json:"display,omitempty"`
	Description  string          `json:"description,omitempty"`
	Comments     string          `json:"comments,omitempty"`
	Accounts     []interface{}   `json:"accounts,omitempty"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	CustomFields interface{}     `json:"custom_fields,omitempty"`
	Created      strfmt.DateTime `json:"created,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated,omitempty"`
	CircuitCount int             `json:"circuit_count,omitempty"`
	Asns         []NestedASN     `json:"asns,omitempty"`
}

// ListCircuitProvidersRequest represents the request parameters for listing circuit providers
type ListCircuitProvidersRequest struct {
	common.ListParams
	Slug        string
	Description string
}

// ListCircuitProvidersResponse represents the response from listing circuit providers
type ListCircuitProvidersResponse struct {
	common.ReturnValues
	Results []CircuitProvider `json:"results"`
}

// NestedCircuitType represents a nested circuit type reference
type NestedCircuitType struct {
	ID          int        `json:"id"`
	URL         strfmt.URI `json:"url,omitempty"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Display     string     `json:"display,omitempty"`
	Description string     `json:"description,omitempty"`
}

// CircuitType represents a circuit type in Netbox
type CircuitType struct {
	ID           int             `json:"id"`
	URL          strfmt.URI      `json:"url,omitempty"`
	Name         string          `json:"name"`
	Slug         string          `json:"slug"`
	Display      string          `json:"display,omitempty"`
	Description  string          `json:"description,omitempty"`
	Color        string          `json:"color,omitempty"`
	Tags         []NestedTag     `json:"tags,omitempty"`
	CustomFields interface{}     `json:"custom_fields,omitempty"`
	Created      strfmt.DateTime `json:"created,omitempty"`
	LastUpdated  strfmt.DateTime `json:"last_updated,omitempty"`
	CircuitCount int             `json:"circuit_count,omitempty"`
}

// ListCircuitTypesRequest represents the request parameters for listing circuit types
type ListCircuitTypesRequest struct {
	common.ListParams
	Slug        string
	Description string
	Color       string
}

// ListCircuitTypesResponse represents the response from listing circuit types
type ListCircuitTypesResponse struct {
	common.ReturnValues
	Results []CircuitType `json:"results"`
}
