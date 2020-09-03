package accountmanager

import (
	"context"
	"users/gen/roles"
	"users/gen/users"
	"users/pkg/db"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	db     storage.Db
}

// NewUsers returns the users service implementation.
func NewUsers(bolt storage.Db) users.Service {
	// Build and return service implementation.
	return &userssrvc{bolt }
}

// List all stored users
func (s *userssrvc) List(ctx context.Context, p *users.ListPayload) (res users.StoredUserCollection, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}
	if err = s.db.LoadAll(storage.UserBucket, &res); err != nil {
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Show user by Email
func (s *userssrvc) Show(ctx context.Context, p *users.ShowPayload) (res *users.StoredUser, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}
	if err = s.db.Load(storage.UserBucket, p.Email, &res); err != nil {
		if err == storage.ErrNotFound {
			return nil, view, &users.NotFound{
				Message: err.Error(),
				ID:      p.Email,
			}
		}
		return nil, view, err // internal error
	}
	return res, view, nil
}

func (s *userssrvc) checkRoleExists(role string) (bool, error) {
	var res *roles.StoredRole

	// Check role existence
	if err := s.db.Load(storage.RoleBucket, role, &res); err != nil {
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

// Add new user and return email.
func (s *userssrvc) Add(ctx context.Context, p *users.User) (res string, err error) {
	if _, err := s.checkRoleExists(p.Role); err != nil {
		return "", err
	}

	res, err = s.db.NewID(storage.UserBucket)
	if err != nil {
		return "", err // internal error
	}
	sb := users.StoredUser{
		//ID:          res,
		Email:     p.Email,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Isactive:  p.Isactive,
		Role:      p.Role,
	}
	if err = s.db.Save(storage.UserBucket, p.Email, &sb); err != nil {
		return "", err // internal error
	}
	return res, nil
}

// Update existing user and return email.
func (s *userssrvc) Update(ctx context.Context, p *users.User) (res string, err error) {
	if _, err := s.checkRoleExists(p.Role); err != nil {
		return "", err
	}

	sb := users.StoredUser{
		Email:     p.Email,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Isactive:  p.Isactive,
		Role:      p.Role,
	}
	if err = s.db.Save(storage.UserBucket, p.Email, &sb); err != nil {
		return "", err // internal error
	}
	return res, nil
}

// Remove user from users data
func (s *userssrvc) Remove(ctx context.Context, p *users.RemovePayload) (err error) {
	return s.db.Delete(storage.UserBucket, p.Email) // internal error if not nil
}

// Activate users by emails
func (s *userssrvc) Activate(ctx context.Context, p []string) (err error) {
	for _, email := range p {
		res := &users.StoredUser{}

		s.db.Load(storage.UserBucket, email, &res)

		res.Isactive = true

		if err = s.db.Save(storage.UserBucket, email, &res); err != nil {
			return err // internal error
		}
	}
	return nil
}
