package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var srv http.Server

	idleConnClosed := make(chan struct{})
	go func() {
		fmt.Println("Waiting signal-1...")
		sig := make(chan os.Signal, 1)

		signal.Notify(sig, os.Interrupt)
		signal.Notify(sig, syscall.SIGTERM)

		fmt.Println("Waiting signal-2...")
		<-sig
		fmt.Println("Waiting signal-3...")

		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Printf("HTTP server Shutdown: %v", err)
		} else {
			fmt.Printf("HTTP server gracefully Shutdown")
		}
		close(idleConnClosed)
	}()

	fmt.Println("Http serving...")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("HTTP server ListenAndserver: %v", err)
	}
	fmt.Println("Ready to shutdown")

	<-idleConnClosed
}
