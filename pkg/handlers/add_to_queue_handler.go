package handlers

import (
	"net/http"
)

func AddToQueueHandler(w http.ResponseWriter, r *http.Request) {
	// queueMutex.Lock()
	// defer queueMutex.Unlock()

	// customer := model.NewCustomer(
	// 	nil,
	// 	time.Now(),
	// 	EstimateTime(),
	// )

	// queue = append(queue, *customer)

	// customer_json, err := json.Marshal(customer)

	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// fmt.Fprintln(w, string(customer_json))
}
