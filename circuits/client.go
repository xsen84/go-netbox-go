// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package circuits

import (
	"github.com/xsen84/go-netbox-go/common"
	"github.com/xsen84/go-netbox-go/models"
)

const basePath = "/api/circuits/"

type NetboxAPI interface {
	common.HTTPConnectable

	// circuit providers
	ListCircuitProviders(opts models.ListCircuitProvidersRequest) (*models.ListCircuitProvidersResponse, error)
	GetCircuitProvider(id int) (*models.CircuitProvider, error)

	// circuit types
	ListCircuitTypes(opts models.ListCircuitTypesRequest) (*models.ListCircuitTypesResponse, error)
	GetCircuitType(id int) (*models.CircuitType, error)
}

type Client struct {
	common.Client
}

// Deprecated: Use NewClient() function instead.
func New(baseURL, authToken string, insecureSkipVerify bool) (*Client, error) {
	client := &Client{}
	return client, common.Initialize(client, baseURL, authToken, insecureSkipVerify)
}

func NewClient(baseURL, authToken string, insecureSkipVerify bool) (NetboxAPI, error) {
	client := &Client{}
	return client, common.Initialize(client, baseURL, authToken, insecureSkipVerify)
}
