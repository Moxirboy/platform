package postgres

// import (
// 	"database/sql"
// 	"context"
// 	"student/internal/models"
// 	"student/internal/service/repo"
// 	logger "student/pkg/log"

// )

// type studentRepo struct{
// 	db  *sql.DB
// 	log logger.Logger
// }

// func NewStudentRepo(db *sql.DB, log logger.Logger ) *studentRepo{
// 	return &studentRepo{
// 		db:db,
// 		log:log,
// 	}
// }

// func (s *studentRepo) TakeExam(ctx context.Context, answers models.CreateAnswers) (models.TestResult, error) {
// 	for _, answer := range answers.Answers {
// 		_, err := s.db.ExecContext(ctx, "insert into answers (question_id, choice_id, is_correct) values ($1, $2, $3)",
// 			answer.QuestionID, answer.ChoiceID, answer.IsCorrect)
// 		if err != nil {
// 			return models.TestResult{}, err
// 		}
// 	}
// 	result, err := s.evaluateTestResult(answers.ExamID)
// 	if err != nil {
// 		return models.TestResult{}, err
// 	}

// 	return models.TestResult{Result: result}, nil
// }

// func (s *studentRepo) GetTestResults(ctx context.Context, examID string) ([]models.TestResult, error) {

// 	rows, err := s.db.QueryContext(ctx, "select question_id, is_correct from answers where question_id IN (select id from questions where exam_id = $1)", examID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	resultsMap := make(map[string]string)

// 	for rows.Next() {
// 		var questionID string
// 		var isCorrect bool
// 		if err := rows.Scan(&questionID, &isCorrect); err != nil {
// 			return nil, err
// 		}
// 		result := "Fail"
// 		if isCorrect {
// 			result = "Pass"
// 		}
// 		resultsMap[questionID] = result
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	var results []models.TestResult
// 	for questionID, result := range resultsMap {
// 		results = append(results, models.TestResult{
// 			QuestionID: questionID,
// 			Result:     result,
// 		})
// 	}

// 	return results, nil
// }

// func (s *studentRepo) evaluateTestResult(examID string) (string, error) {
// 	score := 0
// 	if score >= passingScore {
// 		return "Pass", nil
// 	} else {
// 		return "Fail", nil
// 	}
// }
