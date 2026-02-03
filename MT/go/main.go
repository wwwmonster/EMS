package main

import (
	"ems/mt/golang/router"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	emsRouter := router.SetupEMGRouter()
	router.CreateJwtToken("alex")
	emsRouter.Run(":9090")
}
