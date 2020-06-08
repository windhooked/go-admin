package main

import (
	"go-admin/cmd"
)

// @title go-admin API
// @version 0.0.1
// @description The interface documentation of the front and back end separate permission management system based on Gin + Vue + Element UI
// @description Add qq group: 74520518 Enter the technical exchange group Please note, thank you!
// @license.name MIT
// @license.url https://github.com/wenjianzhang/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

//func main() {
//	configName := "settings"
//
//
//	config.InitConfig(configName)
//
//	gin.SetMode(gin.DebugMode)
//	log.Println(config.DatabaseConfig.Port)
//
//	err := gorm.AutoMigrate(orm.Eloquent)
//	if err != nil {
// log.Fatalln("Database initialization failed err: %v", err)
//	}
//
//	if config.ApplicationConfig.IsInit {
//		if err := models.InitDb(); err != nil {
// log.Fatal("Initialization of database basic data failed!")
//		} else {
//			config.SetApplicationIsInit()
//		}
//	}
//
//	r := router.InitRouter()
//
//	defer orm.Eloquent.Close()
//
//	srv := &http.Server{
//		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
//		Handler: r,
//	}
//
//	go func() {
// // Service connection
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.Fatalf("listen: %s\n", err)
//		}
//	}()
//	log.Println("Server Run ", config.ApplicationConfig.Host+":"+config.ApplicationConfig.Port)
//	log.Println("Enter Control + C Shutdown Server")
// // Wait for an interrupt signal to gracefully shut down the server (set a timeout of 5 seconds)
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//	<-quit
//	log.Println("Shutdown Server ...")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatal("Server Shutdown:", err)
//	}
//	log.Println("Server exiting")
//}

func main() {
	cmd.Execute()
}
