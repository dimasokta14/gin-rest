package product

import (
	"net/http"

	"github.com/dimasokta14/gin-rest/entities"
	"github.com/gin-gonic/gin"
)

func (h *handler) AddProduct(ctx *gin.Context) {

	body := entities.AddProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var product entities.Product

	product.Title = body.Title
	product.Desription = body.Desription
	product.SKU = body.SKU
	product.Stock = body.Stock
	product.ImageUrl = body.ImageUrl

	if result := h.DB.Create(&product); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &product)
}
