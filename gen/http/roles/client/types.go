// Code generated by goa v3.2.3, DO NOT EDIT.
//
// roles HTTP client types
//
// Command:
// $ goa gen users/design

package client

import (
	roles "users/gen/roles"
	rolesviews "users/gen/roles/views"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "roles" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of role
	Name string `form:"name" json:"name" xml:"name"`
	// Description of role
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// UpdateRequestBody is the type of the "roles" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Description of role
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// ListResponseBody is the type of the "roles" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredRoleResponse

// ShowResponseBody is the type of the "roles" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// Name of role
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of role
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "roles" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing user
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredRoleResponse is used to define fields on response body types.
type StoredRoleResponse struct {
	// Name of role
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of role
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "roles" service.
func NewAddRequestBody(p *roles.Role) *AddRequestBody {
	body := &AddRequestBody{
		Name:        p.Name,
		Description: p.Description,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "roles" service.
func NewUpdateRequestBody(p *roles.Role) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Description: p.Description,
	}
	return body
}

// NewListStoredRoleCollectionOK builds a "roles" service "list" endpoint
// result from a HTTP "OK" response.
func NewListStoredRoleCollectionOK(body ListResponseBody) rolesviews.StoredRoleCollectionView {
	v := make([]*rolesviews.StoredRoleView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredRoleResponseToRolesviewsStoredRoleView(val)
	}
	return v
}

// NewShowStoredRoleOK builds a "roles" service "show" endpoint result from a
// HTTP "OK" response.
func NewShowStoredRoleOK(body *ShowResponseBody) *rolesviews.StoredRoleView {
	v := &rolesviews.StoredRoleView{
		Name:        body.Name,
		Description: body.Description,
	}

	return v
}

// NewShowNotFound builds a roles service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *roles.NotFound {
	v := &roles.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateStoredRoleResponse runs the validations defined on StoredRoleResponse
func ValidateStoredRoleResponse(body *StoredRoleResponse) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.name", *body.Name, "[a-z]+[a-z0-9]*"))
	}
	return
}
