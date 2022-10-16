package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bayudava29/go-clean-arc/app"
	"github.com/bayudava29/go-clean-arc/config"
)

func main() {
	// Init configuration
	config.InitConfig()

	// Init Hbase connection
	hbaseConn, errHbase := config.ConnectHbase()
	if errHbase != nil {
		log.Print(errHbase)
	} else {
		defer func() {
			if len(hbaseConn) > 0 {
				for i := range hbaseConn {
					hbaseConn[i].DB.Close()
				}
			}
		}()
	}

	// Init Router
	router := app.InitRouter(hbaseConn)
	log.Print("Route Initialized")

	port := config.CONFIG["PORT"]
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Print("Server Initialized")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
	log.Print("Server exiting")
}
