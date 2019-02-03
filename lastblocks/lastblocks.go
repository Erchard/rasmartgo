package lastblocks

type Block struct {
	Hash         string `json:"hash"`
	Transactions uint16 `json:"transactions"`
}
