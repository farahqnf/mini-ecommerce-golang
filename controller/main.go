package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tugasmeilyanto/go-trial-class/config"
	"github.com/tugasmeilyanto/go-trial-class/entity"
)

type OrderRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductId    int    `json:"product_id"`
}

func HandlerGetProduct(ctx *gin.Context) {
	var products []entity.Product
	err := config.DB.Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func HandlerPostOrder(ctx *gin.Context) {
	// menerima data
	var orderBody OrderRequest
	err := ctx.ShouldBindJSON(&orderBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "terjadi kesalahan dalam parsing data",
		})
		return
	}

	var product entity.Product
	result := config.DB.Where("ID = ?", orderBody.ProductId).First(&product)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "produk tidak ditemukan",
		})
		return
	}

	newOrder := entity.Order{
		BuyerEmail:   orderBody.BuyerEmail,
		BuyerAddress: orderBody.BuyerAddress,
		ProductId:    int(product.ID),
		OrderDate:    time.Now(),
	}

	err = config.DB.Create(&newOrder).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "terjadi kesalahan saat membuat order",
		})
		return
	}
	ctx.JSON(http.StatusOK, "sukses membuat order")
}

func HandlerGetOrder(ctx *gin.Context) {
	var orders []entity.Order

	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
