package routes

import "net/http"

func UserRoutes(router *http.ServeMux) {
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			return
		}

		if r.Method == "POST" {
			return
		}

		if r.Method == "PUT" {
			return
		}

		if r.Method == "DELETE" {
			return
		}
	})
}
