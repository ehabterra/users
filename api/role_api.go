package api

import (
	"context"
	"users/gen/roles"
)

type RoleManager interface {
	List() (res roles.StoredRoleCollection, err error)
	Show(name string) (res *roles.StoredRole, err error)
	Add(p *roles.Role) error
	Update(p *roles.Role) error
	Remove(name string) (err error)
}

// roles service example implementation.
// The example methods log the requests and return zero values.
type RoleAPI struct {
	Service RoleManager
}

// NewRoleAPI returns the roles service implementation.
func NewRoleAPI(manager RoleManager) *RoleAPI {
	// Build and return service implementation.
	return &RoleAPI{manager}
}

// List all stored roles
func (s *RoleAPI) List(_ context.Context, p *roles.ListPayload) (res roles.StoredRoleCollection, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.List()
	if err != nil {
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Show role by name
func (s *RoleAPI) Show(_ context.Context, p *roles.ShowPayload) (res *roles.StoredRole, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.Show(p.Name)
	if err != nil {
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Add new role and return name.
func (s *RoleAPI) Add(_ context.Context, p *roles.Role) (res string, err error) {
	if err = s.Service.Add(p); err != nil {
		return "", err // internal error
	}
	return p.Name, nil
}

// Update existing role and return name.
func (s *RoleAPI) Update(_ context.Context, p *roles.Role) (res string, err error) {
	if err = s.Service.Update(p); err != nil {
		return "", err // internal error
	}
	return p.Name, nil
}

// Remove role from roles data
func (s *RoleAPI) Remove(_ context.Context, p *roles.RemovePayload) (err error) {
	return s.Service.Remove(p.Name) // internal error if not nil
}
