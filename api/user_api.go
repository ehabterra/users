package api

import (
	"context"
	"users/gen/users"
)

type UserManager interface {
	List() (res users.StoredUserCollection, err error)
	Show(email string) (res *users.StoredUser, err error)
	Add(p *users.User) (err error)
	Update(p *users.User) (err error)
	Remove(email string) (err error)
	Activate(p []string) (err error)
}

// users service example implementation.
// The example methods log the requests and return zero values.
type UserAPI struct {
	Service UserManager
}

// NewUserAPI returns the users service implementation.
func NewUserAPI(manager UserManager) *UserAPI {
	// Build and return service implementation.
	return &UserAPI{manager}
}

// List all stored users
func (s *UserAPI) List(_ context.Context, p *users.ListPayload) (res users.StoredUserCollection, view string, err error) {
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

// Show user by Email
func (s *UserAPI) Show(_ context.Context, p *users.ShowPayload) (res *users.StoredUser, view string, err error) {
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.Show(p.Email)
	if err != nil {
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Add new user and return email.
func (s *UserAPI) Add(_ context.Context, p *users.User) (res string, err error) {
	if err = s.Service.Add(p); err != nil {
		return "", err
	}
	return res, nil
}

// Update existing user and return email.
func (s *UserAPI) Update(_ context.Context, p *users.User) (res string, err error) {
	if err = s.Service.Update(p); err != nil {
		return "", err // internal error
	}
	return res, nil
}

// Remove user from users data
func (s *UserAPI) Remove(_ context.Context, p *users.RemovePayload) (err error) {
	return s.Service.Remove(p.Email)
}

// Activate users by emails
func (s *UserAPI) Activate(_ context.Context, p []string) (err error) {
	if err = s.Service.Activate(p); err != nil {
		return err
	}
	return nil
}
