package system

import (
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"

	"github.com/gin-gonic/gin"
)

// @Summary job listing data
// @Description get JSON
// @Tags position
// @Param name query string false "name"
// @Param id query string false "id"
// @Param position query string false "position"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @Security
func GetPostList(c *gin.Context) {
	var data models.Post
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.PostName = c.Request.FormValue("postName")
	id := c.Request.FormValue("postId")
	data.PostId, _ = tools.StringToInt(id)

	data.PostName = c.Request.FormValue("postName")
	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

// @Summary get dictionary data
// @Description get JSON
// @Tags dictionary data
// @Param postId path int true "postId"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{postId} [get]
// @Security
func GetPost(c *gin.Context) {
	var Post models.Post
	Post.PostId, _ = tools.StringToInt(c.Param("postId"))
	result, err := Post.Get()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	app.OK(c, result, "")
}

// @Summary add job
// @Description get JSON
// @Tags position
// @Accept  application/json
// @Product application/json
// @Param data body models.Post true "data"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/post [post]
// @Security Bearer
func InsertPost(c *gin.Context) {
	var data models.Post
	err := c.Bind(&data)
	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Create()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

// @Summary modify position
// @Description get JSON
// @Tags position
// @Accept  application/json
// @Product application/json
// @Param data body models.Dept true "body"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/post/ [put]
// @Security Bearer
func UpdatePost(c *gin.Context) {
	var data models.Post

	err := c.Bind(&data)
	data.UpdateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", -1)
	result, err := data.Update(data.PostId)
	tools.HasError(err, "", -1)
	app.OK(c, result, "Modified successfully")
}

// @Summary delete job
// @Description delete data
// @Tags position
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "message": "Successfully deleted"}"
// @Success 200 {string} string "{"code": -1, "message": "Deletion failed"}"
// @Router /api/v1/post/{postId} [delete]
func DeletePost(c *gin.Context) {
	var data models.Post
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("postId", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "Deletion failed", 500)
	app.OK(c, result, "Delete successfully")
}
