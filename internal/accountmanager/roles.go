package accountmanager

import (
	"context"
	"users/gen/roles"
	storage "users/pkg/db"
)

// roles service example implementation.
// The example methods log the requests and return zero values.
type Rolessrvc struct {
	Db storage.Db
}

// NewRoles returns the roles service implementation.
func NewRoles(bolt storage.Db) roles.Service {
	// Build and return service implementation.
	return &Rolessrvc{bolt}
}

// List all stored roles
func (s *Rolessrvc) List(ctx context.Context, p *roles.ListPayload) (res roles.StoredRoleCollection, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	x, errc := s.Db.LoadAll(storage.RoleBucket)
	res = x.(roles.StoredRoleCollection)
	if errc != nil {
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Show role by name
func (s *Rolessrvc) Show(ctx context.Context, p *roles.ShowPayload) (res *roles.StoredRole, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	x, err := s.Db.Load(storage.RoleBucket, p.Name)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, view, &roles.NotFound{
				Message: err.Error(),
				ID:      p.Name,
			}
		}
		return nil, view, err // internal error
	}
	res = x.(*roles.StoredRole)

	return res, view, nil
}

// Add new role and return name.
func (s *Rolessrvc) Add(ctx context.Context, p *roles.Role) (res string, err error) {
	res, err = s.Db.NewID(storage.RoleBucket)
	if err != nil {
		return "", err // internal error
	}
	sb := roles.StoredRole{
		//ID:          res,
		Name:        p.Name,
		Description: p.Description,
	}
	if err = s.Db.Save(storage.RoleBucket, p.Name, &sb); err != nil {
		return "", err // internal error
	}
	return p.Name, nil
}

// Update existing role and return name.
func (s *Rolessrvc) Update(ctx context.Context, p *roles.Role) (res string, err error) {
	sb := roles.StoredRole{
		Name:        p.Name,
		Description: p.Description,
	}
	if err = s.Db.Save(storage.RoleBucket, p.Name, &sb); err != nil {
		return "", err // internal error
	}
	return p.Name, nil
}

// Remove role from roles data
func (s *Rolessrvc) Remove(ctx context.Context, p *roles.RemovePayload) (err error) {
	return s.Db.Delete(storage.RoleBucket, p.Name) // internal error if not nil
}
