package main

func initializeRoutes() {

	router.GET("/", showIndexPage)

	router.GET("/settings", showSettingsPage)
	router.POST("/settings", changeSettings)

}
