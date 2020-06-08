package system

import (
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

// @Summary list data
// @Description get JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "Sorry no relevant information was found"}"
// @Router /api/v1/sysUserList [get]
// @Security Bearer
func GetSysUserList(c *gin.Context) {
	var data models.SysUser
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

	data.Username = c.Request.FormValue("userName")
	data.Status = c.Request.FormValue("status")
	data.Phone = c.Request.FormValue("phone")

	postId := c.Request.FormValue("postId")
	data.PostId, _ = tools.StringToInt(postId)

	deptId := c.Request.FormValue("deptId")
	data.DeptId, _ = tools.StringToInt(deptId)

	data.DataScope = tools.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

// @Summary get user
// @Description get JSON
// @Tags user
// @Param userId path int true "user code"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser/{userId} [get]
// @Security
func GetSysUser(c *gin.Context) {
	var SysUser models.SysUser
	SysUser.UserId, _ = tools.StringToInt(c.Param("userId"))
	result, err := SysUser.Get()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	var SysRole models.SysRole
	var Post models.Post
	roles, err := SysRole.GetList()
	posts, err := Post.GetList()

	postIds := make([]int, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int, 0)
	roleIds = append(roleIds, result.RoleId)
	app.Custum(c, gin.H{
		"code":    200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
	})
}

// @Summary gets the currently logged in user
// @Description get JSON
// @Tags Personal Center
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security
func GetSysUserProfile(c *gin.Context) {
	var SysUser models.SysUser
	userId := tools.GetUserIdStr(c)
	SysUser.UserId, _ = tools.StringToInt(userId)
	result, err := SysUser.Get()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	var SysRole models.SysRole
	var Post models.Post
	var Dept models.Dept
	//Get the role list
	roles, err := SysRole.GetList()
	//Get the job list
	posts, err := Post.GetList()
	//Get the department list
	Dept.DeptId = result.DeptId
	dept, err := Dept.Get()

	postIds := make([]int, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int, 0)
	roleIds = append(roleIds, result.RoleId)

	app.Custum(c, gin.H{
		"code":    200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
		"dept":    dept,
	})
}

// @Summary Get user roles and positions
// @Description get JSON
// @Tags user
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser [get]
// @Security
func GetSysUserInit(c *gin.Context) {
	var SysRole models.SysRole
	var Post models.Post
	roles, err := SysRole.GetList()
	posts, err := Post.GetList()
	tools.HasError(err, "Sorry no relevant information was found", -1)
	mp := make(map[string]interface{}, 2)
	mp["roles"] = roles
	mp["posts"] = posts
	app.OK(c, mp, "")
}

// @Summary create user
// @Description get JSON
// @Tags user
// @Accept  application/json
// @Product application/json
// @Param data body models.SysUser true "user data"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/sysUser [post]
func InsertSysUser(c *gin.Context) {
	var sysuser models.SysUser
	err := c.BindWith(&sysuser, binding.JSON)
	tools.HasError(err, "Illegal data format", 500)

	sysuser.CreateBy = tools.GetUserIdStr(c)
	id, err := sysuser.Insert()
	tools.HasError(err, "Add failed", 500)
	app.OK(c, id, "added successfully")
}

// @Summary modify user data
// @Description get JSON
// @Tags user
// @Accept  application/json
// @Product application/json
// @Param data body models.SysUser true "body"
// @Success 200 {string} string "{"code": 200, "message": "Modified successfully"}"
// @Success 200 {string} string "{"code": -1, "message": "Modification failed"}"
// @Router /api/v1/sysuser/{userId} [put]
func UpdateSysUser(c *gin.Context) {
	var data models.SysUser
	err := c.Bind(&data)
	tools.HasError(err, "Data parsing failed", -1)
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.UserId)
	tools.HasError(err, "Modification failed", 500)
	app.OK(c, result, "Modified successfully")
}

// @Summary delete user data
// @Description delete data
// @Tags user
// @Param userId path int true "userId"
// @Success 200 {string} string "{"code": 200, "message": "Successfully deleted"}"
// @Success 200 {string} string "{"code": -1, "message": "Deletion failed"}"
// @Router /api/v1/sysuser/{userId} [delete]
func DeleteSysUser(c *gin.Context) {
	var data models.SysUser
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("userId", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "Deletion failed", 500)
	app.OK(c, result, "Delete successfully")
}

// @Summary modify avatar
// @Description get JSON
// @Tags user
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {string} string "{"code": 200, "message": "Add success"}"
// @Success 200 {string} string "{"code": -1, "message": "Add failed"}"
// @Router /api/v1/user/profileAvatar [post]
func InsetSysUserAvatar(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		log.Println(file.Filename)
		// Upload the file to the specified directory
		_ = c.SaveUploadedFile(file, filPath)
	}
	sysuser := models.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	sysuser.Avatar = "/" + filPath
	sysuser.UpdateBy = tools.GetUserIdStr(c)
	sysuser.Update(sysuser.UserId)
	app.OK(c, filPath, "Modified successfully")
}

func SysUserUpdatePwd(c *gin.Context) {
	var pwd models.SysUserPwd
	err := c.Bind(&pwd)
	tools.HasError(err, "Data analysis failed", 500)
	sysuser := models.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	sysuser.SetPwd(pwd)
	app.OK(c, "", "Password changed successfully")
}
