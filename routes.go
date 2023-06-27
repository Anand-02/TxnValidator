package main

import (
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
	// fmt.Println(txn)
}

