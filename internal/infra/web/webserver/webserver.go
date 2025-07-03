package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) AddRoute(method, path string, handler http.HandlerFunc) {
	s.Handlers[method+" "+path] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)

	// CORS middleware
	s.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			if r.Method == "OPTIONS" {
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	for path, handler := range s.Handlers {
		if len(path) > 4 && (path[:4] == "GET " || path[:5] == "POST ") {
			// Handle method-specific routes
			parts := strings.SplitN(path, " ", 2)
			method := parts[0]
			route := parts[1]

			switch method {
			case "GET":
				s.Router.Get(route, handler)
			case "POST":
				s.Router.Post(route, handler)
			}
		} else {
			// Fallback for old-style routes
			s.Router.Handle(path, handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
