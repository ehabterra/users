// Code generated by goa v3.2.3, DO NOT EDIT.
//
// roles gRPC client encoders and decoders
//
// Command:
// $ goa gen users/design

package client

import (
	"context"
	rolespb "users/gen/grpc/roles/pb"
	roles "users/gen/roles"
	rolesviews "users/gen/roles/views"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildListFunc builds the remote method to invoke for "roles" service "list"
// endpoint.
func BuildListFunc(grpccli rolespb.RolesClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.List(ctx, reqpb.(*rolespb.ListRequest), opts...)
		}
		return grpccli.List(ctx, &rolespb.ListRequest{}, opts...)
	}
}

// EncodeListRequest encodes requests sent to roles list endpoint.
func EncodeListRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*roles.ListPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "list", "*roles.ListPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewListRequest(), nil
}

// DecodeListResponse decodes responses from the roles list endpoint.
func DecodeListResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*rolespb.StoredRoleCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "list", "*rolespb.StoredRoleCollection", v)
	}
	res := NewListResult(message)
	vres := rolesviews.StoredRoleCollection{Projected: res, View: view}
	if err := rolesviews.ValidateStoredRoleCollection(vres); err != nil {
		return nil, err
	}
	return roles.NewStoredRoleCollection(vres), nil
}

// BuildShowFunc builds the remote method to invoke for "roles" service "show"
// endpoint.
func BuildShowFunc(grpccli rolespb.RolesClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Show(ctx, reqpb.(*rolespb.ShowRequest), opts...)
		}
		return grpccli.Show(ctx, &rolespb.ShowRequest{}, opts...)
	}
}

// EncodeShowRequest encodes requests sent to roles show endpoint.
func EncodeShowRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*roles.ShowPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "show", "*roles.ShowPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	return NewShowRequest(payload), nil
}

// DecodeShowResponse decodes responses from the roles show endpoint.
func DecodeShowResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*rolespb.ShowResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "show", "*rolespb.ShowResponse", v)
	}
	res := NewShowResult(message)
	vres := &rolesviews.StoredRole{Projected: res, View: view}
	if err := rolesviews.ValidateStoredRole(vres); err != nil {
		return nil, err
	}
	return roles.NewStoredRole(vres), nil
}

// BuildAddFunc builds the remote method to invoke for "roles" service "add"
// endpoint.
func BuildAddFunc(grpccli rolespb.RolesClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*rolespb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &rolespb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to roles add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*roles.Role)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "add", "*roles.Role", v)
	}
	return NewAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the roles add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*rolespb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "add", "*rolespb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}

// BuildUpdateFunc builds the remote method to invoke for "roles" service
// "update" endpoint.
func BuildUpdateFunc(grpccli rolespb.RolesClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Update(ctx, reqpb.(*rolespb.UpdateRequest), opts...)
		}
		return grpccli.Update(ctx, &rolespb.UpdateRequest{}, opts...)
	}
}

// EncodeUpdateRequest encodes requests sent to roles update endpoint.
func EncodeUpdateRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*roles.Role)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "update", "*roles.Role", v)
	}
	return NewUpdateRequest(payload), nil
}

// DecodeUpdateResponse decodes responses from the roles update endpoint.
func DecodeUpdateResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*rolespb.UpdateResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "update", "*rolespb.UpdateResponse", v)
	}
	res := NewUpdateResult(message)
	return res, nil
}

// BuildRemoveFunc builds the remote method to invoke for "roles" service
// "remove" endpoint.
func BuildRemoveFunc(grpccli rolespb.RolesClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Remove(ctx, reqpb.(*rolespb.RemoveRequest), opts...)
		}
		return grpccli.Remove(ctx, &rolespb.RemoveRequest{}, opts...)
	}
}

// EncodeRemoveRequest encodes requests sent to roles remove endpoint.
func EncodeRemoveRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*roles.RemovePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("roles", "remove", "*roles.RemovePayload", v)
	}
	return NewRemoveRequest(payload), nil
}
