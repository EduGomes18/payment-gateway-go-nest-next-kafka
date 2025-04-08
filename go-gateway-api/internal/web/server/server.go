package server

import (
	"go-gateway-api/internal/service"
	"go-gateway-api/internal/web/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	accountService *service.AccountService
	port string
}

func NewServer(port string, accountService *service.AccountService) *Server {
	return &Server{
		router: chi.NewRouter(),
		port: port,
		accountService: accountService,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handler.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
