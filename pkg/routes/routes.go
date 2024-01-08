package routes

import (
	"net/http"

	"github.com/pliavi/go-for-tickets/pkg/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/enqueue", handlers.AddToQueueHandler)
	mux.HandleFunc("/dequeue", handlers.FinishedBuyingHandler)

	return mux
}
