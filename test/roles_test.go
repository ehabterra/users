package test

import (
	"context"
	"reflect"
	"testing"
	"users/gen/roles"
	"users/internal/accountmanager"
	"users/mocks"
	storage "users/pkg/db"
)

func Test_rolessrvc_Add(t *testing.T) {
	name := "admin"
	des := "Administrator"
	mock := mocks.Db{}
	sp := roles.StoredRole{Name: name, Description: &des}
	r := roles.Role{Name: name, Description: &des}

	mock.On("NewID", storage.RoleBucket).Return("1", nil)
	mock.On("Save", storage.RoleBucket, name, &sp).Return(nil)

	type fields struct {
		db storage.Db
	}
	type args struct {
		ctx context.Context
		p   *roles.Role
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
			fields{&mock},
			args{context.TODO(), &r},
			name,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountmanager.Rolessrvc{
				Db: tt.fields.db,
			}
			gotRes, err := s.Add(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("Add() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_rolessrvc_List(t *testing.T) {
	mock := mocks.Db{}
	view := "default"
	r := &roles.ListPayload{&view}
	des := "Administrator"
	wantRes := roles.StoredRoleCollection{{"admin", &des}}

	mock.On("LoadAll", storage.RoleBucket).Return(wantRes, nil)

	type fields struct {
		db storage.Db
	}
	type args struct {
		ctx context.Context
		p   *roles.ListPayload
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantRes  roles.StoredRoleCollection
		wantView string
		wantErr  bool
	}{
		{
			"Test_rolessrvc_List",
			fields{&mock},
			args{context.TODO(), r},
			wantRes,
			view,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountmanager.Rolessrvc{
				Db: tt.fields.db,
			}
			gotRes, gotView, err := s.List(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("List() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotView != tt.wantView {
				t.Errorf("List() gotView = %v, want %v", gotView, tt.wantView)
			}
		})
	}
}

func Test_rolessrvc_Remove(t *testing.T) {
	mock := mocks.Db{}
	wantDelete := "admin"
	type args struct {
		ctx context.Context
		p   *roles.RemovePayload
	}
	removePayload := &roles.RemovePayload{Name: wantDelete}
	test := args{context.TODO(), removePayload}

	mock.On("Delete", storage.RoleBucket, wantDelete).Return(nil)
	s := accountmanager.Rolessrvc{Db: &mock}

	if err := s.Remove(test.ctx, test.p); err != nil {
		t.Errorf("Remove() error = %v", err)
	}
}

func Test_rolessrvc_Show(t *testing.T) {
	mock := mocks.Db{}
	view := "default"
	wantShow := "admin"
	wantShow2 := "user"
	r := &roles.ShowPayload{Name: wantShow, View: &view}
	r2 := &roles.ShowPayload{Name: wantShow2, View: &view}
	wantDes := "Administrator"
	wantRes := &roles.StoredRole{Name: wantShow, Description: &wantDes}

	mock.On("Load", storage.RoleBucket, wantShow).Return(wantRes, nil)
	mock.On("Load", storage.RoleBucket, wantShow2).Return(nil, storage.ErrNotFound)

	type fields struct {
		db storage.Db
	}
	type args struct {
		ctx context.Context
		p   *roles.ShowPayload
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantRes  *roles.StoredRole
		wantView string
		wantErr  bool
	}{
		{
			"Test_rolessrvc_Show",
			fields{&mock},
			args{context.TODO(), r},
			wantRes,
			view,
			false,
		},
		{
			"ShowWithNotFoundError",
			fields{&mock},
			args{context.TODO(), r2},
			nil,
			view,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountmanager.Rolessrvc{
				Db: tt.fields.db,
			}
			gotRes, gotView, err := s.Show(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Show() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotView != tt.wantView {
				t.Errorf("Show() gotView = %v, want %v", gotView, tt.wantView)
			}
		})
	}
}

func Test_rolessrvc_Update(t *testing.T) {
	name := "admin"
	des := "Administrator"
	mock := mocks.Db{}
	sp := roles.StoredRole{Name: name, Description: &des}
	r := roles.Role{Name: name, Description: &des}

	mock.On("Save", storage.RoleBucket, name, &sp).Return(nil)

	type fields struct {
		db storage.Db
	}
	type args struct {
		ctx context.Context
		p   *roles.Role
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
			fields{&mock},
			args{context.TODO(), &r},
			name,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &accountmanager.Rolessrvc{
				Db: tt.fields.db,
			}
			gotRes, err := s.Update(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("Update() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
