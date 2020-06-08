package log

import (
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary login log list
// @Description get JSON
// @Tags login log
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "Number of pages"
// @Param pageIndex query int false "page number"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/loginloglist [get]
// @Security
func GetLoginLogList(c *gin.Context) {
	var data models.LoginLog
	var err error
	var pageSize = 10
	var pageIndex = 1

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Username = c.Request.FormValue("username")
	data.Status = c.Request.FormValue("status")
	data.Ipaddr = c.Request.FormValue("ipaddr")
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	var res app.Response
	res.Data = mp
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary gets the login log by coding
// @Description get JSON
// @Tags login log
// @Param infoId path int true "infoId"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/loginlog/{infoId} [get]
// @Security
func GetLoginLog(c *gin.Context) {
	var LoginLog models.LoginLog
	LoginLog.InfoId, _ = tools.StringToInt(c.Param("infoId"))
	result, err := LoginLog.Get()
	tools.HasError(err, "Sorry no relevant information was found", -1)

	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary add login log
// @Description get JSON
// @Tags login log
// @Accept  application/json
// @Product application/json
// @Param data body models.LoginLog true "data"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/loginlog [post]
// @Security Bearer
func InsertLoginLog(c *gin.Context) {
	var data models.LoginLog
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary modify login log
// @Description get JSON
// @Tags login log
// @Accept  application/json
// @Product application/json
// @Param data body models.LoginLog true "body"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/loginlog [put]
// @Security Bearer
func UpdateLoginLog(c *gin.Context) {
	var data models.LoginLog
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "", -1)
	result, err := data.Update(data.InfoId)
	tools.HasError(err, "", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary delete login logs in batch
// @Description delete data
// @Tags login log
// @Param infoId path string true "infoId separated by comma (,)"
// @Success 200 {string} string "{"code": 200, "message": "Successfully deleted"}"
// @Success 200 {string} string "{"code": -1, "message": "Deletion failed"}"
// @Router /api/v1/loginlog/{infoId} [delete]
func DeleteLoginLog(c *gin.Context) {
	var data models.LoginLog
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("infoId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, "Modification failed", 500)
	var res app.Response
	res.Msg = "Successfully deleted"
	c.JSON(http.StatusOK, res.ReturnOK())
}
