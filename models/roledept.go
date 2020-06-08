package models

import (
	"fmt"
	orm "go-admin/database"
)

//sys_role_dept
type SysRoleDept struct {
	RoleId int `gorm:"type:int(11)"`
	DeptId int `gorm:"type:int(11)"`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}

func (rm *SysRoleDept) Insert(roleId int, deptIds []int) (bool, error) {
	//ORM does not support batch insertion, so you need to splice sql string
	sql := "INSERT INTO `sys_role_dept` (`role_id`,`dept_id`) VALUES "

	for i := 0; i < len(deptIds); i++ {
		if len(deptIds)-1 == i {
			//The last piece of data ends with a semicolon
			sql += fmt.Sprintf("(%d,%d);", roleId, deptIds[i])
		} else {
			sql += fmt.Sprintf("(%d,%d),", roleId, deptIds[i])
		}
	}
	orm.Eloquent.Exec(sql)

	return true, nil
}

func (rm *SysRoleDept) DeleteRoleDept(roleId int) (bool, error) {
	if err := orm.Eloquent.Table("sys_role_dept").Where("role_id = ?", roleId).Delete(&rm).Error; err != nil {
		return false, err
	}
	var role SysRole
	if err := orm.Eloquent.Table("sys_role").Where("role_id = ?", roleId).First(&role).Error; err != nil {
		return false, err
	}

	return true, nil

}
