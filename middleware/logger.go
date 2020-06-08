package middleware

import (
	"go-admin/models"
	"go-admin/tools"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	//	"github.com/sirupsen/logrus"
)

//Instantiate
//var logger = logrus.New()

// log to file
func LoggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Starting time
		startTime := time.Now()

		// process request
		c.Next()

		// End Time
		endTime := time.Now()

		// execution time
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := c.Request.Method

		// Request routing
		reqUri := c.Request.RequestURI

		// status code
		statusCode := c.Writer.Status()

		// Request IP
		clientIP := c.ClientIP()

		// log format
		//		logger.Infof("%s [%s] %3d %13v %15s %s %s \r\n", startTime.Format("2006-01-02 15:04:05.9999"), strings.ToUpper(logger.Level.String()), statusCode, latencyTime, clientIP, reqMethod, reqUri,)
		log.Info().Int("statusCode", statusCode).Dur("latencyTime", latencyTime).Str("clientIP", clientIP).Str("reqMethod", reqMethod).Str("reqUri", reqUri)

		if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" {
			menu := models.Menu{}
			menu.Path = reqUri
			menu.Action = reqMethod
			menuList, _ := menu.Get()
			sysOperLog := models.SysOperLog{}
			sysOperLog.OperIp = clientIP
			sysOperLog.OperLocation = tools.GetLocation(clientIP)
			sysOperLog.Status = tools.IntToString(statusCode)
			sysOperLog.OperName = tools.GetUserName(c)
			sysOperLog.RequestMethod = c.Request.Method
			sysOperLog.OperUrl = reqUri
			if reqUri == "/login" {
				sysOperLog.BusinessType = "10"
				sysOperLog.Title = "user login"
				sysOperLog.OperName = "-"
			} else if strings.Contains(reqUri, "/api/v1/logout") {
				sysOperLog.BusinessType = "11"
			} else if strings.Contains(reqUri, "/api/v1/getCaptcha") {
				sysOperLog.BusinessType = "12"
				sysOperLog.Title = "Captcha"
			} else {
				if reqMethod == "POST" {
					sysOperLog.BusinessType = "1"
				} else if reqMethod == "PUT" {
					sysOperLog.BusinessType = "2"
				} else if reqMethod == "DELETE" {
					sysOperLog.BusinessType = "3"
				}
			}
			sysOperLog.Method = reqMethod
			if len(menuList) > 0 {
				sysOperLog.Title = menuList[0].Title
			}
			b, _ := c.Get("body")
			sysOperLog.OperParam, _ = tools.StructToJsonStr(b)
			sysOperLog.CreateBy = tools.GetUserName(c)
			sysOperLog.OperTime = tools.GetCurrntTime()
			sysOperLog.LatencyTime = (latencyTime).String()
			sysOperLog.UserAgent = c.Request.UserAgent()
			if c.Err() == nil {
				sysOperLog.Status = "0"
			} else {
				sysOperLog.Status = "1"
			}
			_, _ = sysOperLog.Create()
		}
	}
}
