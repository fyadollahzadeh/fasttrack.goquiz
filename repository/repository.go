package repository

import (
    "errors"
)

var ErrInvalidAnswers = errors.New("invalid number of answers")
var ErrQuizNotFound = errors.New("quiz not found")
type Question struct {
	ID       string      `json:"id"`
	Text     string   `json:"text"`
	WrongAnswers  []string `json:"WrongAnswers"`
	RightAnswer   string   `json:"RightAnswer"`
}

type Quiz struct {
	ID  string `json:"id"`
	Title string `json:"title"`
	Questions []Question `json:"questions"`
}
	
type UserQuiz struct {
	UserID string `json:"userID"`
	QuizID string `json:"quizID"`
	Answers []string `json:"answers"`
	Score int `json:"score"`
}

type Repository struct {
    Quizzes     map[string]Quiz
    UserQuizzes []UserQuiz
}

func NewRepository() *Repository {
    return &Repository{
        Quizzes: make(map[string]Quiz),
    }   
}

func (r *Repository) GetAllQuizzes() []Quiz{
	var allQuizzes []Quiz
	for _, quiz := range r.Quizzes {
		allQuizzes = append(allQuizzes, quiz)
	}

    return allQuizzes
}

func (r *Repository) AddQuiz(quiz Quiz) error {
    r.Quizzes[quiz.ID] = quiz
    return nil
}

func (r *Repository) GetQuiz(quizID string) (Quiz, error) {
    quiz, ok := r.Quizzes[quizID]
    if !ok {
        return Quiz{}, ErrQuizNotFound
    }
    return quiz, nil
}

func (r *Repository) UpdateQuiz(quiz Quiz) error {
    r.Quizzes[quiz.ID] = quiz
    return nil
}

func (r *Repository) AddQuestion(quizID string, question Question) error {
    quiz, err := r.GetQuiz(quizID)
    if err != nil {
        return err
    }
    quiz.Questions = append(quiz.Questions, question)
    return r.UpdateQuiz(quiz)
}

func (r *Repository) GetQuestions(quizID string) ([]Question, error) {
    quiz, err := r.GetQuiz(quizID)
    if err != nil {
        return nil, err
    }
    return quiz.Questions, nil
}
func (r *Repository) SubmitAnswers(userID string, quizID string, answers []string) error {
    quiz, err := r.GetQuiz(quizID)
    if err != nil {
        return err
    }
    if len(answers) != len(quiz.Questions) {
        return ErrInvalidAnswers
    }
    userQuiz := UserQuiz{
        UserID: userID,
        QuizID: quizID,
        Answers: answers,
    }
    r.UserQuizzes = append(r.UserQuizzes, userQuiz)
    return nil
}

func (r *Repository) GetUserQuiz(userID string, quizID string) (UserQuiz, error) {
    for _, uq := range r.UserQuizzes {
        if uq.UserID == userID && uq.QuizID == quizID {
            return uq, nil
        }
    }
    return UserQuiz{}, ErrQuizNotFound
}

func (r *Repository) GetScores(quizID string) []int {
    scores := []int{}
    for _, uq := range r.UserQuizzes {
        if uq.QuizID == quizID {
            scores = append(scores, uq.Score)
        }
    }
    return scores
}

func (r *Repository) UpdateScore(userQuiz UserQuiz, score int) {
    for i, uq := range r.UserQuizzes {
        if uq.UserID == userQuiz.UserID && uq.QuizID == userQuiz.QuizID {
            r.UserQuizzes[i].Score = score
        }
    }
}



