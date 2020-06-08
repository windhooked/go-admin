package router

import (
	"go-admin/middleware"
	_ "go-admin/pkg/jwtauth"
	"go-admin/tools"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	// log processing
	r.Use(middleware.LoggerToFile())
	// Custom error handling
	r.Use(middleware.CustomError)
	// NoCache is a middleware function that appends headers
	r.Use(middleware.NoCache)
	// Cross-domain processing
	r.Use(middleware.Options)
	// Secure is a middleware function that appends security
	r.Use(middleware.Secure)
	// Set X-Request-Id header
	r.Use(middleware.RequestId())
	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()
	tools.HasError(err, "JWT Init Error", 500)

	// Register system routing
	InitSysRouter(r, authMiddleware)

	// Register business route
	// TODO: Business routing can be stored here, there is no actual routing inside there is demo code
	InitExamplesRouter(r, authMiddleware)

	return r
}
