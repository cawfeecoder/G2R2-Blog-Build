package main

import (
	"fmt"
	"log"

	"github.com/kabukky/httpscerts"

	"github.com/labstack/echo/engine/fasthttp"
	"github.com/nfrush/G2R2-Blog-Build/routers"
)

func main() {

	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:443")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}
	fmt.Println("Server now running at https://localhost:8443")
	router := routers.InitRoutes()
	router.Run(fasthttp.WithTLS(":8443", "cert.pem", "key.pem"))

}
