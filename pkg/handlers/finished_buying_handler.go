package handlers

import "net/http"

func FinishedBuyingHandler(w http.ResponseWriter, r *http.Request) {
	// queueMutex.Lock()
	// defer queueMutex.Unlock()

	// customerId := r.URL.Query().Get("id")

	// for i, customer := range buying {
	// 	if customer.Id.String() == customerId {
	// 		buying = append(buying[:i], buying[i+1:]...)
	// 		break
	// 	}
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
