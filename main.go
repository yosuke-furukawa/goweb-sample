package main

import (
	"./controllers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/goweb/responders"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	Address string = ":9090"
)

func mapRoutes() {

	tweetsController := new(tweets.TweetsController)
	goweb.MapController(tweetsController)

	goweb.MapStatic("/app", "./static/app")

	/*
	   Catch-all handler for everything that we don't understand
	*/
	goweb.Map(func(c context.Context) error {

		// just return a 404 message
		return goweb.API.Respond(c, 404, nil, []string{"File not found"})

	})

}

func main() {

	apiResponder := responders.NewGowebAPIResponder(goweb.CodecService, goweb.Respond)
	apiResponder.AlwaysEnvelopResponse = false
	goweb.API = apiResponder
	// map the routes
	mapRoutes()

	log.Print("Starting Goweb powered server...")

	// make a http server using the goweb.DefaultHttpHandler()
	s := &http.Server{
		Addr:           Address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	go func() {
		for _ = range c {
			log.Print("Stopping the server...")
			listener.Close()

			log.Print("Tearing down...")
			log.Fatal("Finished - bye bye.  ;-)")

		}
	}()

	log.Fatalf("Error in Serve: %s", s.Serve(listener))
}
