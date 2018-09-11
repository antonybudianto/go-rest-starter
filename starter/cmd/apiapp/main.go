package main

import (
	"log"
	"starter/app"
)

func main() {
	a := app.App{}
	a.Initialize("root", "hello", "rest_api_example")

	log.Println("API running at :8000, try http://localhost:8000/users")

	a.Run(":8000")
}
