package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/mikestefanello/pagoda/pkg/handlers"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/tasks"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// Start a new container.
	c := services.NewContainer()
	defer func() {
		// Gracefully shutdown all services.
		if err := c.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	// Build the router.
	if err := handlers.BuildRouter(c); err != nil {
		log.Fatalf("failed to build the router: %v", err)
	}

	// Register all task queues.
	tasks.Register(c)

	// Start the task runner to execute queued tasks.
	c.Tasks.Start(context.Background())

	// Start the server.
	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.Config.HTTP.Hostname, c.Config.HTTP.Port),
			Handler:      c.Web,
			ReadTimeout:  c.Config.HTTP.ReadTimeout,
			WriteTimeout: c.Config.HTTP.WriteTimeout,
			IdleTimeout:  c.Config.HTTP.IdleTimeout,
		}

		if c.Config.HTTP.TLS.Enabled {
			certs, err := tls.LoadX509KeyPair(c.Config.HTTP.TLS.Certificate, c.Config.HTTP.TLS.Key)
			if err != nil {
				log.Fatalf("cannot load TLS certificate: %v", err)
			}

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{certs},
			}
		}

		if err := c.Web.StartServer(&srv); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the web server and task runner.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
}
