package main

import app "github.com/tuxoo/idler/internal/app/idler-service"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
