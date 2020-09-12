package api

import (
	"context"
	"users/gen/roles"
	storage "users/pkg/db"

	"github.com/dropbox/godropbox/errors"
)

// RoleManager ...
type RoleManager interface {
	List() (res roles.StoredRoleCollection, err error)
	Show(name string) (res *roles.StoredRole, err error)
	Add(p *roles.Role) error
	Update(p *roles.Role) error
	Remove(name string) (err error)
}

// roles service example implementation.
// The example methods log the requests and return zero values.
type Role struct {
	Service RoleManager
}

// NewRole returns the roles service implementation.
func NewRole(manager RoleManager) *Role {
	// Build and return service implementation.
	return &Role{manager}
}

// List all stored roles
func (s *Role) List(_ context.Context, p *roles.ListPayload) (res roles.StoredRoleCollection, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.List()
	if err != nil {
		return nil, view, err
	}
	return res, view, nil
}

// Show role by name
func (s *Role) Show(_ context.Context, p *roles.ShowPayload) (res *roles.StoredRole, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.Show(p.Name)
	if err != nil {
		if errors.IsError(err, storage.ErrNotFound) {
			return nil, "", &roles.NotFound{Message: err.Error(), ID: p.Name}
		}
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Add new role and return name.
func (s *Role) Add(_ context.Context, p *roles.Role) (res string, err error) {
	if err = s.Service.Add(p); err != nil {
		return "", err // internal error
	}
	return p.Name, nil
}

// Update existing role and return name.
func (s *Role) Update(_ context.Context, p *roles.Role) (res string, err error) {
	if err = s.Service.Update(p); err != nil {
		return "", err // internal error
	}
	return p.Name, nil
}

// Remove role from roles data
func (s *Role) Remove(_ context.Context, p *roles.RemovePayload) (err error) {
	return s.Service.Remove(p.Name) // internal error if not nil
}
