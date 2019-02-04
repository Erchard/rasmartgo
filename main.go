package main

import (
	"github.com/Erchard/rasmartgo/controller"
	"log"
)

func main() {
	log.Println("Start Rasmart blockchain explorer on http://localhost:8080/")
	controller.ServerStart()
	log.Println("End")

}
