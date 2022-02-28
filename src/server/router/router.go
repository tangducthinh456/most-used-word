package router

import (
	"github.com/gin-gonic/gin"
	"most-used-word/src/server/handle"
)

const ENDPOINT_TOP_USE = "top-used"

var router *gin.Engine


// Package router uses singleton pattern which init() function
// acts like constructor for router package and variable router create once
 
func init() {
	if router == nil {
		router = gin.New()

		router.Use(gin.Logger())
		router.Use(gin.Recovery())

		router.POST(ENDPOINT_TOP_USE, handle.HandleTopUsed)
	}
}

func Router() *gin.Engine {
	return router
}
