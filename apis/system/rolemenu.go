package system

import (
	"fmt"
	"go-admin/models"
	"go-admin/tools/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary RoleMenu list data
// @Description get JSON
// @Tags character menu
// @Param RoleId query string false "RoleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "Sorry no relevant information was found"}"
// @Router /api/v1/rolemenu [get]
// @Security Bearer
func GetRoleMenu(c *gin.Context) {
	var Rm models.RoleMenu
	err := c.ShouldBind(&Rm)
	result, err := Rm.Get()
	var res app.Response
	if err != nil {
		res.Msg = "Sorry no relevant information was found"
		c.JSON(http.StatusOK, res.ReturnError(200))
		return
	}
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

type RoleMenuPost struct {
	RoleId   string
	RoleMenu []models.RoleMenu
}

func InsertRoleMenu(c *gin.Context) {

	var res app.Response
	res.Msg = "Add successfully"
	c.JSON(http.StatusOK, res.ReturnOK())
	return

}

// @Summary delete user menu data
// @Description delete data
// @Tags character menu
// @Param id path string true "id"
// @Param menu_id query string false "menu_id"
// @Success 200 {string} string "{"code": 200, "message": "Successfully deleted"}"
// @Success 200 {string} string "{"code": -1, "message": "Deletion failed"}"
// @Router /api/v1/rolemenu/{id} [delete]
func DeleteRoleMenu(c *gin.Context) {
	var t models.RoleMenu
	id := c.Param("id")
	menuId := c.Request.FormValue("menu_id")
	fmt.Println(menuId)
	_, err := t.Delete(id, menuId)
	if err != nil {
		var res app.Response
		res.Msg = "Deletion failed"
		c.JSON(http.StatusOK, res.ReturnError(200))
		return
	}
	var res app.Response
	res.Msg = "Successfully deleted"
	c.JSON(http.StatusOK, res.ReturnOK())
	return
}
