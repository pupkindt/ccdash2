package main

import (
	//"math/rand"
	//"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sirupsen/logrus"
)

func showSettingsPage(c *gin.Context) {

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"settings.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Settings Page",
			"TargetTemperature" : targetTemperature,
		},
	)

}

func changeSettings(c *gin.Context) {

	var err error
	var newTargetTemperature int

	// Obtain the POSTed targetTemperature
	targetTemperatureString := c.PostForm("target-temperature")

	newTargetTemperature, err = strconv.Atoi(targetTemperatureString)

	if err != nil {

		c.HTML(http.StatusBadRequest, "settings.html", gin.H{
			"ErrorTitle":   "Invalid temperature value!",
			"ErrorMessage": "Can't convert to integer!",
			"TargetTemperature" : targetTemperatureString })

		logFile.WithFields(logrus.Fields{
			"targetTemperature" : targetTemperatureString,
		}).Error("Can't convert to integer!")

	} else {

		if newTargetTemperature < 17 || newTargetTemperature > 25 {

			c.HTML(http.StatusBadRequest, "settings.html", gin.H{
				"ErrorTitle":   "Invalid temperature value!",
				"ErrorMessage": "targetTemperature < 17 || targetTemperature > 25!",
				"TargetTemperature" : targetTemperatureString })

			logFile.WithFields(logrus.Fields{
				"targetTemperature" : targetTemperatureString,
			}).Error("targetTemperature < 17 || targetTemperature > 25!")

		} else {

			targetTemperature = newTargetTemperature

			logFile.WithFields(logrus.Fields{
				"targetTemperature" : targetTemperatureString,
			}).Info("OK!")

			//
			ConfigSetTargetTemperature(targetTemperature)

			showSettingsPage(c)
		}
	}

}
