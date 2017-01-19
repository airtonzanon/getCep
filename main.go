package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cep := params["cep"]

	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		log.Fatal(err)
	}

	returnData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(returnData))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/{cep}", handler)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
