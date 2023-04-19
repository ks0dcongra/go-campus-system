package service

import (
	"example1/app/model"
	"reflect"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	type args struct {
		condition *model.LoginStudent
	}
	tests := []struct {
		name            string
		h               *UserService
		args            args
		wantStudent     model.Student
		wantStatus      string
		wantTokenResult string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStudent, gotStatus, gotTokenResult := tt.h.Login(tt.args.condition)
			if !reflect.DeepEqual(gotStudent, tt.wantStudent) {
				t.Errorf("UserService.Login() gotStudent = %v, want %v", gotStudent, tt.wantStudent)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("UserService.Login() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotTokenResult != tt.wantTokenResult {
				t.Errorf("UserService.Login() gotTokenResult = %v, want %v", gotTokenResult, tt.wantTokenResult)
			}
		})
	}
}
