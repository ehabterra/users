package role

import (
	"users/gen/roles"
	storage "users/pkg/db"
)

type Manager struct {
	Db storage.Db
}

// NewManager returns the roles service implementation.
func NewManager(bolt storage.Db) *Manager {
	// Build and return service implementation.
	return &Manager{bolt}
}

func (m *Manager) List() (res roles.StoredRoleCollection, err error) {
	err = m.Db.LoadAll(storage.RoleBucket, &res)
	if err != nil {
		return nil, err // internal error
	}
	return res, nil
}

// Show role by name
func (m *Manager) Show(name string) (res *roles.StoredRole, err error) {
	res = &roles.StoredRole{}
	err = m.Db.Load(storage.RoleBucket, name, res)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, &roles.NotFound{
				Message: err.Error(),
				ID:      name,
			}
		}
		return nil, err // internal error
	}
	return res, nil
}

// Add new role and return name.
func (m *Manager) Add(p *roles.Role) error {
	sb := roles.StoredRole{
		Name:        p.Name,
		Description: p.Description,
	}
	if err := m.Db.Save(storage.RoleBucket, p.Name, &sb); err != nil {
		return err // internal error
	}
	return nil
}

func (m *Manager) Update(p *roles.Role) error {
	sb := roles.StoredRole{
		Name:        p.Name,
		Description: p.Description,
	}
	if err := m.Db.Save(storage.RoleBucket, p.Name, &sb); err != nil {
		return err // internal error
	}
	return nil
}

// Remove role from roles data
func (m *Manager) Remove(name string) (err error) {
	return m.Db.Delete(storage.RoleBucket, name) // internal error if not nil
}
