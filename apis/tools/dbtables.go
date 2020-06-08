package tools

import (
	"go-admin/models/tools"
	tools2 "go-admin/tools"
	"go-admin/tools/app"
	config2 "go-admin/tools/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary page list data / page list data
// @Description database table page list / database table page list
// @Tags tools / Tools
// @Param tableName query string false "tableName / data table name"
// @Param pageSize query int false "pageSize / number of pages"
// @Param pageIndex query int false "pageIndex / page number"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/tables/page [get]
func GetDBTableList(c *gin.Context) {
	var res app.Response
	var data tools.DBTables
	var err error
	var pageSize = 10
	var pageIndex = 1
	if config2.DatabaseConfig.Dbtype == "sqlite3" {
		res.Msg = "Sorry, sqlite3 does not support code generation yet!"
		c.JSON(http.StatusOK, res.ReturnError(500))
		return
	}

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools2.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools2.StrToInt(err, index)
	}

	data.TableName = c.Request.FormValue("tableName")
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools2.HasError(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	res.Data = mp

	c.JSON(http.StatusOK, res.ReturnOK())
}
