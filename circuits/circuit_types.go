// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package circuits

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/xsen84/go-netbox-go/models"
)

// ListCircuitTypes retrieves a list of circuit types from Netbox
func (c *Client) ListCircuitTypes(opts models.ListCircuitTypesRequest) (*models.ListCircuitTypesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"circuit-types/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListCircuitTypesParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListCircuitTypesResponse{}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}

// setListCircuitTypesParams sets the query parameters for listing circuit types
func setListCircuitTypesParams(req *http.Request, opts models.ListCircuitTypesRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Slug != "" {
		q.Set("slug", opts.Slug)
	}
	if opts.Description != "" {
		q.Set("description", opts.Description)
	}
	if opts.Color != "" {
		q.Set("color", opts.Color)
	}
	req.URL.RawQuery = q.Encode()
}

// GetCircuitType retrieves a single circuit type by ID
func (c *Client) GetCircuitType(id int) (*models.CircuitType, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"circuit-types/"+strconv.Itoa(id)+"/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.CircuitType{}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}
