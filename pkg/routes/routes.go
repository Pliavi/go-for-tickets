package routes

import (
	"net/http"

	"github.com/pliavi/go-for-tickets/pkg/controllers"
	"github.com/pliavi/go-for-tickets/pkg/database"
	"github.com/pliavi/go-for-tickets/pkg/repositories"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	db := database.MustGetConnection()

	concertRepo := repositories.NewConcertRepository(db)
	// customerRepo := repositories.NewCustomerRepository(db)

	// customerService := services.NewCustomerService(customerRepo)

	concertController := controllers.NewConcertController(concertRepo)

	mux.HandleFunc("/concerts", concertController.Index)
	mux.HandleFunc("/concerts/show", concertController.Show)

	return mux
}
