package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"restapi/internal/features/feature1/repository"
	"restapi/internal/features/feature1/service"
	"restapi/internal/features/feature1/transport"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conn, err := repository.Connection(ctx)
	if err != nil {
		panic(err)
	}

	postgres := repository.NewPostgres(ctx, conn)

	service := service.NewUserService(postgres)
	handlers := transport.NewHTTPHandlers(service)
	httpServer := transport.NewHTTPServer(handlers)

	go func() {
		if err := httpServer.StartServer(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("failed to start http server: %v\n", err)
			stop()
		}
	}()

	fmt.Println("App is running...")

	<-ctx.Done()
	fmt.Println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Stop(shutdownCtx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("App stop correctly!")
}
