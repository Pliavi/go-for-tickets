package routes

import (
	"net/http"

	"github.com/pliavi/go-for-tickets/pkg/database"
	"github.com/pliavi/go-for-tickets/pkg/repositories"
	"github.com/pliavi/go-for-tickets/pkg/services"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	db := database.MustGetConnection()

	concertRepo := repositories.NewConcertRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)

	customerService := services.NewCustomerService(customerRepo)

	return mux
}
