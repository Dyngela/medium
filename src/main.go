package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github/Dyngela/medium/src/product"
	"github/Dyngela/medium/src/user"
	"github/Dyngela/medium/src/utils"
	"log"
	"os"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	utils.ConnectToPostgres()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Printf("Starting server on port %s\n", os.Getenv("APP_PORT"))
	//Setup gin router
	router := gin.Default()
	//Parameterize cors to allow remote connection e.g. your front-end
	router.Use(cors.Default())
	//Call the method that gather all our endpoints
	controllers(router)
	//Launch the actual server
	err = router.Run(os.Getenv("APP_PORT"))
	//Check for eventual error, logging it if it happens
	utils.CheckForError(&err, "error while trying to run gin's server")
}

//controllers is going to gather all of our domain's controllers
func controllers(router *gin.Engine) {
	product.ProductsController(router)
	user.UserController(router)
}
