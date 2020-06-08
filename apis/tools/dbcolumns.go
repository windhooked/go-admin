package tools

import (
	"go-admin/models/tools"
	tools2 "go-admin/tools"
	"go-admin/tools/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary page list data / page list data
// @Description database table column page list / database table column page list
// @Tags tools / Tools
// @Param tableName query string false "tableName / data table name"
// @Param pageSize query int false "pageSize / number of pages"
// @Param pageIndex query int false "pageIndex / page number"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/columns/page [get]
func GetDBColumnList(c *gin.Context) {
	var data tools.DBColumns
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools2.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools2.StrToInt(err, index)
	}

	data.TableName = c.Request.FormValue("tableName")
	tools2.Assert(data.TableName == "", "table name cannot be empty！", 500)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools2.HasError(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	var res app.Response
	res.Data = mp

	c.JSON(http.StatusOK, res.ReturnOK())
}
