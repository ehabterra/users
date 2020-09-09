package role

import (
	"fmt"
	"reflect"
	"testing"
	"users/gen/roles"
	"users/mocks"
	storage "users/pkg/db"

	"github.com/stretchr/testify/mock"
)

func Test_Add(t *testing.T) {
	name := "admin"
	des := "Administrator"
	roleMock := mocks.Db{}
	sp := roles.StoredRole{Name: name, Description: &des}
	r := roles.Role{Name: name, Description: &des}

	roleMock.On("NewID", storage.RoleBucket).Return("1", nil)
	roleMock.On("Save", storage.RoleBucket, name, &sp).Return(nil)

	type fields struct {
		db storage.Db
	}
	type args struct {
		p *roles.Role
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes string
		wantErr bool
	}{
		{
			"Add",
			fields{&roleMock},
			args{&r},
			name,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Manager{
				Db: tt.fields.db,
			}
			err := s.Add(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_List(t *testing.T) {
	roleMock := mocks.Db{}
	des := "Administrator"
	wantRes := roles.StoredRoleCollection{{"admin", &des}}
	var res roles.StoredRoleCollection

	roleMock.On("LoadAll", storage.RoleBucket, &res).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*roles.StoredRoleCollection)
		*arg = append(*arg, &roles.StoredRole{Name: "admin", Description: &des})
		fmt.Printf("value %v, type %T \n", arg, arg)
	})

	type fields struct {
		db storage.Db
	}

	tests := []struct {
		name    string
		fields  fields
		wantRes roles.StoredRoleCollection
		wantErr bool
	}{
		{
			"Test_List",
			fields{&roleMock},
			wantRes,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Manager{
				Db: tt.fields.db,
			}
			gotRes, err := s.List()
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

func Test_Remove(t *testing.T) {
	roleMock := mocks.Db{}
	wantDelete := "admin"

	roleMock.On("Delete", storage.RoleBucket, wantDelete).Return(nil)
	s := Manager{Db: &roleMock}

	if err := s.Remove(wantDelete); err != nil {
		t.Errorf("Remove() error = %v", err)
	}
}

func Test_Show(t *testing.T) {
	roleMock := mocks.Db{}
	wantShow := "admin"
	wantShow2 := "user"
	wantDes := "Administrator"
	res := &roles.StoredRole{}
	wantRes := &roles.StoredRole{Name: wantShow, Description: &wantDes}

	type args struct {
		name string
	}
	roleMock.On("Load", storage.RoleBucket, wantShow, res).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*roles.StoredRole)
		fmt.Print(arg)
		arg.Name = wantShow
		arg.Description = &wantDes
	})
	roleMock.On("Load", storage.RoleBucket, wantShow2, res).Return(storage.ErrNotFound)

	type fields struct {
		db storage.Db
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *roles.StoredRole
		wantErr bool
	}{
		{
			"Test_Show",
			fields{&roleMock},
			args{wantShow},
			wantRes,
			false,
		},
		{
			"ShowWithNotFoundError",
			fields{&roleMock},
			args{wantShow2},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Manager{
				Db: tt.fields.db,
			}
			gotRes, err := s.Show(tt.args.name)
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

func Test_Update(t *testing.T) {
	name := "admin"
	des := "Administrator"
	roleMock := mocks.Db{}
	sp := roles.StoredRole{Name: name, Description: &des}
	r := roles.Role{Name: name, Description: &des}

	roleMock.On("Save", storage.RoleBucket, name, &sp).Return(nil)

	type fields struct {
		db storage.Db
	}
	type args struct {
		p *roles.Role
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes string
		wantErr bool
	}{
		{
			"Update",
			fields{&roleMock},
			args{&r},
			name,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Manager{
				Db: tt.fields.db,
			}
			err := s.Update(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
