// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package extras

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xsen84/go-netbox-go/models"
)

// ListCustomFields retrieves custom field definitions from Netbox
func (c *Client) ListCustomFields(opts models.ListCustomFieldsRequest) (*models.ListCustomFieldsResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"custom-fields/", http.NoBody)
	if err != nil {
		return nil, err
	}
	c.ApplyAuthTokenToHeader(&request.Header)
	setListCustomFieldsParams(request, opts)
	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListCustomFieldsResponse{}
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

func setListCustomFieldsParams(req *http.Request, opts models.ListCustomFieldsRequest) {
	q := req.URL.Query()
	opts.SetCustomFieldsParams(&q)
	req.URL.RawQuery = q.Encode()
}
