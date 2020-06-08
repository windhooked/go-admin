package api

import (
	"context"
	"fmt"
	"go-admin/database"
	"go-admin/router"
	"go-admin/tools"
	config2 "go-admin/tools/config"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "go-admin server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")

	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFieldName = "created"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorFieldName = "message"
	//zerolog.ErrorStackMarshaler = MarshalStack
	log.Logger = log.With().Caller().Logger()

	if os.Getenv("APP_ENV") == "" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05",
			//MarshalIndent: true,
		})
	}
}

func usage() {
	usageStr := `starting api server`
	fmt.Printf("%s\n", usageStr)
}

func setup() {

	//1. Read configuration
	config2.ConfigSetup(config)
	//2. Set log
	tools.InitLogger()
	//3. Initialize the database link
	database.Setup()

}

func run() error {
	if mode != "" {
		config2.SetConfig(config, "settings.application.mode", mode)
	}
	if viper.GetString("settings.application.mode") == string(tools.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()

	defer database.Eloquent.Close()
	if port != "" {
		config2.SetConfig(config, "settings.application.port", port)
	}

	srv := &http.Server{
		Addr:    config2.ApplicationConfig.Host + ":" + config2.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		// Service connection
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()
	content, _ := ioutil.ReadFile("./static/go-admin.txt")
	log.Info().Msgf("%v", string(content))
	log.Info().Msgf("Server Run :%s/", config2.ApplicationConfig.Port)
	log.Info().Msgf("Swagger URL :%s  /swagger/index.html", config2.ApplicationConfig.Port)

	log.Info().Msgf("Enter Control + C Shutdown Server")
	// Wait for an interrupt signal to gracefully shut down the server (set a timeout of 5 seconds)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msgf("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server Shutdown:", err)
	}
	log.Info().Msgf("Server exiting")
	return nil
}
