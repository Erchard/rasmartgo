package main

import (
	"github.com/Erchard/rasmartgo/counters"
	"log"
)

func main() {
	log.Println("Hello Rasmart!")
	counters.GetCounters()
	log.Println("End")

}
