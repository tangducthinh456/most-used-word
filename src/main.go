package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"most-used-word/src/config"
	"most-used-word/src/server/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	fmt.Println("Start service!")


	// initial server config
	serverConf := config.GetServerConfig()


	// initial http server
	srv := &http.Server{
		Addr:    ":" + serverConf.ServerPort,
		Handler: router.Router(),
	}

	// This is implement of server with graceful shutdown which 
	// can continue all demain job until service is shutdown when 
	// it receive a cancel signal like Ctrl + C from keyboard


	// start a server in other go routine so it can not block graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// create channel for send cancel signal 
	quit := make(chan os.Signal)


	// get signal from keyboard
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	
	// send signal to go routine running server
	<-quit
	log.Println("Shutting down server...")


	// add context to wait for server to finish all current job
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
