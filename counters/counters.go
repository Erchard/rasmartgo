package counters

import "github.com/Erchard/rasmartgo/rsapi"

type Counters struct {
	Blocks       uint8 `json:"blocks"`
	Transactions uint8 `json:"transactions"`
}

func GetCounters() Counters {
	const get_counters = "0500" + "00000000"
	rsapi.Request(get_counters)
	return Counters{0, 0}
}
