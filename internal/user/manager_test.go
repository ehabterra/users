package user

import (
	"fmt"
	"reflect"
	"testing"
	"users/gen/roles"
	"users/gen/users"
	"users/internal/role"
	"users/mocks"
	storage "users/pkg/db"

	"github.com/stretchr/testify/mock"
)

func TestManager_Activate(t *testing.T) {
	userMock := &mocks.Db{}
	userData := []*users.User{
		{
			Email:     "ehab@test.com",
			Firstname: "Ehab",
			Lastname:  "Terra",
			Role:      "admin",
			Isactive:  false,
		},
		{
			Email:     "khalifa@test.com",
			Firstname: "Khalifa",
			Lastname:  "Hassan",
			Role:      "admin",
			Isactive:  false,
		},
	}
	var sp []*users.StoredUser

	for _, user := range userData {
		sp = append(sp, &users.StoredUser{
			Email:     user.Email,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Isactive:  user.Isactive,
			Role:      user.Role,
		})
	}

	userMock.On("Save", mock.AnythingOfType("string"), mock.AnythingOfType("*users.StoredUser")).Return(nil)
	userMock.On("Load", mock.AnythingOfType("string"), &users.StoredUser{}).Return(func(email string, res interface{}) error {
		data := res.(*users.StoredUser)
		for _, user := range userData {
			if user.Email == email {
				data.Email = user.Email
				data.Firstname = user.Firstname
				data.Lastname = user.Lastname
				data.Isactive = user.Isactive
				data.Role = user.Role
				break
			}
		}
		return nil
	})

	type fields struct {
		Db storage.Db
	}
	type args struct {
		p []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Activate",
			fields{userMock},
			args{[]string{"ehab@test.com", "khalifa@test.com"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			if err := m.Activate(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Activate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_Add(t *testing.T) {
	userMock := &mocks.Db{}
	roleManager := role.NewManager(userMock)

	user := &users.User{
		Email:     "ehab@test.com",
		Firstname: "Ehab",
		Lastname:  "Terra",
		Role:      "admin",
		Isactive:  true,
	}
	sp := &users.StoredUser{
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Isactive:  user.Isactive,
		Role:      user.Role,
	}
	wantDes := "Administrator"
	testRole := &roles.StoredRole{Name: user.Role, Description: &wantDes}

	userMock.On("Save", user.Email, sp).Return(nil)
	userMock.On("Load", user.Role, &roles.StoredRole{}).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*roles.StoredRole)
		fmt.Print(arg)
		arg.Name = user.Role
		arg.Description = testRole.Description
	})

	type fields struct {
		Db   storage.Db
		role *role.Manager
	}
	type args struct {
		p *users.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Add",
			fields{userMock, roleManager},
			args{user},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db:   tt.fields.Db,
				role: roleManager,
			}
			if err := m.Add(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_List(t *testing.T) {
	userMock := &mocks.Db{}

	userData := users.StoredUserCollection{
		{
			Email:     "ehab@test.com",
			Firstname: "Ehab",
			Lastname:  "Terra",
			Role:      "admin",
			Isactive:  false,
		},
		{
			Email:     "khalifa@test.com",
			Firstname: "Khalifa",
			Lastname:  "Hassan",
			Role:      "admin",
			Isactive:  false,
		},
	}
	var res users.StoredUserCollection

	userMock.On("LoadAll", &res).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*users.StoredUserCollection)
		*arg = append(*arg, userData...)
		fmt.Printf("value %v, type %T \n", arg, arg)
	})

	type fields struct {
		Db storage.Db
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes users.StoredUserCollection
		wantErr bool
	}{
		{
			"List",
			fields{userMock},
			userData,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			gotRes, err := m.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("List() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestManager_Remove(t *testing.T) {
	userMock := mocks.Db{}
	email := "ehab@test.com"
	userMock.On("Delete", email).Return(nil)

	type fields struct {
		Db storage.Db
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Remove",
			fields{&userMock},
			args{email},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			if err := m.Remove(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_Show(t *testing.T) {
	userMock := &mocks.Db{}

	userData := users.StoredUser{
		Email:     "ehab@test.com",
		Firstname: "Ehab",
		Lastname:  "Terra",
		Role:      "admin",
		Isactive:  false,
	}

	var res users.StoredUser

	userMock.On("Load", userData.Email, &res).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*users.StoredUser)
		*arg = userData
		fmt.Printf("value %v, type %T \n", arg, arg)
	})

	type fields struct {
		Db storage.Db
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *users.StoredUser
		wantErr bool
	}{
		{
			"Show",
			fields{userMock},
			args{userData.Email},
			&userData,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			gotRes, err := m.Show(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Show() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestManager_Update(t *testing.T) {
	userMock := &mocks.Db{}
	roleManager := role.NewManager(userMock)

	user := &users.User{
		Email:     "ehab@test.com",
		Firstname: "Ehab",
		Lastname:  "Terra",
		Role:      "admin",
		Isactive:  true,
	}
	sp := &users.StoredUser{
		Email:     user.Email,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Isactive:  user.Isactive,
		Role:      user.Role,
	}
	wantDes := "Administrator"
	testRole := &roles.StoredRole{Name: user.Role, Description: &wantDes}

	userMock.On("Save", user.Email, sp).Return(nil)
	userMock.On("Load", user.Role, &roles.StoredRole{}).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*roles.StoredRole)
		fmt.Print(arg)
		arg.Name = user.Role
		arg.Description = testRole.Description
	})

	type fields struct {
		Db storage.Db
		M  *role.Manager
	}
	type args struct {
		p *users.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Update",
			fields{userMock, roleManager},
			args{user},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db:   tt.fields.Db,
				role: tt.fields.M,
			}
			if err := m.Update(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewManager(t *testing.T) {
	userMock := &mocks.Db{}
	roleManager := role.NewManager(userMock)
	type args struct {
		db storage.Db
	}
	tests := []struct {
		name string
		args args
		want *Manager
	}{
		{
			"NewManager",
			args{userMock},
			&Manager{userMock, roleManager},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewManager(tt.args.db, roleManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
