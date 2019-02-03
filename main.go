package main

import (
	"github.com/Erchard/rasmartgo/controller"
	"log"
)

func main() {
	log.Println("Hello Rasmart!")
	controller.ServerStart()
	log.Println("End")

}
