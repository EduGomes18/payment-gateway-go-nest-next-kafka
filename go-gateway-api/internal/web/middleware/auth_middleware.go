package middleware

import (
	"go-gateway-api/internal/domain"
	"go-gateway-api/internal/service"
	"net/http"
)

type AuthMiddleware struct {
	accountService *service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{accountService: accountService}
}

func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			http.Error(w, "API key is required", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.FindByApiKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			} 
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// we can return more context on the request if needed
		// like account id, account name, etc.
		// we can add it to the request context
		// and use it in the next handler
		// for example:
		// ctx := context.WithValue(r.Context(), "accountId", account.ID)
		// next.ServeHTTP(w, r.WithContext(ctx))

		next.ServeHTTP(w, r)
	})
}
