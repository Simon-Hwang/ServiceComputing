package main

import (
    "os"
    "cloud_go/service"
	"flag"
)

const (
    PORT string = "8080"
)

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }

    flag.StringVar(&port, "p", PORT, "PORT for httpd listening")
    flag.Parse()
    server := service.NewServer()
    server.Run(":" + port)
}