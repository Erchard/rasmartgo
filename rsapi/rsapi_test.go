package rsapi

import (
	"log"
	"testing"
)

func TestRequest(t *testing.T) {
	get_counters := "0500" + "00000000"

	get_blocks := "0900" + "0a000000" + "0000000000000000" + "0100"

	get_transactions := "0d00" + "4a000000" + "e85d84b9609640419c3fc3652e593e0d859e28eea4aeaeffba9decce5f46ca793d049a4d0c17d76b050a812e1292f3e6b5283b9944ad730924e69fc357b96493" + "0000000000000000" + "ff00"

	commands := [3]string{get_counters, get_blocks, get_transactions}

	for i, comm := range commands {
		log.Println("i: ", i)
		Request(comm, i > 0)

	}
}
