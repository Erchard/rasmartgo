package main

import (
	"encoding/hex"
	"log"
)

func main() {
	log.Println("Hello Rasmart!")

	const_pubkey, err := hex.DecodeString("50f86b12dbdb50ae9197980787198e278dc9ec94ec8491e3b79df03157ad0bd1")
	if err != nil {
		log.Fatal("Parse pubkey error")
	}
	log.Println("Pub key len: ", len(const_pubkey))
}
