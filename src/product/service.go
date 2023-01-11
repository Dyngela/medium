package product

import (
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
	response, err := findProductById(uri.IdProduct)
	if err != nil {
		// utils function to make a response already formatted
		utils.ThrowExceptionSQLError(c, err, response)
		return
	}
	// if everything ran smoothly we return our response from the repository and a 200 status code
	c.JSON(http.StatusOK, response)
}

func getAllProduct(c *gin.Context) {
	
}

func createProduct(c *gin.Context) {

}

func updateProduct(c *gin.Context) {

}

func deleteProduct(c *gin.Context) {

}
