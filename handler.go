package main

import "github.com/gin-gonic/gin"

func handler() {
	r := gin.Default()

	r.GET("/person/:person_id/info", GetPersonInformation)
	r.POST("/person/create", CreatePerson)

	r.Run(":8080")

}
