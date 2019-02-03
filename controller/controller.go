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

func countersHandler(writer http.ResponseWriter, request *http.Request) {
	raw := counters.GetCounters()
	json, _ := json.Marshal(raw)
	fmt.Fprintf(writer, "%s", json)
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, `
<!doctype html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css"
          integrity="sha384-GJzZqFGwb1QTTN6wy59ffF1BuGJpLSa9DkKMp0DgiMDm4iYMj70gZWKYbI706tWS" crossorigin="anonymous">

    <title>Hello, world!</title>
</head>
<body>
<div class="container-fluid text-light bg-dark">
    <nav class="navbar navbar-dark bg-dark">
        <a class="navbar-brand" href="#">
            <img src="https://en.rasmart.io/img/logo.png" height="30"
                 class="d-inline-block align-top" alt="">
        </a>
    </nav>
    <div class="container">
        <div class="row">
            <div class="col-sm">
                <table class="table table-sm table-dark">
                    <thead>
                    <tr>
                        <th scope="col">Blocks</th>
                        <th scope="col">Transactions</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td id="blocks_number">0</td>
                        <td id="tx_numbers">0</td>
                    </tr>

                    </tbody>
                </table>
            </div>
            <div class="col-sm">
            </div>
            <div class="col-sm">
            </div>
        </div>
    </div>
    <div class="container">
        <div class="row">
            <div class="col-sm">
                <table class="table table-sm table-dark">
                    <thead>
                    <tr>
                        <th scope="col">Last blocks</th>
                        <th scope="col">Transactions</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>
                    <tr>
                        <td>8c854b1cf465e00749ed2c6dd426b595423171e755dd0584b1df28da765fdb76</td>
                        <td>320</td>
                    </tr>

                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<!-- Optional JavaScript -->
<!-- jQuery first, then Popper.js, then Bootstrap JS -->
<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
        integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo"
        crossorigin="anonymous"></script>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.6/umd/popper.min.js"
        integrity="sha384-wHAiFfRlMFy6i5SRaxvfOCifBUQy1xHdJ/yoi7FRNXMRBu5WHdZYu1hA6ZOblgut"
        crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js"
        integrity="sha384-B0UglyR+jN6CkvvICOB2joaf5I4l3gm9GU6Hc1og6Ls7i6U/mkkaduKaBhlAXv9k"
        crossorigin="anonymous"></script>
<script>
    $.ajax({
        url: "http://localhost:8080/counters", // path to file
        dataType: 'json', // type of file (text, json, xml, etc)
        success: function (data) { // callback for successful completion
            $("#blocks_number").html(data.blocks);
            $("#tx_numbers").html(data.transactions);
        },
        error: function () { // callback if there's an error
            console.log("error");
        }
    });
</script>
</body>
</html>
`)
}
