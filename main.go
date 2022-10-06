package main

import (
	// "fmt"
	// "os"
	"os"
	"products/router"
	// "regexp"
)

func main() {
	r := router.StartRouter()
	port := os.Getenv("PORT")

	r.Run(":"+port)
}
