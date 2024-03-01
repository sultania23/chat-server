package main

import app "github.com/sultania23/chat-server/internal/app/idler-service"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
