package counters

import (
	json "encoding/json"
	"log"
	"testing"
)

func TestGetCounters(t *testing.T) {

	raw := GetCounters()
	json, _ := json.Marshal(raw)
	log.Printf("Counters: %s \n", json)

}
