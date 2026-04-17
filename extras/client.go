// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package extras

import (
	"github.com/xsen84/go-netbox-go/common"
	"github.com/xsen84/go-netbox-go/models"
)

const basePath = "/api/extras/"

type NetboxAPI interface {
	common.HTTPConnectable

	// tags
	ListTags(opts models.ListTagsRequest) (*models.ListTagsResponse, error)

	// custom fields
	ListCustomFields(opts models.ListCustomFieldsRequest) (*models.ListCustomFieldsResponse, error)

	// custom field choice sets
	ListCustomFieldChoiceSets(opts models.ListCustomFieldChoiceSetsRequest) (*models.ListCustomFieldChoiceSetsResponse, error)
	GetCustomFieldChoiceSet(id int) (*models.CustomFieldChoiceSet, error)
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
