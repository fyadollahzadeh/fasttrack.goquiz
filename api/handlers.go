
package api

import (
	"encoding/json"
	"goquiz/services"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

type UserAnswers struct {
	Answers []string `json:"answers"`
}

type UserQuizResult struct {
	Score int `json:"score"`
	Percentage int `json:"percentage"`
	Rank int `json:"rank"`
	NumberOfUsers int `json:"numberOfUsers"`
	CustomizedMessage string `json:"customizedMessage"`
}

//@Summary Get a quiz by ID
//@Description Get a quiz by ID
//@Security ApiKeyAuth
//@Produce json
//@Param id path string true "Quiz ID"
//@Success 200 {object} services.QuizView
//@Failure 400 {string} string "Invalid request"
//@Failure 500 {string} string "Internal server error"
//@Router /get-quiz/{id} [get]
func GetQuizByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	quizID := vars["id"]
	quiz, err := services.GetQuiz(quizID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, quiz)
}

//@Summary Submit answers for a quiz
//@Description Submit answers for a quiz
// @Security ApiKeyAuth
//@Produce json
//@Param id path string true "Quiz ID"
//@Param answers body UserAnswers true "Answers"
//@Success 200 {string} string "Answers submitted successfully"
//@Failure 400 {string} string "Invalid request"
//@Failure 500 {string} string "Internal server error"
//@Router /submit-answers/{id} [post]
func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userId")

	vars := mux.Vars(r)
	quizID := vars["id"]
	
	var userAnswers UserAnswers
	err := json.NewDecoder(r.Body).Decode(&userAnswers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.SubmitAnswers(userID, quizID, userAnswers.Answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, "Answers submitted successfully")
}

func FillDatabase(w http.ResponseWriter, r *http.Request) {
	services.FillRepository()
}

//@Summary Get the result of a quiz
//@Description Get the result of a quiz
//@Security ApiKeyAuth
//@Produce json
//@Param id path string true "Quiz ID"
//@Success 200 {object} UserQuizResult
//@Failure 400 {string} string "Invalid request"
//@Failure 500 {string} string "Internal server error"
//@Router /get-results/{id} [get]
func GetResult(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userId")
	vars := mux.Vars(r)
	quizID := vars["id"]
	result, err := services.GetResult(userID, quizID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := UserQuizResult{
		Score: result.Score,
		Percentage: int(result.Percentage),
		Rank: result.Rank,
		NumberOfUsers: result.TotalUsers,
		CustomizedMessage: "You're better than " +  fmt.Sprintf("%.2f", result.BetterThan*100) + "% of users",
	}
	respondWithJSON(w, http.StatusOK, response)
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

