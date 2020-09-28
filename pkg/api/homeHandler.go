package api

import (
	"fmt"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>METRON API IS RUNNING</h1>")
}
