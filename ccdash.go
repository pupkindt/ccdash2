package main

import (
	"context"
	//"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	//"github.com/spf13/viper"

	//"testing"
	//"fmt"
	//ioutil "io/ioutil"
	//"crypto/x509"
	//"crypto/tls"
	//"github.com/kat-co/vala"
	"fmt"
)

const (

	TEMPERATURE_MAX_VALID = 27
	TEMPERATURE_MIN_VALID = 16
	TEMPERATURE_DEFAULT = 21

)
var logFile = logrus.New()

var router *gin.Engine

var useTemplates bool = true
var useLocalBootstrap = true
var targetTemperature int = 21

func main() {

	// for the first open or create log
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("ccdash.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
	  logFile.Out = file
	} else {
	  logFile.Errorln("Failed to log to file, using default stderr")
	}

	// then load config

	ConfigLoad()

	targetTemperature = CCConfig.TargetTemperature

	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	if useTemplates {

		router.LoadHTMLGlob("templates/*.html")
		router.Static("/templates/static", "./templates/static")

		initializeRoutes()

	} else {

		router.GET("/", func(c *gin.Context) {
			//time.Sleep(5 * time.Second)
			c.String(http.StatusOK, "Welcome CCDash Server")
		})

	}

	strPort := fmt.Sprintf(":%v", CCConfig.Port)

	srv := &http.Server{
		Addr:    strPort,
		Handler: router,
	}

	logFile.Infof("Start CCDash Server! Listen on port%s", strPort)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			logFile.Errorf("listen error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// kill -INT $( pgrep CCDash_run )
	// pkill -INT CCDash_run

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	logFile.Infoln("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logFile.Errorf("Server Shutdown error: %v", err)
	}

	logFile.Infoln("Server exit")

}

/*
func initAppConfig() error {

	var err error


	viper.SetConfigType("toml")
	viper.SetConfigName("ccdash")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()



	return err

}


func setAppConfigDefaults() {

	//viper.AutomaticEnv()

	if !viper.IsSet("climate.targetTemperature") || viper.GetString("climate.targetTemperature") == "" {
		viper.Set("climate.targetTemperature", TEMPERATURE_DEFAULT)
	}

	if viper.GetString("server.port") == "" {
		viper.Set("server.port", 8080)
	}

	if viper.GetString("server.mode") == "" {
		viper.Set("server.mode", "debug")
	}

	viper.SaveConfig()

}

func validateAppConfig() error {
	// /-*
	vala.BeginValidation().Validate(
		vala.StringNotEmpty(viper.GetString("postgres.user"), "postgres.user"),
		vala.StringNotEmpty(viper.GetString("postgres.host"), "postgres.host"),
		vala.Not(vala.Equals(viper.GetInt("postgres.port"), 0, "postgres.port")),
		vala.StringNotEmpty(viper.GetString("postgres.dbname"), "postgres.dbname"),
		vala.StringNotEmpty(viper.GetString("postgres.sslmode"), "postgres.sslmode"),
	).Check()

	//-/

	return vala.BeginValidation().Validate(
		vala.StringNotEmpty(viper.GetString("climate.targetTemperature"), "climate.targetTemperature"),
		vala.Not(vala.GreaterThan(viper.GetInt("climate.targetTemperature"), TEMPERATURE_MAX_VALID, "climate.targetTemperature") ),
		vala.Not(vala.GreaterThan(TEMPERATURE_MIN_VALID, viper.GetInt(  "climate.targetTemperature"), "climate.targetTemperature") ),
	).Check()
}

*/