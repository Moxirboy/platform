package handler
import (pb "student/proto/exam"
	"student/internal/models"
)


func ConvertToTestResponse(questions []models.QuestionsAll) *pb.TestResponse {
    var tests []*pb.Test
	var choices []*pb.Choices
    for _, q := range questions {
        
        for _, choice := range q.Choices {
            c := &pb.Choices{
                Id:         choice.Id,
                Choices: choice.ChoiceText,
                IsCorrect:  choice.Is_correct,
            }
            choices = append(choices, c)
        }

        test := &pb.Test{
            Id:       q.Id,
            Question: q.QuestionText,
            Choices:  choices,
            
        }
        tests = append(tests, test)
    }

    return &pb.TestResponse{
        Test: tests,
    }
}

// func choicesToStrings(choices []*pb.Choices) []string {
//     var choiceStrings []string
//     for _, choice := range choices {
		
//         choiceStrings = append(choiceStrings, choice.Choices)
//     }
//     return choiceStrings
// }
