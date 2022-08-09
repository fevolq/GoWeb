package controller

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// api耗时

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiPath := "test"

		start := time.Now()

		c.Next()

		cost := time.Since(start)
		log.Printf("%v消耗时间：%v", apiPath, cost)
	}
}
