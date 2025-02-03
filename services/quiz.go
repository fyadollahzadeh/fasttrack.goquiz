package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"goquiz/repository"
	"io"
	"math/rand"
	"os"
	"time"
)

type QuizView struct {
	ID        string         `json:"id"`
	Title     string         `json:"title"`
	Questions []QuestionView `json:"questions"`
}

type QuestionView struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Choices []string `json:"choices"`
}

type QuizResultView struct {
	Score      int     `json:"score"`
	Percentage float64 `json:"percentage"`
	Rank       int     `json:"rank"`
	TotalUsers int     `json:"totalUsers"`
	BetterThan float64 `json:"betterThan"`
}

var repo *repository.Repository

func init() {
	repo = repository.NewRepository()
}

func FillRepository() {
	bytes, err := readSeedData()
	if err != nil {
		fmt.Println("Error reading seed data:", err)
		return
	}

	var quizzes []repository.Quiz
	err = json.Unmarshal(bytes, &quizzes)
	if err != nil {
		fmt.Println("Error unmarshalling seed data:", err)
		return
	}
	

	for _, quiz := range quizzes {
		repo.AddQuiz(quiz)
	}
}

func readSeedData() ([]byte, error) {
	file, err := os.Open("config/seed.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return bytes, nil
}

func GetQuiz(quizID string) (QuizView, error) {
	quiz, err := repo.GetQuiz(quizID)
	if err != nil {
		return QuizView{}, err
	}

	quizView := QuizView{
		ID:        quiz.ID,
		Title:     quiz.Title,
		Questions: generateShuffledQuestions(quiz.Questions),
	}

	return quizView, nil
}

func generateShuffledQuestions(questions []repository.Question) []QuestionView {
	questionViews := make([]QuestionView, len(questions))
	for i, question := range questions {
		choices := make([]string, len(question.WrongAnswers)+1)
		copy(choices, question.WrongAnswers)
		choices[len(question.WrongAnswers)] = question.RightAnswer

		shuffle(choices)

		questionViews[i] = QuestionView{
			ID:      question.ID,
			Text:    question.Text,
			Choices: choices,
		}
	}
	return questionViews
}

func shuffle(choices []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})
}

var ErrAlreadySubmitted = errors.New("answers already submitted")

func SubmitAnswers(userID string, quizID string, answers []string) error {
	_, err := repo.GetUserQuiz(userID, quizID)
	if err == nil {
		return ErrAlreadySubmitted
	}

	result := repo.SubmitAnswers(userID, quizID, answers)

	quiz, err := repo.GetQuiz(quizID)
	if err != nil {
		return err
	}
	userQuiz, err := repo.GetUserQuiz(userID, quizID)
	if err != nil {
		return err
	}
	score := calculateScore(answers, quiz.Questions)
	repo.UpdateScore(userQuiz, score)

	return result
}

func GetResult(userID string, quizID string) (QuizResultView, error) {
	userQuiz, err := repo.GetUserQuiz(userID, quizID)
	if err != nil {
		return QuizResultView{}, err
	}

	quiz, err := repo.GetQuiz(quizID)
	if err != nil {
		return QuizResultView{}, err
	}

	if len(userQuiz.Answers) != len(quiz.Questions) {
		return QuizResultView{}, repository.ErrInvalidAnswers
	}

	allScores := repo.GetScores(quizID)
	rank := calculateRank(userQuiz.Score, allScores)
	betterThan := calculateBetterThan(userQuiz.Score, allScores)
	percentage := calculatePercentage(userQuiz.Score, len(quiz.Questions))

	result := QuizResultView{
		Score:      userQuiz.Score,
		Percentage: percentage,
		Rank:       rank,
		TotalUsers: len(allScores),
		BetterThan: betterThan,
	}
	return result, nil
}

func calculateScore(answers []string, questions []repository.Question) int {
	score := 0
	for i, answer := range answers {
		if answer == questions[i].RightAnswer {
			score++
		}
	}
	return score
}

func calculateRank(score int, allScores []int) int {
	rank := 1
	for _, s := range allScores {
		if s > score {
			rank++
		}
	}
	return rank
}

func calculatePercentage(score int, totalQuestions int) float64 {
	return float64(score) / float64(totalQuestions) * 100
}

func calculateBetterThan(score int, allScores []int) float64 {
	usersWithLessScore := 0.0
	for _, s := range allScores {
		if s < score {
			usersWithLessScore++
		}
	}
	allUsersCountFloat := float64(len(allScores))

	betterThan := 0.0
	if allUsersCountFloat <= 1 {
		betterThan = 1.0
	} else {
		betterThan = (usersWithLessScore / allUsersCountFloat)
	}
	return betterThan
}
