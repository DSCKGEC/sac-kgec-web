package main

import (
	"os"

	"github.com/dsckgec/sac-kgec-web/server"

	"github.com/joho/godotenv"
)

var app = server.Server{}

func main() {
	godotenv.Load()
	Port := os.Getenv("PORT")
	app.Run(Port)
}
