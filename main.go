package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", home)

	http.ListenAndServe(":8089", router)
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Rizvan!!!"))
}
