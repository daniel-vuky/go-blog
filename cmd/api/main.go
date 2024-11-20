package main

import (
	"context"
	"github.com/daniel-vuky/go-blog/internal/delivery/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGINT,
	syscall.SIGTERM,
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	errGroup, ctx := errgroup.WithContext(ctx)
	server, err := gin.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	err = server.Start(ctx, errGroup)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	err = errGroup.Wait()
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
