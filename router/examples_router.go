package router

import (
	"go-admin/pkg/jwtauth"
	jwt "go-admin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

// routing example
func InitExamplesRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// routing without authentication
	examplesNoCheckRoleRouter(r)
	// Routes requiring authentication
	examplesCheckRoleRouter(r, authMiddleware)

	return r
}

// Example routing without authentication
func examplesNoCheckRoleRouter(r *gin.Engine) {

	//v1 := r.Group("/api/v1")
	//v1.GET("/examples/list", examples.apis)

}

// Examples of routes that require authentication
func examplesCheckRoleRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {

	//v1auth := r.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	//{
	//	v1auth.GET("/examples/list", examples.apis)
	//}
}
