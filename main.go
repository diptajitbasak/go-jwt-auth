package main

import ("github.com/gin-gonic/gin"
	controller "gitlab.com/test-auth/controller"
	database "gitlab.com/test-auth/database"
)

func main() {
	database.ExampleNewClient()

	r := gin.Default()

	r.GET("/tokenGet", controller.GetToken)
	r.GET("/verify", controller.VerifyToken)
	
	r.Run()
}