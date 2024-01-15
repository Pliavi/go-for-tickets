package controllers

import (
	"net/http"

	"github.com/pliavi/go-for-tickets/pkg/repositories"
	"github.com/pliavi/go-for-tickets/pkg/utils"
)

type concertController struct {
	concertRepository repositories.ConcertRepository
}

// ConcertController defines the interface for concert controllers.
// It's a resource controller, for now methods are not implemented.
//
// TODO: Use Chi to handle methods and be able to use url wildcards
type ConcertController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
}

func NewConcertController(concertRepository repositories.ConcertRepository) ConcertController {
	return &concertController{
		concertRepository: concertRepository,
	}
}

func (cc *concertController) Index(w http.ResponseWriter, r *http.Request) {
	searchFilter := r.URL.Query().Get("search")

	// TODO: implement find filters struct instead of use map
	concerts, err := cc.concertRepository.FindAll(
		repositories.ConcertFindAllFilters{
			Search: searchFilter,
		},
	)

	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "something went wrong", err)
		return
	}

	utils.SendJsonResponse(w, http.StatusOK, concerts)
}

func (cc *concertController) Show(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	concert, err := cc.concertRepository.FindById(id)

	if err != nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "concert not found", err)
		return
	}

	utils.SendJsonResponse(w, http.StatusOK, concert)
}
