package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guruorgoru/buffery/internal/config"
	"github.com/guruorgoru/buffery/internal/db"
	"github.com/guruorgoru/buffery/internal/router"
)

func main() {
        port, err := config.GetPort()
        if err != nil {
                log.Fatalln("Error Loading port env:", err)
        }
        host, err := config.GetHost()
        if err != nil {
                log.Fatalln("Error Loading host env:", err)
        }
	dbURL, err := config.GetDbURL()
	if err != nil {
		log.Fatalln("Error loading db_url env:", err)
	}
	gormDB, err := db.Init(dbURL)
	if err != nil {
		log.Fatalln("Error while initializing database:", err)
	}

        router := router.GetServer(gormDB)
        server := &http.Server{
                Addr:    net.JoinHostPort(host, port),
                Handler: router,
        }
        done := make(chan os.Signal, 1)
        signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
        go func() {
                if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
                        log.Fatalln("Error Starting The Server:", err)
                }
        }()
        <-done
        log.Println("Shutting down Server gracefully...")

        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()

        if err := server.Shutdown(ctx); err != nil {
                log.Fatalf("Server Shutdown Failed:%+v", err)
        }
        log.Println("Server exited properly")
}

