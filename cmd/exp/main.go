package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {

	http.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
		//ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		//defer cancel()

		res, err := json.Marshal(Response{
			Message: "test",
		})
		if err != nil {
			log.Printf("Ошибка сериализации JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write(res); err != nil {
			log.Printf("Ошибка записи ответа: %v", err)
		}
	})

	log.Println("HTTP сервер слушает :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
