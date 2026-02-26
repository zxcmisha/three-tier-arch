package transport

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	server *http.Server // Добавляем поле для доступа к методу Shutdown
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(httpHandler.HandleCreateTask)
	router.Path("/tasks/{id}").Methods("GET").HandlerFunc(httpHandler.HandleGetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(httpHandler.HandleGetAllUncompletedTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(httpHandler.HandleGetAllTasks)
	router.Path("/tasks/{id}").Methods("PATCH").HandlerFunc(httpHandler.HandleCompleteTask)
	router.Path("/tasks/{id}").Methods("DELETE").HandlerFunc(httpHandler.HandleDeleteTask)

	return &HTTPServer{
		server: &http.Server{
			Addr:    ":9091",
			Handler: router,
		},
	}
}

// StartServer теперь вызывает ListenAndServe у конкретного объекта сервера
func (s *HTTPServer) StartServer() error {
	return s.server.ListenAndServe()
}

// Stop позволяет корректно завершить работу сервера
func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
