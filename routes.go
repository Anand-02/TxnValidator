package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Handler(c *gin.Context) {
	var txn []map[string]Data
	if err := c.ShouldBindJSON(&txn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, mp2 := range txn {
		Validator(mp2)
	}
	fmt.Println(txn)
}

func createRoute() {
	router.POST("/post", Handler)
	err := router.Run("localhost:8000")
	Err(err)
}