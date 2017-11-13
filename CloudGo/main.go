package main

import "os"

import "github.com/echobeee/Service-Computing/tree/master/CloudGo/Service"
import flag "github.com/spf13/pflag"


const (
	PORT string = "3000"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	// Define server port
	flag.StringVarP(&port, "port", "p", "3000", "define server port")
	flag.Parse()

	server := service.NewServer()
	server.Run(":" + port)
}
