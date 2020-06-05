package pingmodule

import (
	"gousers/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BootstrapRoutes ... to bootstrap all routes of this module
func BootstrapRoutes(ir *core.GousersRouter) {
	ir.GET("/ping", func(c *gin.Context) (core.ReturnType, error) {
		return core.NewResponse(http.StatusOK, gin.H{
			"message": "pong",
		}), nil
	})
}
