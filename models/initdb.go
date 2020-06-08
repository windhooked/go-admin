package models

import (
	"fmt"
	orm "go-admin/database"
	config2 "go-admin/tools/config"
	"io/ioutil"
	"strings"
)

func InitDb() error {
	//filePath := "config/db.sql"
	filePath := "db.sql"
	if config2.DatabaseConfig.Dbtype == "sqlite3" {
		fmt.Println("sqlite3 database does not need to be initialized!")
		return nil
	}
	sql, err := Ioutil(filePath)
	if err != nil {
		fmt.Println("database basic data initialization script failed ", err.Error())
		return err
	}
	sqlList := strings.Split(sql, ";")
	for i := 0; i < len(sqlList)-1; i++ {
		if strings.Contains(sqlList[i], "--") {
			fmt.Println(sqlList[i])
			continue
		}
		sql := strings.Replace(sqlList[i]+";", "\n", "", 0)
		if err = orm.Eloquent.Exec(sql).Error; err != nil {
			if !strings.Contains(err.Error(), "Query was empty") {
				return err
			}
		}
	}
	return nil
}

func Ioutil(name string) (string, error) {
	if contents, err := ioutil.ReadFile(name); err == nil {
		//Because contents is []byte type, there will be one more space after directly converted to string type, you need to replace newline with strings.
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("Use ioutil.ReadFile to read a file:", result)
		return result, nil
	} else {
		return "", err
	}
}
