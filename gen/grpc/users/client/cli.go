// Code generated by goa v3.2.3, DO NOT EDIT.
//
// users gRPC client CLI support package
//
// Command:
// $ goa gen users/design

package client

import (
	"encoding/json"
	"fmt"
	userspb "users/gen/grpc/users/pb"
	users "users/gen/users"

	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the users list endpoint from CLI
// flags.
func BuildListPayload(usersListView string) (*users.ListPayload, error) {
	var err error
	var view *string
	{
		if usersListView != "" {
			view = &usersListView
			if view != nil {
				if !(*view == "default" || *view == "tiny") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &users.ListPayload{}
	v.View = view

	return v, nil
}

// BuildShowPayload builds the payload for the users show endpoint from CLI
// flags.
func BuildShowPayload(usersShowMessage string, usersShowView string) (*users.ShowPayload, error) {
	var err error
	var message userspb.ShowRequest
	{
		if usersShowMessage != "" {
			err = json.Unmarshal([]byte(usersShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Illo omnis ea sit nam.\"\n   }'")
			}
		}
	}
	var view *string
	{
		if usersShowView != "" {
			view = &usersShowView
			if view != nil {
				if !(*view == "default" || *view == "tiny") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &users.ShowPayload{
		Email: message.Email,
	}
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the users add endpoint from CLI flags.
func BuildAddPayload(usersAddMessage string) (*users.User, error) {
	var err error
	var message userspb.AddRequest
	{
		if usersAddMessage != "" {
			err = json.Unmarshal([]byte(usersAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"ehabterra@hotmail.com\",\n      \"firstname\": \"Ehab\",\n      \"isactive\": true,\n      \"lastname\": \"Terra\",\n      \"role\": \"admin\"\n   }'")
			}
		}
	}
	v := &users.User{
		Email:     message.Email,
		Firstname: message.Firstname,
		Lastname:  message.Lastname,
		Role:      message.Role,
		Isactive:  message.Isactive,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the users update endpoint from CLI
// flags.
func BuildUpdatePayload(usersUpdateMessage string) (*users.User, error) {
	var err error
	var message userspb.UpdateRequest
	{
		if usersUpdateMessage != "" {
			err = json.Unmarshal([]byte(usersUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"ehabterra@hotmail.com\",\n      \"firstname\": \"Ehab\",\n      \"isactive\": false,\n      \"lastname\": \"Terra\",\n      \"role\": \"admin\"\n   }'")
			}
		}
	}
	v := &users.User{
		Email:     message.Email,
		Firstname: message.Firstname,
		Lastname:  message.Lastname,
		Role:      message.Role,
		Isactive:  message.Isactive,
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the users remove endpoint from CLI
// flags.
func BuildRemovePayload(usersRemoveMessage string) (*users.RemovePayload, error) {
	var err error
	var message userspb.RemoveRequest
	{
		if usersRemoveMessage != "" {
			err = json.Unmarshal([]byte(usersRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Sapiente consequuntur modi nisi.\"\n   }'")
			}
		}
	}
	v := &users.RemovePayload{
		Email: message.Email,
	}

	return v, nil
}

// BuildActivatePayload builds the payload for the users activate endpoint from
// CLI flags.
func BuildActivatePayload(usersActivateMessage string) ([]string, error) {
	var err error
	var message userspb.ActivateRequest
	{
		if usersActivateMessage != "" {
			err = json.Unmarshal([]byte(usersActivateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"field\": [\n         \"Possimus veniam iure.\",\n         \"Autem autem.\"\n      ]\n   }'")
			}
		}
	}
	v := make([]string, len(message.Field))
	for i, val := range message.Field {
		v[i] = val
	}
	return v, nil
}
