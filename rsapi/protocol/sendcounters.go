package protocol

import "encoding/binary"

func SendCounters(args []byte) {
	var bl_num uint16 = binary.LittleEndian.Uint16(args[0:8])
	var tx_num uint16 = binary.LittleEndian.Uint16(args[8:16])

}
