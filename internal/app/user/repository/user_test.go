package repository

import (
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/model"
	"testing"
)

func TestUserRepository_GetRole(t *testing.T) {
	sto, err := CreateUserRepository(app.Middle)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userId: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetRole(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_ChangeRoleExp(t *testing.T) {
	sto, err := CreateUserRepository(app.Middle)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
	}
	type args struct {
		userID int
		exp    int
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userID: 1, exp: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.ChangeRoleExp(tt.args.userID, tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.ChangeRoleExp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_ChangeRoleAttr(t *testing.T) {
	sto, err := CreateUserRepository(app.Middle)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
	}
	type args struct {
		userID int
		attr   string
		val    int
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userID: 1, attr: model.RoleStrength, val: 1}, false},
		{"case2", sto, args{userID: 1, attr: model.RoleCulture, val: 1}, false},
		{"case3", sto, args{userID: 1, attr: model.RoleEnvironment, val: 1}, false},
		{"case4", sto, args{userID: 1, attr: model.RoleCharisma, val: 1}, false},
		{"case5", sto, args{userID: 1, attr: model.RoleTalent, val: 1}, false},
		{"case5", sto, args{userID: 1, attr: model.RoleIntellect, val: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.ChangeRoleAttr(tt.args.userID, tt.args.attr, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.ChangeRoleAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}