package repository

import (
	"testing"
)

func TestAddQuiz(t *testing.T) {
	repo := NewRepository()
	quiz := Quiz{ID: "1", Title: "Sample Quiz"}
	err := repo.AddQuiz(quiz)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(repo.Quizzes) != 1 {
		t.Errorf("expected 1 quiz, got %d", len(repo.Quizzes))
	}
}

func TestGetQuiz(t *testing.T) {
	repo := NewRepository()
	quiz := Quiz{ID: "1", Title: "Sample Quiz"}
	repo.AddQuiz(quiz)
	retrievedQuiz, err := repo.GetQuiz("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if retrievedQuiz.ID != "1" {
		t.Errorf("expected quiz ID '1', got %s", retrievedQuiz.ID)
	}
}

func TestGetQuizNotFound(t *testing.T) {
	repo := NewRepository()
	_, err := repo.GetQuiz("1")
	if err != ErrQuizNotFound {
		t.Errorf("expected error %v, got %v", ErrQuizNotFound, err)
	}
}

func TestUpdateQuiz(t *testing.T) {
	repo := NewRepository()
	quiz := Quiz{ID: "1", Title: "Sample Quiz"}
	repo.AddQuiz(quiz)
	updatedQuiz := Quiz{ID: "1", Title: "Updated Quiz"}
	err := repo.UpdateQuiz(updatedQuiz)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	retrievedQuiz, _ := repo.GetQuiz("1")
	if retrievedQuiz.Title != "Updated Quiz" {
		t.Errorf("expected quiz title 'Updated Quiz', got %s", retrievedQuiz.Title)
	}
}

func TestAddQuestion(t *testing.T) {
	repo := NewRepository()
	quiz := Quiz{ID: "1", Title: "Sample Quiz"}
	repo.AddQuiz(quiz)
	question := Question{ID: "1", Text: "Sample Question", RightAnswer: "A"}
	err := repo.AddQuestion("1", question)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	retrievedQuiz, _ := repo.GetQuiz("1")
	if len(retrievedQuiz.Questions) != 1 {
		t.Errorf("expected 1 question, got %d", len(retrievedQuiz.Questions))
	}
}

func TestGetQuestions(t *testing.T) {
	repo := NewRepository()
	quiz := Quiz{ID: "1", Title: "Sample Quiz"}
	repo.AddQuiz(quiz)
	question := Question{ID: "1", Text: "Sample Question", RightAnswer: "A"}
	repo.AddQuestion("1", question)
	questions, err := repo.GetQuestions("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(questions) != 1 {
		t.Errorf("expected 1 question, got %d", len(questions))
	}
}

func TestSubmitAnswers(t *testing.T) {
	repo := NewRepository()
	quiz := Quiz{ID: "1", Title: "Sample Quiz"}
	question := Question{ID: "1", Text: "Sample Question", RightAnswer: "A"}
	quiz.Questions = append(quiz.Questions, question)
	repo.AddQuiz(quiz)
	answers := []string{"A"}
	err := repo.SubmitAnswers("user1", "1", answers)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(repo.UserQuizzes) != 1 {
		t.Errorf("expected 1 user quiz, got %d", len(repo.UserQuizzes))
	}
}

func TestSubmitAnswersInvalid(t *testing.T) {
	repo := NewRepository()
	quiz := Quiz{ID: "1", Title: "Sample Quiz"}
	question := Question{ID: "1", Text: "Sample Question", RightAnswer: "A"}
	quiz.Questions = append(quiz.Questions, question)
	repo.AddQuiz(quiz)
	answers := []string{"A", "B"}
	err := repo.SubmitAnswers("user1", "1", answers)
	if err != ErrInvalidAnswers {
		t.Errorf("expected error %v, got %v", ErrInvalidAnswers, err)
	}
}