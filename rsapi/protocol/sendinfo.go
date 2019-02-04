package protocol

import "log"

func SendInfo(args []byte) {
	log.Printf("Receive: %x \n", args)
}
