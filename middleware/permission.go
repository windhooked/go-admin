package middleware

import (
	mycasbin "go-admin/pkg/casbin"
	"go-admin/pkg/jwtauth"
	_ "go-admin/pkg/jwtauth"
	"go-admin/tools"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Permission check middleware
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(jwtauth.MapClaims)
		e, err := mycasbin.Casbin()
		tools.HasError(err, "", 500)
		//Check permissions
		res, err := e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		log.Println("----------------", v["rolekey"], c.Request.URL.Path, c.Request.Method)

		tools.HasError(err, "", 500)

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "Sorry, you do not have access to this interface, please contact the administrator",
			})
			c.Abort()
			return
		}
	}
}
