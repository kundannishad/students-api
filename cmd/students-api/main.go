package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kundannishad/students-api/internal/config"
	"github.com/kundannishad/students-api/internal/http/handler/student"
)

func main() {
	//fmt.Println("Welcome to student api")

	//Load Confg
	//Database Setup
	//setup router

	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Addr))
	fmt.Printf("Server started %s:", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Faild To start server", slog.String("error", err.Error()))
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("Faild to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server Shutdown successfully !")

}
