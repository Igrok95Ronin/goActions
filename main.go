package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Home)

	http.ListenAndServe(":8080", router)
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("I HOME !?!!!"))
}
