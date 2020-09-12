package api

import (
	"context"
	"users/gen/users"
	storage "users/pkg/db"

	"github.com/dropbox/godropbox/errors"
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
type User struct {
	Service UserManager
}

// NewUser returns the users service implementation.
func NewUser(manager UserManager) *User {
	// Build and return service implementation.
	return &User{manager}
}

// List all stored users
func (s *User) List(_ context.Context, p *users.ListPayload) (res users.StoredUserCollection, view string, err error) {
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
func (s *User) Show(ctx context.Context, p *users.ShowPayload) (res *users.StoredUser, view string, err error) {

	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res, err = s.Service.Show(p.Email)
	if err != nil {
		if errors.IsError(err, storage.ErrNotFound) {
			return nil, "", &users.NotFound{Message: err.Error(), ID: p.Email}
		}
		return nil, view, err // internal error
	}
	return res, view, nil
}

// Add new user and return email.
func (s *User) Add(_ context.Context, p *users.User) (res string, err error) {
	if err = s.Service.Add(p); err != nil {
		return "", err
	}
	return res, nil
}

// Update existing user and return email.
func (s *User) Update(_ context.Context, p *users.User) (res string, err error) {
	if err = s.Service.Update(p); err != nil {
		return "", err // internal error
	}
	return res, nil
}

// Remove user from users data
func (s *User) Remove(_ context.Context, p *users.RemovePayload) (err error) {
	return s.Service.Remove(p.Email)
}

// Activate users by emails
func (s *User) Activate(_ context.Context, p []string) (err error) {
	if err = s.Service.Activate(p); err != nil {
		return err
	}
	return nil
}
