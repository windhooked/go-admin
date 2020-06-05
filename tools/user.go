package tools

import (
	"fmt"
	jwt "go-admin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func GetUserId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return int((data["identity"]).(float64))
	}
	fmt.Println("****************************** path：" + c.Request.URL.Path + "  request method：" + c.Request.Method + "  desctiption：no identity")
	return 0
}

func GetUserIdStr(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return Int64ToString(int64((data["identity"]).(float64)))
	}
	fmt.Println("****************************** path：" + c.Request.URL.Path + "  request method：" + c.Request.Method + "  no identity")
	return ""
}

func GetUserName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["nice"] != nil {
		return (data["nice"]).(string)
	}
	fmt.Println("****************************** path：" + c.Request.URL.Path + "  request method：" + c.Request.Method + "  no nice")
	return ""
}

func GetRoleName(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["rolekey"] != nil {
		return (data["rolekey"]).(string)
	}
	fmt.Println("****************************** path：" + c.Request.URL.Path + "  request method：" + c.Request.Method + "  no rolekey")
	return ""
}

func GetRoleId(c *gin.Context) int {
	data := ExtractClaims(c)
	if data["roleid"] != nil {
		i := int((data["roleid"]).(float64))
		return i
	}
	fmt.Println("****************************** path：" + c.Request.URL.Path + "  request method：" + c.Request.Method + "  no roleid")
	return 0
}
