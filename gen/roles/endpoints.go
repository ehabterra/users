// Code generated by goa v3.2.3, DO NOT EDIT.
//
// roles endpoints
//
// Command:
// $ goa gen users/design

package roles

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "roles" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Show   goa.Endpoint
	Add    goa.Endpoint
	Update goa.Endpoint
	Remove goa.Endpoint
}

// NewEndpoints wraps the methods of the "roles" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:   NewListEndpoint(s),
		Show:   NewShowEndpoint(s),
		Add:    NewAddEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Remove: NewRemoveEndpoint(s),
	}
}

// Use applies the given middleware to all the "roles" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Add = m(e.Add)
	e.Update = m(e.Update)
	e.Remove = m(e.Remove)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "roles".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		res, view, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredRoleCollection(res, view)
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "roles".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredRole(res, view)
		return vres, nil
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "roles".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Role)
		return s.Add(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "roles".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Role)
		return s.Update(ctx, p)
	}
}

// NewRemoveEndpoint returns an endpoint function that calls the method
// "remove" of service "roles".
func NewRemoveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemovePayload)
		return nil, s.Remove(ctx, p)
	}
}
