package usersapi

import (
	"context"
	"log"
	roles "users/gen/roles"
	"users/pkg/db"
)

// roles service example implementation.
// The example methods log the requests and return zero values.
type rolessrvc struct {
	db     *storage.Db
	logger *log.Logger
}

// NewRoles returns the roles service implementation.
func NewRoles(bolt *storage.Db, logger *log.Logger) roles.Service {
	// Build and return service implementation.
	return &rolessrvc{bolt, logger}
}

// List all stored roles
func (s *rolessrvc) List(ctx context.Context, p *roles.ListPayload) (res roles.StoredRoleCollection, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}
	if err = s.db.LoadAll(storage.RoleBucket, &res); err != nil {
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Show role by name
func (s *rolessrvc) Show(ctx context.Context, p *roles.ShowPayload) (res *roles.StoredRole, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}
	if err = s.db.Load(storage.RoleBucket, p.Name, &res); err != nil {
		if err == storage.ErrNotFound {
			return nil, view, &roles.NotFound{
				Message: err.Error(),
				ID:      p.Name,
			}
		}
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Add new role and return name.
func (s *rolessrvc) Add(ctx context.Context, p *roles.Role) (res string, err error) {
	res, err = s.db.NewID(storage.RoleBucket)
	if err != nil {
		return "", err // internal error
	}
	sb := roles.StoredRole{
		//ID:          res,
		Name:        p.Name,
		Description: p.Description,
	}
	if err = s.db.Save(storage.RoleBucket, p.Name, &sb); err != nil {
		return "", err // internal error
	}
	return res, nil
}

// Update existing role and return name.
func (s *rolessrvc) Update(ctx context.Context, p *roles.Role) (res string, err error) {
	sb := roles.StoredRole{
		Name:        p.Name,
		Description: p.Description,
	}
	if err = s.db.Save(storage.RoleBucket, p.Name, &sb); err != nil {
		return "", err // internal error
	}
	return res, nil
}

// Remove role from roles data
func (s *rolessrvc) Remove(ctx context.Context, p *roles.RemovePayload) (err error) {
	return s.db.Delete(storage.RoleBucket, p.Name) // internal error if not nil
}
