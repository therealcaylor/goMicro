package main

import (
	"context"
	"gomicro/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// we need to create a reference to our hello handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	// now we need to register the handler to the serveMux
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	//* e dato che non vogliamo usare il defaultServeMux allora
	//* dobbiamo specificare che vogliamo usare sm
	// http.ListenAndServe(":9090", sm)
	//* ci sono delle cose come il timeout che dobbiamo gestire
	//* in modo da poter avere una performance migliore
	//* per una cosa piu specifica é meglio creare il proprio server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//* graceful shotdown
	sigChan := make(chan os.Signal)
	//creo un chanel
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	//il chanel verra notificato ogni volta che avviene un interruzione o un kill
	sig := <-sigChan
	l.Println("received terminates, graceful shotdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
