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

// ListCircuitProviders retrieves a list of circuit providers from Netbox
func (c *Client) ListCircuitProviders(opts models.ListCircuitProvidersRequest) (*models.ListCircuitProvidersResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"providers/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListCircuitProvidersParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListCircuitProvidersResponse{}
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

// setListCircuitProvidersParams sets the query parameters for listing circuit providers
func setListCircuitProvidersParams(req *http.Request, opts models.ListCircuitProvidersRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	if opts.Slug != "" {
		q.Set("slug", opts.Slug)
	}
	if opts.Description != "" {
		q.Set("description", opts.Description)
	}
	req.URL.RawQuery = q.Encode()
}

// GetCircuitProvider retrieves a single circuit provider by ID
func (c *Client) GetCircuitProvider(id int) (*models.CircuitProvider, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"providers/"+strconv.Itoa(id)+"/", http.NoBody)
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
	resObj := models.CircuitProvider{}
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
