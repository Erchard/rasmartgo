package protocol

import "log"

func Terminate(args []byte) {
	if len(args) > 0 {
		log.Fatal("Terminate command with arguments")
	}
}
