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
// @Router /api/v1/operloglist [get]
// @Security
func GetOperLogList(c *gin.Context) {
	var data models.SysOperLog
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

	data.OperName = c.Request.FormValue("operName")
	data.Status = c.Request.FormValue("status")
	data.OperIp = c.Request.FormValue("operIp")
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
// @Router /api/v1/operlog/{infoId} [get]
// @Security
func GetOperLog(c *gin.Context) {
	var OperLog models.SysOperLog
	OperLog.OperId, _ = tools.StringToInt(c.Param("operId"))
	result, err := OperLog.Get()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary add operation log
// @Description get JSON
// @Tags operation log
// @Accept  application/json
// @Product application/json
// @Param data body models.SysOperLog true "data"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/operlog [post]
// @Security Bearer
func InsertOperLog(c *gin.Context) {
	var data models.SysOperLog
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary bulk delete operation logs
// @Description delete data
// @Tags operation log
// @Param operId path string true "operId separated by comma (,)"
// @Success 200 {string} string "{"code": 200, "message": "Successfully deleted"}"
// @Success 200 {string} string "{"code": -1, "message": "Deletion failed"}"
// @Router /api/v1/operlog/{operId} [delete]
func DeleteOperLog(c *gin.Context) {
	var data models.SysOperLog
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("operId", c)
	_, err := data.BatchDelete(IDS)
	tools.HasError(err, "Deletion failed", 500)
	var res app.Response
	res.Msg = "Successfully deleted"
	c.JSON(http.StatusOK, res.ReturnOK())
}
