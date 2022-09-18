package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//create a webserver --> http package
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//* per rispondere con un errore avrei potuto usare
			//* questi comandi
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("opps"))
			//* ma il pachetto http ha il metodo Error che mi permette di
			//* fare tutto in una sola riga di codice
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "hello: %s\n", d)
	})
	http.HandleFunc("/ciao", func(w http.ResponseWriter, r *http.Request) {
		log.Println("ciao amici!!!!!!")
	})
	http.ListenAndServe(":9090", nil)
}

/*
what is a handle func ?
its a convenience function on http package
what handlefunc does is register a function to a path in the default serveMux.
the default servemux is an http handler and everything related to server in go
is an http handler

praticamente le funzione vengono legate a dei path e
in questo modo quando arrivano le chiamate al server il serveMux determina quale funzione deve
essere eseguita.
*/
