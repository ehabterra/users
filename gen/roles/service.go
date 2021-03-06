// Code generated by goa v3.2.3, DO NOT EDIT.
//
// roles service
//
// Command:
// $ goa gen users/design

package roles

import (
	"context"
	rolesviews "users/gen/roles/views"
)

// The roles service performs role data.
type Service interface {
	// List all stored roles
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	List(context.Context, *ListPayload) (res StoredRoleCollection, view string, err error)
	// Show role by name
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Show(context.Context, *ShowPayload) (res *StoredRole, view string, err error)
	// Add new role and return name.
	Add(context.Context, *Role) (res string, err error)
	// Update existing role and return name.
	Update(context.Context, *Role) (res string, err error)
	// Remove role from roles data
	Remove(context.Context, *RemovePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "roles"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list", "show", "add", "update", "remove"}

// ListPayload is the payload type of the roles service list method.
type ListPayload struct {
	// View to render
	View *string
}

// StoredRoleCollection is the result type of the roles service list method.
type StoredRoleCollection []*StoredRole

// ShowPayload is the payload type of the roles service show method.
type ShowPayload struct {
	// Name of role to show
	Name string
	// View to render
	View *string
}

// StoredRole is the result type of the roles service show method.
type StoredRole struct {
	// Name of role
	Name string
	// Description of role
	Description *string
}

// Role is the payload type of the roles service add method.
type Role struct {
	// Name of role
	Name string
	// Description of role
	Description *string
}

// RemovePayload is the payload type of the roles service remove method.
type RemovePayload struct {
	// Name of role to remove
	Name string
}

// NotFound is the type returned when attempting to show or delete a user that
// does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing user
	ID string
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when attempting to show or delete a user that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return e.Message
}

// NewStoredRoleCollection initializes result type StoredRoleCollection from
// viewed result type StoredRoleCollection.
func NewStoredRoleCollection(vres rolesviews.StoredRoleCollection) StoredRoleCollection {
	var res StoredRoleCollection
	switch vres.View {
	case "default", "":
		res = newStoredRoleCollection(vres.Projected)
	case "tiny":
		res = newStoredRoleCollectionTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredRoleCollection initializes viewed result type
// StoredRoleCollection from result type StoredRoleCollection using the given
// view.
func NewViewedStoredRoleCollection(res StoredRoleCollection, view string) rolesviews.StoredRoleCollection {
	var vres rolesviews.StoredRoleCollection
	switch view {
	case "default", "":
		p := newStoredRoleCollectionView(res)
		vres = rolesviews.StoredRoleCollection{Projected: p, View: "default"}
	case "tiny":
		p := newStoredRoleCollectionViewTiny(res)
		vres = rolesviews.StoredRoleCollection{Projected: p, View: "tiny"}
	}
	return vres
}

// NewStoredRole initializes result type StoredRole from viewed result type
// StoredRole.
func NewStoredRole(vres *rolesviews.StoredRole) *StoredRole {
	var res *StoredRole
	switch vres.View {
	case "default", "":
		res = newStoredRole(vres.Projected)
	case "tiny":
		res = newStoredRoleTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredRole initializes viewed result type StoredRole from result
// type StoredRole using the given view.
func NewViewedStoredRole(res *StoredRole, view string) *rolesviews.StoredRole {
	var vres *rolesviews.StoredRole
	switch view {
	case "default", "":
		p := newStoredRoleView(res)
		vres = &rolesviews.StoredRole{Projected: p, View: "default"}
	case "tiny":
		p := newStoredRoleViewTiny(res)
		vres = &rolesviews.StoredRole{Projected: p, View: "tiny"}
	}
	return vres
}

// newStoredRoleCollection converts projected type StoredRoleCollection to
// service type StoredRoleCollection.
func newStoredRoleCollection(vres rolesviews.StoredRoleCollectionView) StoredRoleCollection {
	res := make(StoredRoleCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredRole(n)
	}
	return res
}

// newStoredRoleCollectionTiny converts projected type StoredRoleCollection to
// service type StoredRoleCollection.
func newStoredRoleCollectionTiny(vres rolesviews.StoredRoleCollectionView) StoredRoleCollection {
	res := make(StoredRoleCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredRoleTiny(n)
	}
	return res
}

// newStoredRoleCollectionView projects result type StoredRoleCollection to
// projected type StoredRoleCollectionView using the "default" view.
func newStoredRoleCollectionView(res StoredRoleCollection) rolesviews.StoredRoleCollectionView {
	vres := make(rolesviews.StoredRoleCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredRoleView(n)
	}
	return vres
}

// newStoredRoleCollectionViewTiny projects result type StoredRoleCollection to
// projected type StoredRoleCollectionView using the "tiny" view.
func newStoredRoleCollectionViewTiny(res StoredRoleCollection) rolesviews.StoredRoleCollectionView {
	vres := make(rolesviews.StoredRoleCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredRoleViewTiny(n)
	}
	return vres
}

// newStoredRole converts projected type StoredRole to service type StoredRole.
func newStoredRole(vres *rolesviews.StoredRoleView) *StoredRole {
	res := &StoredRole{
		Description: vres.Description,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newStoredRoleTiny converts projected type StoredRole to service type
// StoredRole.
func newStoredRoleTiny(vres *rolesviews.StoredRoleView) *StoredRole {
	res := &StoredRole{}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newStoredRoleView projects result type StoredRole to projected type
// StoredRoleView using the "default" view.
func newStoredRoleView(res *StoredRole) *rolesviews.StoredRoleView {
	vres := &rolesviews.StoredRoleView{
		Name:        &res.Name,
		Description: res.Description,
	}
	return vres
}

// newStoredRoleViewTiny projects result type StoredRole to projected type
// StoredRoleView using the "tiny" view.
func newStoredRoleViewTiny(res *StoredRole) *rolesviews.StoredRoleView {
	vres := &rolesviews.StoredRoleView{
		Name: &res.Name,
	}
	return vres
}
