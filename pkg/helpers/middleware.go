package helpers

import (
	"net/http"

	"github.com/codebyaadi/rss-agg/internal/database"
)

func MethodMiddleware(method string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
            return
        }
        next(w, r)
    }
}

type authHandler func (http.ResponseWriter, *http.Request, database.User)

func (cfg *) authMiddleware(handler authHandler) http.HandlerFunc {
    
}