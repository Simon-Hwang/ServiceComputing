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
    port := os.Getenv("PORT")//get custom environment variables
    if len(port) == 0 {
        port = PORT
    }
    flag.StringVar(&port, "p", PORT, "PORT for httpd listening")
    flag.Parse()
    server := service.NewServer() // get server
    server.Run(":" + port) // run the server at specific port
}