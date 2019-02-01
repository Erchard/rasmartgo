package counters

type Counters struct {
	Blocks       uint8 `json:"blocks"`
	Transactions uint8 `json:"transactions"`
}

func GetCounters() Counters {

}
