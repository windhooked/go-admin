package database

import (
	"bytes"
	"go-admin/tools/config"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql" //Load mysql
	"github.com/jinzhu/gorm"

	"strconv"
)

var Eloquent *gorm.DB

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

func Setup() {

	DbType = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password

	if DbType != "mysql" {
		log.Error().Msgf("db type unknow")
	}
	var err error

	conn := GetMysqlConnect()

	log.Info().Msgf(conn)

	var db Database
	if DbType == "mysql" {
		db = new(Mysql)
		Eloquent, err = db.Open(DbType, conn)

	} else {
		log.Fatal().Msg("db type unknow")
	}
	if err != nil {
		log.Fatal().Msgf("%s connect error %v", DbType, err)
	} else {
		log.Info().Msgf("%s connect success!", DbType)
	}

	if Eloquent.Error != nil {
		log.Fatal().Msgf("database error %v", Eloquent.Error)
	}

	Eloquent.LogMode(true)
}

func GetMysqlConnect() string {
	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	return conn.String()
}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}

type SqlLite struct {
}

func (*SqlLite) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}
