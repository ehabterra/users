package user

import (
	"fmt"
	"users/gen/roles"
	"users/gen/users"
	storage "users/pkg/db"
)

type Manager struct {
	Db storage.Db
}

func NewManager(db storage.Db) *Manager {
	return &Manager{db}
}

func (m *Manager) List() (res users.StoredUserCollection, err error) {
	err = m.Db.LoadAll(storage.UserBucket, &res)
	if err != nil {
		return nil, err // internal error
	}
	return res, nil
}

func (m *Manager) Show(email string) (res *users.StoredUser, err error) {
	res = &users.StoredUser{}
	err = m.Db.Load(storage.UserBucket, email, res)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, &users.NotFound{
				Message: err.Error(),
				ID:      email,
			}
		}
		return nil, err // internal error
	}
	return res, nil
}

func (m *Manager) CheckRoleExists(role string) (bool, error) {
	res := &roles.StoredRole{}
	err := m.Db.Load(storage.RoleBucket, role, res)
	fmt.Printf("res: %v\n", *res)

	// Check role existence
	if err != nil {
		if err == storage.ErrNotFound {
			return false, &roles.NotFound{
				Message: err.Error(),
				ID:      role,
			}
		}
		return false, err // internal error
	}
	return true, nil

}

func (m *Manager) Add(p *users.User) (err error) {
	if _, err := m.CheckRoleExists(p.Role); err != nil {
		return err
	}

	sb := users.StoredUser{
		Email:     p.Email,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Isactive:  p.Isactive,
		Role:      p.Role,
	}
	if err = m.Db.Save(storage.UserBucket, p.Email, &sb); err != nil {
		return err // internal error
	}
	return nil

}

func (m *Manager) Update(p *users.User) (err error) {
	if _, err := m.CheckRoleExists(p.Role); err != nil {
		return err
	}

	sb := users.StoredUser{
		Email:     p.Email,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Isactive:  p.Isactive,
		Role:      p.Role,
	}
	if err = m.Db.Save(storage.UserBucket, p.Email, &sb); err != nil {
		return err // internal error
	}
	return nil
}

func (m *Manager) Remove(email string) (err error) {
	return m.Db.Delete(storage.UserBucket, email) // internal error if not nil
}

func (m *Manager) Activate(p []string) (err error) {
	for _, email := range p {
		res := &users.StoredUser{}

		err := m.Db.Load(storage.UserBucket, email, res)

		res.Isactive = true

		if err = m.Db.Save(storage.UserBucket, email, res); err != nil {
			return err // internal error
		}
	}
	return nil
}
