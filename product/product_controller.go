package product

import (
	"context"
	"gocommerce/presentation"
	"gocommerce/product/dto"
	"gocommerce/product/services"
	"gocommerce/shared/entities"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	ProductViewerAdmin(c *gin.Context)
}

type ProductController struct {
	ProductService services.IProductService
}

func NewProductController(
	productService services.IProductService,
) IProductController {
	return &ProductController{
		ProductService: productService,
	}
}

//This is function to add products
//this is an admin part
//json should look like this
// post request : http://localhost:8080/admin/addproduct
/*
json
{
"product_name" : "pencil"
"price"        : 98
"rating"       : 10
"image"        : "image-url"
}
*/
func (d *ProductController) ProductViewerAdmin(c *gin.Context) {
	in := dto.AddProductRequest{}

	if err := presentation.ReadRestIn(c, &in); err != nil {
		o := entities.CommonResult{}
		o.SetCode(400, err)
		presentation.WriteRestOut(c, o, &o)
		return
	}

	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	out := d.ProductService.ProductViewerAdminCtx(ctx, &in)

	presentation.WriteRestOut(c, out, &out.CommonResult)
}
