package controller

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
}

func UserController() *userController {
	return &userController{}
}

func (h *userController) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.SearchItem)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil))
			return
		}
		item, status := service.NewItemService().Get(requestData)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, item))
	}
}
