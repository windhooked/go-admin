package migrate

import (
	"fmt"
	"go-admin/database"
	orm "go-admin/database"
	"go-admin/models"
	"go-admin/models/gorm"
	"go-admin/tools"
	config2 "go-admin/tools/config"
	"log"

	"github.com/spf13/cobra"
)

var (
	config   string
	mode     string
	StartCmd = &cobra.Command{
		Use:   "init",
		Short: "initialize the database",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func run() {
	usage := `start init`
	fmt.Println(usage)
	//1. read config
	config2.ConfigSetup(config)
	//2. setup logs
	tools.InitLogger()
	//3. init db
	database.Setup()
	//4. migrate
	_ = migrateModel()
	log.Println("The database structure is initialized successfully!")
	//5. data init complete
	if err := models.InitDb(); err != nil {
		log.Fatal("Database basic data initialization failed!")
	}

	usage = `The database basic data is initialized successfully`
	fmt.Println(usage)
}

func migrateModel() error {
	if config2.DatabaseConfig.Dbtype == "mysql" {
		orm.Eloquent = orm.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return gorm.AutoMigrate(orm.Eloquent)
}
