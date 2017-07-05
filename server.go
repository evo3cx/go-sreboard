package main

import (
	"flag"
	"log"
	"net/http"
	"sre-onboard-golang/bukalapak"

	"github.com/julienschmidt/httprouter"
)

func main() {
	flag.Parse()
	addr := ":8080"
	route := httprouter.New()
	route.PanicHandler = InternalServerErrorHandler
	route.GET("/", bukalapak.HandleHome)
	route.GET("/:command", bukalapak.HandleService)

	log.Printf("versi 2 server start at %s", addr)
	log.Printf("\naddress local server start at http://127.0.0.1%s", addr)
	log.Fatal(http.ListenAndServe(addr, route))
}

//InternalServerErrorHandler ...
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request, param interface{}) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	text := "Internal Server Error, please contact our support "
	w.Write([]byte(text))
}
