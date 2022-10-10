package presentation

import (
	"gocommerce/shared/entities"

	"github.com/gin-gonic/gin"
)

func ReadRestIn[T interface{}](c *gin.Context, in *T) error {
	if err := c.ShouldBindJSON(in); err != nil {
		return err
	}
	return nil
}

func WriteRestOut[T interface{}](c *gin.Context, out T, cr *entities.CommonResult) {
	if cr.ResCode == 0 {
		c.JSON(200, out)
		return
	}

	if cr.ResCode >= 200 && cr.ResCode < 300 {
		c.JSON(cr.ResCode, out)
		return
	}

	if cr.ResCode >= 400 && cr.ResCode < 500 {
		c.AbortWithStatusJSON(cr.ResCode, struct {
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Message: cr.ResMessage,
			Data:    nil,
		})
		return
	}

	c.AbortWithStatusJSON(500, struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "internal server error",
		Data:    nil,
	})
}
