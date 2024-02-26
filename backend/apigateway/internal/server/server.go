package server

import (
	configs "gateway/internal/config"

	ht "gateway/internal/controller/http"
	client "gateway/pkg/grpc"
	logger "gateway/pkg/log"
	"log"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	cfg *configs.Config
	log logger.Logger
}

func NewServer(
	cfg *configs.Config,
	log logger.Logger,
) *Server {
	return &Server{
		cfg: cfg,
		log: log,
	}
}

func (s *Server) Run(port string) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Current Working Directory:", wd)
	r := gin.New()
	r.StaticFile("/","internal/controller/http/v1/templates/login.html")
	r.StaticFile("/signup","internal/controller/http/v1/templates/sign.html")
	r.StaticFile("/dashboard","internal/controller/http/v1/templates/dashboard.html")
	r.StaticFile("/test","internal/controller/http/v1/templates/test.html")
	r.GET("/questions", getQuestions)
	r.POST("/check-answers", checkAnswers)
	r.GET("/test-available", Available)

	r.LoadHTMLGlob("internal/controller/http/v1/templates/*")
	log.Println(s.cfg.Admin_RPC.Hosts, s.cfg.Admin_RPC.Port)
	grpc_client, err := client.New(*s.cfg)
	if err != nil {
		log.Println("here")
		return err
	}

	ht.SetUp(r, *s.cfg, grpc_client, s.log)

	return r.Run(port)
}

type Question struct {
	ID                int      `json:"id"`
	Text              string   `json:"text"`
	Choices           []string `json:"choices"`
	CorrectChoiceIndex int      `json:"correctChoiceIndex"`
}

var questions = []Question{
	{1, "Question 1: Allergiya nima?", []string{"infeksion allergik kasallik", "ortirilgan immunitet tanqisligi", "surunkali kasallik asorati"}, 0},
		{2, "Question 2: What is the capital of Germany?", []string{"Berlin", "Paris", "London"}, 0},
		{3, "Question 3: What is the largest ocean?", []string{"Pacific Ocean", "Atlantic Ocean", "Indian Ocean"}, 0},
		{4, "Question 4: What is the chemical symbol for water?", []string{"H2O", "CO2", "O2"}, 0},
		{5, "Question 5: Who wrote 'To Kill a Mockingbird'?", []string{"Harper Lee", "Stephen King", "J.K. Rowling"}, 0},
		{6, "Question 6: What is the capital of France?", []string{"Berlin", "Paris", "London"}, 1},
		{7, "Question 7: What is the capital of Japan?", []string{"Tokyo", "Beijing", "Seoul"}, 0},
		{8, "Question 8: Who painted the Mona Lisa?", []string{"Leonardo da Vinci", "Pablo Picasso", "Vincent van Gogh"}, 0},
		{9, "Question 9: What is the currency of Brazil?", []string{"Peso", "Real", "Ruble"}, 1},
		{10, "Question 10: Which planet is known as the Red Planet?", []string{"Mars", "Venus", "Jupiter"}, 0},
		{11, "Question 11: What is the tallest mountain in the world?", []string{"Mount Everest", "K2", "Kangchenjunga"}, 0},
		{12, "Question 12: Who is known as the 'Father of Computers'?", []string{"Charles Babbage", "Alan Turing", "Ada Lovelace"}, 0},
		{13, "Question 13: What is the chemical symbol for gold?", []string{"Au", "Ag", "Cu"}, 0},
		{14, "Question 14: What is the main ingredient in guacamole?", []string{"Avocado", "Tomato", "Onion"}, 0},
		{15, "Question 15: Who wrote '1984'?", []string{"George Orwell", "Aldous Huxley", "Ray Bradbury"}, 0},
		{16, "Question 16: What is the capital of Australia?", []string{"Canberra", "Sydney", "Melbourne"}, 0},
		{17, "Question 17: Which gas do plants use for photosynthesis?", []string{"Carbon dioxide", "Oxygen", "Nitrogen"}, 0},
		{18, "Question 18: Who discovered penicillin?", []string{"Alexander Fleming", "Louis Pasteur", "Marie Curie"}, 0},
		{19, "Question 19: Which country is commonly associated with the Great Wall?", []string{"China", "India", "Japan"}, 0},
		{20, "Question 20: Who painted the ceiling of the Sistine Chapel?", []string{"Michelangelo", "Leonardo da Vinci", "Raphael"}, 0},
		// Add more questions here
}

func getQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"questions": questions})
}
func Available(c *gin.Context){
	if len(questions)==0{
		c.JSON(http.StatusOK, gin.H{"testAvailable": false})
	}

        // Return JSON response indicating if the test is available
        c.JSON(http.StatusOK, gin.H{"testAvailable": true})
}
func checkAnswers(c *gin.Context) {
	var userAnswers []struct {
		QuestionID  int `json:"questionId"`
		AnswerIndex int `json:"answerIndex"`
	}
	if err := c.BindJSON(&userAnswers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var score int
	for _, ans := range userAnswers {
		if ans.AnswerIndex == questions[ans.QuestionID-1].CorrectChoiceIndex {
			score++
		}
	}

	c.JSON(http.StatusOK, gin.H{"score": score, "totalQuestions": len(questions)})
}
