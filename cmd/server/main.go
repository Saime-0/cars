package main

import (
	"carstore/internal/app"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	log.SetPrefix("[carstore] ")
	log.Println("starting...")
	ctx, cancel := context.WithCancel(context.Background())
	go runApplication(ctx)
	gracefulShutdown(cancel)
}

func runApplication(ctx context.Context) {
	err := app.Run(ctx)
	if err != nil {
		log.Fatalf("application: %s", err)
	}
	os.Exit(0)
}

func gracefulShutdown(cancel context.CancelFunc) {
	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	log.Printf("%s signal was received", <-quit)
	var timeout = 2 * time.Second
	log.Printf("after %v seconds, the program will force exit", timeout.Seconds())
	cancel()
	time.Sleep(timeout)
	os.Exit(0)
}
