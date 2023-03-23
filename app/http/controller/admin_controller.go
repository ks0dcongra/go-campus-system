package controller

import (
	"example1/app/model"
	"example1/app/model/responses"
	"example1/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminController struct {
}

func AdminController() *adminController {
	return &adminController{}
}

func (h *adminController) AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.AddItem)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil, "From DB"))
			return
		}
		item_id, status := service.NewItemService().Add(requestData)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil, "From DB"))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, item_id, "From DB"))
	}
}
func (h *adminController) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestData := new(model.SearchItem)
		if err := c.ShouldBindJSON(requestData); err != nil {
			c.JSON(http.StatusOK, responses.Status(responses.ParameterErr, nil, "From DB"))
			return
		}
		item, status := service.NewItemService().Get(requestData)
		if status != responses.Success {
			c.JSON(http.StatusOK, responses.Status(responses.Error, nil, "From DB"))
			return
		}
		c.JSON(http.StatusOK, responses.Status(responses.Success, item, "From DB"))
	}
}
