package main

import (
	"nest/app/config"
)

func main() {
	app := config.NewApp()
	app.Run()
}
