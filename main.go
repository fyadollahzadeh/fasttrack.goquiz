// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name userId

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"goquiz/api"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "goquiz/docs"
	"goquiz/middleware"
)

func main() {
	r := mux.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
		handlers.AllowCredentials(),
	)

	r.Use(mux.CORSMethodMiddleware(r))

	api.FillDatabase(nil, nil)

	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	protectedRoutes.HandleFunc("/get-quiz/{id}", api.GetQuizByID).Methods("GET", "OPTIONS")
	protectedRoutes.HandleFunc("/submit-answers/{id}", api.SubmitAnswers).Methods("POST", "OPTIONS")
	protectedRoutes.HandleFunc("/get-results/{id}", api.GetResult).Methods("GET", "OPTIONS")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusFound)
	})

	log.Println("Server is running on port 8080 🚀")
	err := http.ListenAndServe(":8080", corsHandler(r))
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
