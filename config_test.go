package main

import (
	"testing"
	"fmt"
)

func TestConfigLoad(t *testing.T) {

	ConfigLoad()

	fmt.Printf("server port: %v\n", CCConfig.Port)
	fmt.Printf("server mode: %v\n", CCConfig.Mode)

	fmt.Printf("target temperature: %v\n", CCConfig.TargetTemperature)

	CCConfig.TargetTemperature = 23


	temperature, err := ConfigFile.Section("application").Key("target temperature").Int()
	if err != nil {
		fmt.Printf("%v", err)
	}


	fmt.Printf("new target temperature: %v\n", temperature)

	ConfigSave()


}