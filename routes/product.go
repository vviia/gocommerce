package routes

import (
	"gocommerce/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, controller product.IProductController) {
	r.POST("/admin/addproduct", controller.ProductViewerAdmin)
}
