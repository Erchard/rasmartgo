package rsapi

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/Erchard/rasmartgo/rsapi/protocol"
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

func responceProcess() {

	head := make([]byte, 6)
	n, err := connect.Read(head)
	if n < 6 || err != nil {
		log.Fatal("Error responce n = ", n)
	}
	argumentsLen := binary.LittleEndian.Uint32(head[2:6])

	args := make([]byte, argumentsLen)

	buff := make([]byte, 1024)
	var offset uint32 = 0
	for offset < argumentsLen {
		n, err := connect.Read(buff)
		if err != nil {
			log.Fatal("Error responce n = ", n)
		}
		newOffset := offset + uint32(n)
		copy(args[offset:newOffset], buff[:n])
		offset = newOffset
	}

	switch comm := binary.LittleEndian.Uint16(head[:2]); comm {
	case 0:
		protocol.Terminate(args)
	case 2:
		protocol.SendInfo(args)
	case 6:
		protocol.SendCounters(args)
	default:
		log.Fatal("Unknown command from srver")
	}

}
