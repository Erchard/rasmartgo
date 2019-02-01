package rsapi

import (
	"encoding/hex"
	"log"
	"net"
)

var connect net.Conn = nil

func Request(raw_string string, arr bool) ([]byte, []byte) {

	terminate := "0000" + "00000000"

	log.Printf("Request: %s", raw_string)

	request, err := hex.DecodeString(raw_string + terminate)
	if err != nil {
		log.Fatal("Parse pubkey error")
		log.Fatal(err)
	}
	if connect == nil {
		conn, err := net.Dial("tcp", "95.84.138.232:38101")
		defer conn.Close()
		if err != nil {
			log.Fatal("Connection error")
			log.Fatal(err)
		}
		connect = conn

		pub_key := "50f86b12dbdb50ae9197980787198e278dc9ec94ec8491e3b79df03157ad0bd1"
		get_info := "0100" + "20000000" + pub_key

		Request(get_info, false)
	}

	log.Println("Connected!")
	len_n, err := connect.Write(request)
	if err != nil {
		log.Fatal("Send request error")
		log.Fatal(err)
	}
	log.Println("conn len: ", len_n)

	buff := make([]byte, 1024)
	n, err := connect.Read(buff)
	if err != nil {
		log.Fatal("Responce error")
		log.Fatal(err)
	}
	log.Printf("Receive: %x", buff[:n])

	if arr {
		arrbuff := make([]byte, 1024)
		n, err := connect.Read(arrbuff)
		if err != nil {
			log.Fatal("Responce error")
			log.Fatal(err)
		}
		log.Printf("Receive: %x", arrbuff[:n])
		return buff[:n], arrbuff[:n]
	}
	return buff[:n], nil
}
