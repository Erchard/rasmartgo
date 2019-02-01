package rsapi

import (
	"encoding/hex"
	"log"
	"net"
)

func main() {
	pub_key := "50f86b12dbdb50ae9197980787198e278dc9ec94ec8491e3b79df03157ad0bd1"
	get_info := "0100" + "20000000" + pub_key

	get_counters := "0500" + "00000000"

	get_blocks := "0900" + "0a000000" + "0000000000000000" + "0100"

	get_transactions := "0d00" + "40000000" + "e85d84b9609640419c3fc3652e593e0d859e28eea4aeaeffba9decce5f46ca793d049a4d0c17d76b050a812e1292f3e6b5283b9944ad730924e69fc357b96493" + "0000000000000000" + "0100"

	terminate := "0000" + "00000000"

	commands := [4]string{get_info, get_counters, get_blocks, get_transactions}

	for i, comm := range commands {
		log.Println("i: ", i)
		request, err := hex.DecodeString(comm + terminate)
		if err != nil {
			log.Fatal("Parse pubkey error")
			log.Fatal(err)
		}
		log.Println("Pub key len: ", len(request))

		conn, err := net.Dial("tcp", "95.84.138.232:38101")
		defer conn.Close()
		if err != nil {
			log.Fatal("Connection error")
			log.Fatal(err)
		}

		log.Println("Connected!")
		len_n, err := conn.Write(request)
		if err != nil {
			log.Fatal("Send request error")
			log.Fatal(err)
		}
		log.Println("conn len: ", len_n)

		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			log.Fatal("Responce error")
			log.Fatal(err)
		}
		log.Printf("Receive: %x", buff[:n])

		if i > 1 {

			n, err := conn.Read(buff)
			if err != nil {
				log.Fatal("Responce error")
				log.Fatal(err)
			}
			log.Printf("Receive: %x", buff[:n])

		}

	}
}
