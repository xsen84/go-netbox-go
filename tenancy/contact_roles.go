package tenancy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xsen84/go-netbox-go/models"
)

// ListContactRoles returns a list of contact roles from Netbox
func (c *Client) ListContactRoles(opts models.ListContactRolesRequest) (*models.ListContactRolesResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"contact-roles/", http.NoBody)
	if err != nil {
		return nil, err
	}

	c.ApplyAuthTokenToHeader(&request.Header)

	query := request.URL.Query()
	if opts.Name != "" {
		query.Add("name", opts.Name)
	}
	if opts.Slug != "" {
		query.Add("slug", opts.Slug)
	}
	if opts.Limit > 0 {
		query.Add("limit", fmt.Sprintf("%d", opts.Limit))
	}
	if opts.Offset > 0 {
		query.Add("offset", fmt.Sprintf("%d", opts.Offset))
	}
	request.URL.RawQuery = query.Encode()

	response, err := c.HTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}

	resObj := models.ListContactRolesResponse{}
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
