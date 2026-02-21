package tenancy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xsen84/go-netbox-go/models"
)

// ListContactAssignments returns a list of contact assignments from Netbox
func (c *Client) ListContactAssignments(opts models.ListContactAssignmentsRequest) (*models.ListContactAssignmentsResponse, error) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, c.BaseURL().String()+basePath+"contact-assignments/", http.NoBody)
	if err != nil {
		return nil, err
	}

	c.ApplyAuthTokenToHeader(&request.Header)

	query := request.URL.Query()
	if opts.ContactId != 0 {
		query.Add("contact_id", fmt.Sprintf("%d", opts.ContactId))
	}
	if opts.RoleId != 0 {
		query.Add("role_id", fmt.Sprintf("%d", opts.RoleId))
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

	resObj := models.ListContactAssignmentsResponse{}
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
