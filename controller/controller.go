package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Erchard/rasmartgo/counters"
	"log"
	"net/http"
)

func ServerStart() {

	http.HandleFunc("/counters", countersHandler)
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>RASMART</title>
</head>
<body>
<h1>Hello Rasmart!</h1>
</body>
</html>
`)
}

func countersHandler(writer http.ResponseWriter, request *http.Request) {
	raw := counters.GetCounters()
	json, _ := json.Marshal(raw)
	fmt.Fprintf(writer, "%s", json)
}
