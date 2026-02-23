package models

// Contact defines the structure for a Tenancy Contact in Netbox
type Contact struct {
	Id           int                 `json:"id"`
	Url          string              `json:"url"`
	Display      string              `json:"display"`
	Group        *NestedContactGroup `json:"group"`
	Name         string              `json:"name"`
	Title        string              `json:"title"`
	Phone        string              `json:"phone"`
	Email        string              `json:"email"`
	Address      string              `json:"address"`
	Link         string              `json:"link"`
	Description  string              `json:"description"`
	Comments     string              `json:"comments"`
	Tags         []NestedTag         `json:"tags"`
	CustomFields interface{}         `json:"custom_fields"`
	Created      string              `json:"created"`
	LastUpdated  string              `json:"last_updated"`
}

// ListContactsResponse defines the API response for listing contacts
type ListContactsResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Contact `json:"results"`
}

// ListContactsRequest defines filters that can be applied to list contacts
type ListContactsRequest struct {
	Name   string
	Group  string
	Limit  int
	Offset int
}

// NestedContactGroup defines a nested struct for contact group references
type NestedContactGroup struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
}

// ContactRole defines the structure for a Tenancy Contact Role in Netbox
type ContactRole struct {
	Id           int         `json:"id"`
	Url          string      `json:"url"`
	Display      string      `json:"display"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	Description  string      `json:"description"`
	CustomFields interface{} `json:"custom_fields"`
	Created      string      `json:"created"`
	LastUpdated  string      `json:"last_updated"`
}

// ListContactRolesResponse defines the API response for listing contact roles
type ListContactRolesResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []ContactRole `json:"results"`
}

// ListContactRolesRequest defines filters that can be applied to list contact roles
type ListContactRolesRequest struct {
	Name   string
	Slug   string
	Limit  int
	Offset int
}

// ContactPriority defines the label/value structure for Contact Priority
type ContactPriority struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// ContactAssignment defines the structure for a Contact assigned to a given Netbox object
type ContactAssignment struct {
	Id          int              `json:"id"`
	Url         string           `json:"url"`
	Display     string           `json:"display"`
	Object      interface{}      `json:"object"`
	Contact     Contact          `json:"contact"`
	Role        ContactRole      `json:"role"`
	Priority    *ContactPriority `json:"priority"`
	Created     string           `json:"created"`
	LastUpdated string           `json:"last_updated"`
}

// ListContactAssignmentsResponse defines the response for listing contact assignments
type ListContactAssignmentsResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []ContactAssignment `json:"results"`
}

// ListContactAssignmentsRequest defines filters that can be applied to list contact assignments
type ListContactAssignmentsRequest struct {
	ContactId   int
	RoleId      int
	ObjectId    int
	ContentType string
	Limit       int
	Offset      int
}
