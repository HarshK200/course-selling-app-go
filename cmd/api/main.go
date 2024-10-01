package main

import (
	"log"
	"os"

	"github.com/harshk200/course-selling-app-go/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := ":" + os.Getenv("PORT")

	app := server.NewApp(PORT)

	app.LoadRoutes()
	if err := app.Listen(); err != nil {
		log.Fatalf("cannot start the server: %v", err)
	}
}
