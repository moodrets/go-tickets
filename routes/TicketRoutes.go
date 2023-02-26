package routes

import (
	"encoding/json"
	"net/http"

	"github.com/moodrets/go-tickets/repositories"
	"github.com/moodrets/go-tickets/response"
)

func TicketRoutes(router *http.ServeMux) {
	ticketRepository := new(repositories.TicketRepository)

	router.HandleFunc("/tickets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			requestParams := make(map[string]string)
			requestParams["status"] = r.URL.Query().Get("status")
			requestParams["title"] = r.URL.Query().Get("title")
			requestParams["creator_id"] = r.URL.Query().Get("creator_id")
			requestParams["in_work_user_id"] = r.URL.Query().Get("in_work_user_id")

			tickets, err := ticketRepository.GetList(requestParams)

			if len(tickets) > 0 {
				w.WriteHeader(http.StatusOK)
				tJson, _ := json.Marshal(tickets)
				w.Write(tJson)
			}

			if len(tickets) == 0 {
				w.WriteHeader(http.StatusOK)
				errorMessage := response.Response{Message: "Тикеты не найдены"}
				jsonError, _ := json.Marshal(errorMessage)
				w.Write(jsonError)
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				errorMessage := response.Response{Message: "Ошибка сервера"}
				jsonError, _ := json.Marshal(errorMessage)
				w.Write(jsonError)
			}

		case http.MethodPost:
		case http.MethodPut:
		case http.MethodDelete:
		}
	})
}
