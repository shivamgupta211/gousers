package core

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// GousersRouter ... custom gin router
type GousersRouter struct {
	router   *gin.Engine
	basePath string
}

// ReturnType ... struct for all api return
type ReturnType struct {
	Status   int
	Response interface{}
}

// RouteHandler ... It's the callback function of each router
type RouteHandler func(*gin.Context) (ReturnType, error)

// NewResponse ... constructor for response
func NewResponse(status int, response interface{}) ReturnType {
	return ReturnType{Status: status, Response: response}
}

// custom router : GousersRouter

// GET ... similar to git router GET
func (ir *GousersRouter) GET(relativePath string, handler RouteHandler) gin.IRoutes {
	return ir.router.GET(ir.basePath+relativePath, func(c *gin.Context) {
		r, err := handler(c)
		status := 200
		if err != nil {
			status = 500
		}
		if r.Status != 0 {
			status = r.Status
		}
		c.JSON(status, gin.H{
			"result": r.Response,
		})
	})
}

// POST ... similar to git router GET
func (ir *GousersRouter) POST(relativePath string, handler RouteHandler) gin.IRoutes {
	return ir.router.POST(ir.basePath+relativePath, func(c *gin.Context) {
		r, err := handler(c)
		status := 200
		if err != nil {
			status = 500
		}
		if r.Status != 0 {
			status = r.Status
		}
		c.JSON(status, gin.H{
			"result": r.Response,
		})
	})
}

// PUT ... similar to git router GET
func (ir *GousersRouter) PUT(relativePath string, handler RouteHandler) gin.IRoutes {
	return ir.router.PUT(ir.basePath+relativePath, func(c *gin.Context) {
		r, err := handler(c)
		status := 200
		if err != nil {
			status = 500
		}
		if r.Status != 0 {
			status = r.Status
		}
		c.JSON(status, gin.H{
			"result": r.Response,
		})
	})
}

// DELETE ... similar to git router GET
func (ir *GousersRouter) DELETE(relativePath string, handler RouteHandler) gin.IRoutes {
	return ir.router.DELETE(ir.basePath+relativePath, func(c *gin.Context) {
		r, err := handler(c)
		status := 200
		if err != nil {
			status = 500
		}
		if r.Status != 0 {
			status = r.Status
		}
		c.JSON(status, gin.H{
			"result": r.Response,
		})
	})
}

// Run ... similar to git router GET
func (ir *GousersRouter) Run(port string) (err error) {
	return ir.router.Run(":" + port)
}

// Bootstrap ... scaffold it from server.go
func Bootstrap(basePath string) *GousersRouter {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          500 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	return &GousersRouter{
		router:   router,
		basePath: basePath,
	}
}
