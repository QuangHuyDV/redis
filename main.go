package main

import (
	"redis/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/",HomePage)
	r.GET("/ex7",handler.Ex7)
	r.GET("/ex6",handler.Ex6)

//ex8
	r.GET("/names",handler.GetAll)
	r.POST("/names",handler.InsertName)
	r.GET("/names/:index",handler.ReadName)
	r.POST("/names/:index",handler.UpdateName)
	r.DELETE("/names/:index",handler.DeleteName)

	
//ex9
	r.GET("/user",handler.ReadUser)
	r.POST("/user",handler.InsertUser)

	r.Run()
}



func HomePage(c *gin.Context) {
	c.JSON(200,gin.H{
		"message" : "Hello ",
	})
}