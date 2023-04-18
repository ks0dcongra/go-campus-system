package repository

import (
	"example1/app/model"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func Test_UserRepository_CheckUserPassword(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		name            string
		h               *_UserRepository
		args            args
		wantStudent     model.Student
		wantResult      *gorm.DB
		wantTokenResult string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStudent, gotResult, gotTokenResult := tt.h.CheckUserPassword(tt.args.condition)
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("_UserRepository.CheckUserPassword() gotStudent = %v, want %v", gotStudent, tt.wantStudent)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("_UserRepository.CheckUserPassword() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotTokenResult != tt.wantTokenResult {
				t.Errorf("_UserRepository.CheckUserPassword() gotTokenResult = %v, want %v", gotTokenResult, tt.wantTokenResult)
			}
		})
	}
}
