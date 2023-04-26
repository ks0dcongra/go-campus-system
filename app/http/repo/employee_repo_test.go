package repo

import (
	"testing"

	"example1/app/model"
	"github.com/stretchr/testify/assert"
)

func TestFindEmployeesAgeGreaterThan30(t *testing.T) {
	testCase := struct {
		age      int
		expected []model.Employee
	}{
		30,
		[]model.Employee{
			{ID: 1, Name: "John", Age: 33},
			{ID: 3, Name: "Mike", Age: 45},
		},
	}

	empRepo := EmployeeRepoImpl{}
	actial := empRepo.FindEmployeesAgeGreaterThan(30)

	assert.Equal(t, testCase.expected, actial)
}
