package product

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github/Dyngela/medium/src/utils"
	"net/http"
)

func getProductById(c *gin.Context) {
	// We create a struct which correspond to our argument.
	type URI struct {
		// We don't have to use the same properties name that we use in our controller.
		// To do so you can pass an argument uri which gonna do the mapping.
		// We can also do some binding, here we want an argument greater or equal to one
		IdProduct uint `binding:"gte=1" uri:"id"`
	}
	uri := URI{}
	// We bind gin's context to our struct
	if err := c.ShouldBindUri(&uri); err != nil {
		// utils function to make a response already formatted.
		utils.ThrowExceptionBadArgument(c, err)
		return
	}

	// We get the actual response from our repository
	response, err := findProductById(uri.IdProduct, c)
	if err != nil {
		// utils function to make a response already formatted
		utils.ThrowExceptionSQLError(c, err, response)
		return
	}
	// if everything ran smoothly we return our response from the repository and a 200 status code
	c.JSON(http.StatusOK, response)
}

func createProduct(c *gin.Context) {
	body := CreateProductRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ThrowExceptionBadArgument(c, err)
		return
	}
	// We now validated our json, now we bind it to an actual business object
	product := Product{
		Name:  body.Name,
		Price: body.Price,
		// Since image is nullable and sql.NullString doesn't take string right away, we have to do that.
		Image:      sql.NullString{String: body.Image, Valid: true},
		IsInStock:  body.IsInStock,
		CategoryId: body.CategoryId,
	}

	response, err := createProductRepo(product, c)
	if err != nil {
		utils.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func getAllProduct(c *gin.Context) {
	response, err := findAllProduct(c)
	if err != nil {
		utils.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func updateProduct(c *gin.Context) {
	body := UpdateProductRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ThrowExceptionBadArgument(c, err)
		return
	}
	// We now validated our json, now we bind it to an actual business object
	product := Product{
		ProductId: body.ProductId,
		Name:      body.Name,
		Price:     body.Price,
		// Since image is nullable and sql.NullString doesn't take string right away, we have to do that.
		Image:      sql.NullString{String: body.Image, Valid: true},
		IsInStock:  body.IsInStock,
		CategoryId: body.CategoryId,
	}

	response, err := updateProductRepo(product, c)
	if err != nil {
		utils.ThrowExceptionSQLError(c, err, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func deleteProduct(c *gin.Context) {
	type URI struct {
		IdProduct uint `binding:"gte=1" uri:"id"`
	}
	uri := URI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		utils.ThrowExceptionBadArgument(c, err)
		return
	}

	response, err := deleteProductRepo(uri.IdProduct, c)
	if err != nil {
		utils.ThrowExceptionSQLError(c, err, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
