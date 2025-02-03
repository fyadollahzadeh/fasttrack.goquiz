package services

import (
	"goquiz/repository"
	"testing"
)
var testQuiz = repository.Quiz{
	ID: "1",
	Title: "TestQuiz",
	Questions: []repository.Question{
		{
			ID:           "1",
			Text:         "4",
			WrongAnswers: []string{"A", "B"},
			RightAnswer:  "C",
		},
		{
			ID:           "2",
			Text:         "4",
			WrongAnswers: []string{"A", "C"},
			RightAnswer:  "B",
		},
		{
			ID:           "3",
			Text:         "4",
			WrongAnswers: []string{"C", "B"},
			RightAnswer:  "A",
		},
	},
}

func TestGetQuiz(t *testing.T) {
	repo.AddQuiz(testQuiz)
	quizID := "1"
	quiz, err := GetQuiz(quizID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if quiz.ID != quizID {
		t.Errorf("Expected quiz ID %v, got %v", quizID, quiz.ID)
	}
}

func TestSubmitAnswers(t *testing.T) {
	
	repo.AddQuiz(testQuiz)

	userID := "user1"
	quizID := "1"
	answers := []string{"A", "B", "C"}
	err := SubmitAnswers(userID, quizID, answers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = SubmitAnswers(userID, quizID, answers)
	if err != ErrAlreadySubmitted {
		t.Errorf("Expected error %v, got %v", ErrAlreadySubmitted, err)
	}
}

func TestGetResult(t *testing.T) {

	userID := "user1"
	quizID := "1"
	answers := []string{"A", "B", "C","D","E"}
	SubmitAnswers(userID, quizID, answers)
	result, err := GetResult(userID, quizID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Score == 0 {
		t.Error("Expected non-zero score")
	}
}

func TestCalculateScore(t *testing.T) {
	questions := []repository.Question{
		{RightAnswer: "A"},
		{RightAnswer: "B"},
		{RightAnswer: "C"},
	}
	answers := []string{"A", "B", "C"}
	score := calculateScore(answers, questions)
	if score != 3 {
		t.Errorf("Expected score 3, got %v", score)
	}
}

func TestCalculateRank(t *testing.T) {
	allScores := []int{1, 2, 3, 4, 5}
	rank := calculateRank(3, allScores)
	if rank != 3 {
		t.Errorf("Expected rank 3, got %v", rank)
	}
}

func TestCalculatePercentage(t *testing.T) {
	percentage := calculatePercentage(3, 5)
	if percentage != 60 {
		t.Errorf("Expected percentage 60, got %v", percentage)
	}
}

func TestCalculateBetterThan(t *testing.T) {
	allScores := []int{1, 2, 3, 4, 5}
	betterThan := calculateBetterThan(3, allScores)
	if betterThan != 0.4 {
		t.Errorf("Expected betterThan 0.4, got %v", betterThan)
	}
}