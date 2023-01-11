package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github/Dyngela/medium/src/product"
	"github/Dyngela/medium/src/utils"
	"log"
)

const webport = ":8080"

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	utils.ConnectToPostgres()
}

func main() {
	log.Printf("Starting server on port %s\n", webport)
	//Setup gin router
	router := gin.Default()
	//Parameterize cors to allow remote connection e.g. your front-end
	router.Use(cors.Default())
	//Call the method that gather all our endpoints
	controllers(router)
	//Launch the actual server
	err := router.Run(webport)
	//Check for eventual error, logging it if it happens
	utils.CheckForError(&err, "error while trying to run gin's server")
}

//controllers is going to gather all of our domain's controllers
func controllers(router *gin.Engine) {
	product.ProductsController(router)
}
