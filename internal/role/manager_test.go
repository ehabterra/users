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

	roleMock.On("NewID").Return("1", nil)
	roleMock.On("Save", name, &sp).Return(nil)

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
	wantRes := roles.StoredRoleCollection{{Name: "admin", Description: &des}}
	var res roles.StoredRoleCollection

	roleMock.On("LoadAll", &res).Return(func(args interface{}) error {
		arg := args.(*roles.StoredRoleCollection)
		*arg = append(*arg, &roles.StoredRole{Name: "admin", Description: &des})
		fmt.Printf("value %v, type %T \n", arg, arg)
		return nil
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

	roleMock.On("Delete", wantDelete).Return(nil)
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
	roleMock.On("Load", wantShow, res).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*roles.StoredRole)
		fmt.Print(arg)
		arg.Name = wantShow
		arg.Description = &wantDes
	})
	roleMock.On("Load", wantShow2, res).Return(storage.ErrNotFound)

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

	roleMock.On("Save", name, &sp).Return(nil)

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

func TestManager_CheckRoleExists(t *testing.T) {
	userMock := &mocks.Db{}
	roleName := "admin"
	roleDesc := "Administrator"
	emptyRole := &roles.StoredRole{}

	userMock.On("Load", roleName, emptyRole).Return(func(key string, res interface{}) error {
		role := res.(*roles.StoredRole)
		role.Name = "admin"
		role.Description = &roleDesc

		return nil
	})
	type fields struct {
		Db storage.Db
	}
	type args struct {
		role string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"CheckRoleExistence",
			fields{userMock},
			args{roleName},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				Db: tt.fields.Db,
			}
			got, err := m.CheckRoleExists(tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckRoleExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckRoleExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
