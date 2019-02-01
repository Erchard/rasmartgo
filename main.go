package main

import (
	"github.com/Erchard/rasmartgo/counters"
	"github.com/Erchard/rasmartgo/rsapi"
	"log"
)

func main() {
	log.Println("Hello Rasmart!")
	counters.GetCounters()
	rsapi.Main()
	log.Println("End")

}
