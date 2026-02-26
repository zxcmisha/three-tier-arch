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
	// Добавляем SIGINT (для Ctrl+C) и SIGTERM (для Docker)
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

	// 1. Запускаем сервер в горутине, чтобы он не блокировал main
	go func() {
		if err := httpServer.StartServer(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("failed to start http server: %v\n", err)
			stop() // Отменяем контекст, если сервер не смог запуститься
		}
	}()

	fmt.Println("App is running...")

	// 2. Блокируем main здесь до получения сигнала от Docker
	<-ctx.Done()
	fmt.Println("Shutting down gracefully...")

	// 3. Настраиваем дедлайн для завершения работы (Graceful Shutdown)
	// Даем серверу, например, 5 секунд, чтобы дообработать текущие запросы
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. Вызываем метод остановки вашего сервера
	// Убедитесь, что внутри StartServer используется http.Server,
	// и добавьте метод Stop/Shutdown в структуру httpServer
	if err := httpServer.Stop(shutdownCtx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("App stop correctly!")
}
