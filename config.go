package main

import (
	_ "fmt"
	//"github.com/go-ozzo/ozzo-config"
	//"fmt"
	//"github.com/go-ozzo/ozzo-config"
	"github.com/go-ini/ini"
	"os"
	"fmt"
	//"log"
)

type Server struct {
	Port int `ini:"port"`
	Mode string `ini:"mode"`
}

type ClimateControl struct {

	TargetTemperature int `ini:"target temperature"`
	TargetTemperatureMin int `ini:"target temperature min"`
	TargetTemperatureMax int `ini:"target temperature max"`
}

type Config struct {

	*Server `ini:"server"`

	*ClimateControl `ini:"application""`

}

var CCConfig *Config
var ConfigFile *ini.File

const (

	CONFIG_FILE_NAME = "ccdash.ini"

	DEFAULT_SERVER_PORT = 8080
	DEFAULT_SERVER_MODE = "debug"

	DEFAULT_CC_TARGET_TEMPERATURE = 21
	DEFAULT_CC_TARGET_TEMPERATURE_MIN = 17
	DEFAULT_CC_TARGET_TEMPERATURE_MAX = 27
)

func configCreateDefault(file string) {

	c := &Config{
		&Server{
			Port: DEFAULT_SERVER_PORT,
			Mode: DEFAULT_SERVER_MODE,
		},
		&ClimateControl{
			TargetTemperature: DEFAULT_CC_TARGET_TEMPERATURE,
			TargetTemperatureMin: DEFAULT_CC_TARGET_TEMPERATURE_MIN,
			TargetTemperatureMax: DEFAULT_CC_TARGET_TEMPERATURE_MAX,
		},
	}

	configFile := ini.Empty()
	var err error

	err = ini.ReflectFrom(configFile, c)
	if err != nil {
		fmt.Println(err)
		logFile.Errorf("error reflect default config file!")
		os.Exit(-1)
	}

	err = configFile.SaveTo(file)
	if err != nil {
		fmt.Println(err)
		logFile.Errorf("error save default config file!")
		os.Exit(-2)
	}


}

func configCheckExistence(file string) bool {

	_, err := os.Stat(file)
	if err == nil {
		logFile.Infof("config file %s found!\n", file)
		return true
	} else if os.IsNotExist(err) {
		logFile.Infof("config file %s not exists. create default!\n", file)
		configCreateDefault(file)
		return true
	} else {
		logFile.Errorf("config file %s stat error: %v\n", file, err)
		os.Exit(-3)
	}
	return false
}

func ConfigLoad() {

	// check existence and create if absent
	if configCheckExistence(CONFIG_FILE_NAME) {
		var err error
		// read to internal structs
		ConfigFile, err = ini.Load(CONFIG_FILE_NAME)
		if err != nil {
			fmt.Println(err)
			logFile.Errorf("error load config file!")
			os.Exit(-5)
		}

		CCConfig = new(Config)

		ConfigFile.MapTo(CCConfig)
	}

}

func ConfigSetTargetTemperature(t int) {

	CCConfig.TargetTemperature = t

	ConfigSave()
}


func ConfigSave() {

	if ConfigFile == nil {
		logFile.Errorf("config file pointer is nil!")
		os.Exit(-6)
	}

	ConfigFile.ReflectFrom(CCConfig)
	ConfigFile.SaveTo(CONFIG_FILE_NAME)

}

