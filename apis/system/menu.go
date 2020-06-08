package system

import (
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary Menu list data
// @Description get JSON
// @Tags menu
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "Sorry no relevant information was found"}"
// @Router /api/v1/menulist [get]
// @Security Bearer
func GetMenuList(c *gin.Context) {
	var Menu models.Menu
	Menu.MenuName = c.Request.FormValue("menuName")
	Menu.Visible = c.Request.FormValue("visible")
	Menu.Title = c.Request.FormValue("title")
	Menu.DataScope = tools.GetUserIdStr(c)
	result, err := Menu.SetMenu()
	tools.HasError(err, "Sorry no relevant information was found", -1)

	app.OK(c, result, "")
}

// @Summary Menu list data
// @Description get JSON
// @Tags menu
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "Sorry no relevant information was found"}"
// @Router /api/v1/menu [get]
// @Security Bearer
func GetMenu(c *gin.Context) {
	var data models.Menu
	id, err := tools.StringToInt(c.Param("id"))
	data.MenuId = id
	result, err := data.GetByMenuId()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	app.OK(c, result, "")
}

func GetMenuTreeRoleselect(c *gin.Context) {
	var Menu models.Menu
	var SysRole models.SysRole
	id, err := tools.StringToInt(c.Param("roleId"))
	SysRole.RoleId = id
	result, err := Menu.SetMenuLable()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = SysRole.GetRoleMeunId()
		tools.HasError(err, "Sorry no relevant information was found", -1)
	}
	app.Custum(c, gin.H{
		"code":        200,
		"menus":       result,
		"checkedKeys": menuIds,
	})
}

// @Summary get the menu tree
// @Description get JSON
// @Tags menu
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/menuTreeselect [get]
// @Security Bearer
func GetMenuTreeelect(c *gin.Context) {
	var data models.Menu
	result, err := data.SetMenuLable()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	app.OK(c, result, "")
}

// @Summary create menu
// @Description get JSON
// @Tags menu
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param menuName formData string true "menuName"
// @Param Path formData string false "Path"
// @Param Action formData string true "Action"
// @Param Permission formData string true "Permission"
// @Param ParentId formData string true "ParentId"
// @Param IsDel formData string true "IsDel"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/menu [post]
// @Security Bearer
func InsertMenu(c *gin.Context) {
	var data models.Menu
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "Sorry no relevant information was found", -1)
	data.CreateBy = tools.GetUserIdStr(c)
	result, err := data.Create()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	app.OK(c, result, "")
}

// @Summary modify the menu
// @Description get JSON
// @Tags menu
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param id path int true "id"
// @Param data body models.Menu true "body"
// @Success 200 {string} string "{"code": 200, "message": "Modified successfully"}"
// @Success 200 {string} string "{"code": -1, "message": "Modification failed"}"
// @Router /api/v1/menu/{id} [put]
// @Security Bearer
func UpdateMenu(c *gin.Context) {
	var data models.Menu
	err2 := c.BindWith(&data, binding.JSON)
	data.UpdateBy = tools.GetUserIdStr(c)
	tools.HasError(err2, "Modification failed", -1)
	_, err := data.Update(data.MenuId)
	tools.HasError(err, "", 501)
	app.OK(c, "", "Modified successfully")

}

// @Summary delete menu
// @Description delete data
// @Tags menu
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "message": "Successfully deleted"}"
// @Success 200 {string} string "{"code": -1, "message": "Deletion failed"}"
// @Router /api/v1/menu/{id} [delete]
func DeleteMenu(c *gin.Context) {
	var data models.Menu
	id, err := tools.StringToInt(c.Param("id"))
	data.UpdateBy = tools.GetUserIdStr(c)
	_, err = data.Delete(id)
	tools.HasError(err, "Deletion failed", 500)
	app.OK(c, "", "Delete successful")
}

// @Summary gets menu list data based on the role name (used in the left menu)
// @Description get JSON
// @Tags menu
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "Sorry no relevant information was found"}"
// @Router /api/v1/menurole [get]
// @Security Bearer
func GetMenuRole(c *gin.Context) {
	var Menu models.Menu
	result, err := Menu.SetMenuRole(tools.GetRoleName(c))
	tools.HasError(err, "Get failed", 500)
	app.OK(c, result, "")
}

// @Summary gets the menu id array corresponding to the role
// @Description get JSON
// @Tags menu
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "Sorry no relevant information was found"}"
// @Router /api/v1/menuids/{id} [get]
// @Security Bearer
func GetMenuIDS(c *gin.Context) {
	var data models.RoleMenu
	data.RoleName = c.GetString("role")
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.GetIDS()
	tools.HasError(err, "Get failed", 500)
	app.OK(c, result, "")
}
