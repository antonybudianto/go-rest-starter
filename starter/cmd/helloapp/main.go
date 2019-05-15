package main

import (
	"fmt"
	"log"
	"starter/app"
)

const port = 8000

func main() {
	a := app.App{}
	a.Initialize("root", "hello", "rest_api_example")
	log.Println(fmt.Sprintf("API running at :%v, try http://localhost:%v/users", port, port))
	a.Run(fmt.Sprintf(":%v", port))
}
