package product

import "github.com/gin-gonic/gin"

// ProductsController Every API routes mainly related to product table
func ProductsController(router *gin.Engine) {
	// This is a base path for our route.
	v1 := router.Group("/api/v1/product/")
	{
		// We have to precise the verb, then the path and finally the method which will be called
		// Argument such as uri, or body are contained in the gin.Context
		// that's why there's no parameter on method call
		// Here the final path will looks like: http://localhost:8080/api/v1/product/all
		v1.GET("all", getAllProduct)
		// If we want to pass an url parameter
		v1.GET(":id", getProductById)
		v1.POST("", updateProduct)
		v1.PUT("", createProduct)
		v1.DELETE(":id", deleteProduct)
	}

}
