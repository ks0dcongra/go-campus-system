package test

import (
	
	"example1/app/model"
	"example1/app/model/responses"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserController struct {
    mock.Mock
}

func (controllerMock *MockUserController) LoginUser() gin.HandlerFunc{
    args := controllerMock.Called()
    log.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",args)
    mockUser := model.Student{
        Id:   123,
        Name: "testuser",
    }
   
    return mockUser,"test"
}

func Test_userController_LoginUser3(t *testing.T) {
	controllerMock := new(MockUserController)
    controllerMock.On("LoginUser").
        Return(func(c *gin.Context) {
			c.JSON(200, responses.Status("0", nil))
		})
    expected := 13

    es := controller.UserController{
        UserController: controllerMock,
    }
    actial := es.LoginUser()
    log.Println("AAA",actial)
    assert.Equal(t, expected, 1)
}

// use mock to implments CalculatorService's method
func (repoMock *EmployeeRepoImplMock) FindEmployeesAgeGreaterThan(age int) []model.Employee {
	
	args := repoMock.Called(age)
	log.Println("argsargsargsargsargsargsargsargsargsargsargs",args)
	return args.Get(0).([]model.Employee)
}

func TestGetSrEmployeeNumbers_Age40(t *testing.T) {

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
