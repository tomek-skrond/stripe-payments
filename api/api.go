package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tomek-skrond/stripe-tests/lib/httpwrap"
)

type APIServer struct {
	ListenPort string
}

func NewAPIServer(listenPort string) *APIServer {
	return &APIServer{
		ListenPort: listenPort,
	}
}

func (s *APIServer) Run() {
	r := chi.NewRouter()
	r.Use(JSONLogger)
	r.Use(corsMiddleware)
	// r.Use(middleware.Logger)

	r.Get("/", httpwrap.MakeHTTPHandler(hello))

	r.Post("/api/payment_intents", httpwrap.MakeHTTPHandler(paymentIntents))

	log.Printf("Listening on port %v\n", s.ListenPort)
	log.Fatalln(http.ListenAndServe(s.ListenPort, r))
}
