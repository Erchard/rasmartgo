package counters

import (
	"encoding/binary"
	"github.com/Erchard/rasmartgo/rsapi"
)

type Counters struct {
	Blocks       uint16 `json:"blocks"`
	Transactions uint16 `json:"transactions"`
}

func GetCounters() Counters {

	const get_counters = "0500" + "00000000"
	raw := rsapi.Request(get_counters)

	var bl_num uint16 = binary.LittleEndian.Uint16(raw[6:14])
	var tx_num uint16 = binary.LittleEndian.Uint16(raw[14:22])
	rsapi.Request("0000" + "00000000")
	return Counters{bl_num, tx_num}
}
