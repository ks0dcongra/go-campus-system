package test

import (
	"example1/app/http/controller"
	"example1/app/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepoImplMock struct {
    mock.Mock
}

func (serviceMock *EmployeeRepoImplMock) Login(*model.LoginStudent) (model.Student, string){
    args := serviceMock.Called()
    log.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",args.Get(0))
	
    return args.Get(0).(model.Student),"test"
}

func Test_userController_LoginUser2(t *testing.T) {
	serviceMock := new(EmployeeRepoImplMock)
    serviceMock.On("Login").
        Return(model.Student{Id: 99, Name: "Jack", Password: "12345678"},"test")
    expected := 13

    es := controller.UserController{
        UserService: serviceMock,
    }
    actial := es.LoginUser()
    log.Println("AAA", actial)
    assert.Equal(t, expected, 1)
}

// use mock to implments CalculatorService's method
func (repoMock *EmployeeRepoImplMock) FindEmployeesAgeGreaterThan(age int) []model.Employee {
	
	args := repoMock.Called(age)
	log.Println("argsargsargsargsargsargsargsargsargsargsargs",args.Get(0).([]model.Employee))
	return args.Get(0).([]model.Employee)
}

func TestGetSrEmployeeNumbers_Age400(t *testing.T) {

	repoMock := new(EmployeeRepoImplMock)
	repoMock.On("FindEmployeesAgeGreaterThan", 40).
		Return([]model.Employee{
			{ID: 99, Name: "Jack", Age: 70},
		})

	expected := 1

	es := controller.EmployeeSerivceImpl{
		EmpRepo: repoMock,
	}

	actial := es.GetSrEmployeeNumbers(40)
    log.Println("BBB",actial)
	assert.Equal(t, expected, actial)
}


// 