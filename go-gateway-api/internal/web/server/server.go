package server

import (
	"go-gateway-api/internal/service"
	"go-gateway-api/internal/web/handler"
	"go-gateway-api/internal/web/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port string
}

func NewServer(port string, accountService *service.AccountService, invoiceService *service.InvoiceService) *Server {
	return &Server{
		router: chi.NewRouter(),
		port: port,
		accountService: accountService,
		invoiceService: invoiceService,
	}
}

func (s *Server) ConfigureRoutes() {
	// Add logging middleware
	logMiddleware := middleware.NewLogMiddleware()
	s.router.Use(logMiddleware.Handle)

	accountHandler := handler.NewAccountHandler(s.accountService)
	invoiceHandler := handler.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	// we can group the routes that need authentication
	// and apply the middleware to them
	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Handle)
		r.Post("/invoices", invoiceHandler.CreateInvoice)
		r.Get("/invoices", invoiceHandler.ListInvoicesByAccountApiKey)
		r.Get("/invoices/{id}", invoiceHandler.GetInvoiceById)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
