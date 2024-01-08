package handlers

import (
	"net/http"
)

func FinishedBuyingHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:  - get the queue, from taking the concert id from the url
	//        - check if the cookie customer id is the same as the one in the queue
	//        - check if the purchase deadline is set
	// 			  - delete the queue
	//				- check if the customer is in another queue
	//				- if not, delete the customer

}
