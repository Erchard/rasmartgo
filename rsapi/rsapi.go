package rsapi

import (
	"bytes"
	"encoding/hex"
	"log"
	"net"
)

var connect net.Conn = nil

func Request(raw_string string) []byte {

	terminate := "0000" + "00000000"

	request, err := hex.DecodeString(raw_string + terminate)

	if err != nil {
		log.Fatal("Parse pubkey error")
		log.Fatal(err)
	}
	if connect == nil {
		conn, err := net.Dial("tcp", "95.84.138.232:38101")
		//defer conn.Close()
		if err != nil {
			log.Fatal("Connection error")
			log.Fatal(err)
		}
		connect = conn

		pub_key := "50f86b12dbdb50ae9197980787198e278dc9ec94ec8491e3b79df03157ad0bd1"
		get_info := "0100" + "20000000" + pub_key

		Request(get_info)
		log.Println("Connected!")
	}

	len_n, err := connect.Write(request)
	if err != nil {
		log.Fatal("Send request error")
		log.Fatal(err)
	}
	log.Printf("Request %d: %x", len_n, request)

	//var result []byte
	//t, _ := hex.DecodeString(terminate)
	//for result == nil || len(result) < 6 ||
	//	!bytes.Equal(result[len(result)-6:], t) {
	//	buff := read(connect)
	//	result = append(result, buff...)
	//}
	//return result

	var result []byte
	t, _ := hex.DecodeString(terminate)
	var buff []byte
	for !bytes.Equal(buff, t) {
		buff = read(connect)
		result = append(result, buff...)
	}
	return result
}

func read(connect net.Conn) []byte {
	buff := make([]byte, 1024)
	n, err := connect.Read(buff)
	if err != nil {
		log.Fatal("Responce error")
		log.Fatal(err)
	}
	log.Printf("Receive: %x", buff[:n])
	return buff[:n]
}
