package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

const webPort = "8080"

func main() {
	srv := gin.Default()

	registerRoutes(srv)

	err := srv.Run(fmt.Sprintf(":%v", webPort))
	if err != nil {
		log.Panic(err)
	}

}
