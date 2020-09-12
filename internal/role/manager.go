package role

import (
	"fmt"
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

// List roles
func (m *Manager) List() (res roles.StoredRoleCollection, err error) {
	err = m.Db.LoadAll(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Show role by name
func (m *Manager) Show(name string) (res *roles.StoredRole, err error) {
	res = &roles.StoredRole{}
	err = m.Db.Load(name, res)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, &roles.NotFound{
				Message: err.Error(),
				ID:      name,
			}
		}
		return nil, err
	}
	return res, nil
}

// Add new role and return name.
func (m *Manager) Add(p *roles.Role) error {
	sb := roles.StoredRole{
		Name:        p.Name,
		Description: p.Description,
	}
	if err := m.Db.Save(p.Name, &sb); err != nil {
		return err
	}
	return nil
}

func (m *Manager) Update(p *roles.Role) error {
	sb := roles.StoredRole{
		Name:        p.Name,
		Description: p.Description,
	}
	if err := m.Db.Save(p.Name, &sb); err != nil {
		return err
	}
	return nil
}

// Remove role from roles data
func (m *Manager) Remove(name string) (err error) {
	return m.Db.Delete(name)
}

func (m *Manager) CheckRoleExists(role string) (bool, error) {
	res := &roles.StoredRole{}
	err := m.Db.Load(role, res)
	fmt.Printf("res: %v\n", *res)

	// Check role existence
	if err != nil {
		if err == storage.ErrNotFound {
			return false, &roles.NotFound{
				Message: err.Error(),
				ID:      role,
			}
		}
		return false, err
	}
	return true, nil

}
