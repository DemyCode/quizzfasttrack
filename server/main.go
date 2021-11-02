package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	// "strings"

	"github.com/DemyCode/quizzfasttrack"
	"github.com/DemyCode/quizzfasttrack/lib/errorhandler"
	"github.com/DemyCode/quizzfasttrack/lib/question"
)

func main() {
	// set Flags for logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var err error
	var fatal bool = true

	// Checking if the questions.json file is not corrupted
	jsonFile, err := ioutil.ReadFile(config.QUESTIONPATH)
	errorhandler.ErrorHandler(err, fatal)
	

	// Checking if the questions.json file is not corrupted
	var questions []question.Question
	err = json.Unmarshal(jsonFile, &questions)
	errorhandler.ErrorHandler(err, fatal)

	// for _, question := range questions {
	// 	fmt.Println(question)
	// }

	// Turning on the server
	var adress = config.HOST + ":" + config.PORT
	http.HandleFunc("/questions", handleQuestions)
	http.HandleFunc("/answer", handleAnswers)
	fmt.Println("Server running on ", adress)
	err = http.ListenAndServe(adress, nil)
	errorhandler.ErrorHandler(err, fatal)
}

func handleQuestions(w http.ResponseWriter, r *http.Request){
	var err error
	var fatal bool = false
	w.Header().Set("Content-Type", "application/json")

	// Reading JSON everytime not the best but lets roll with that for now
	jsonFile, err := ioutil.ReadFile(config.QUESTIONPATH)
	errorhandler.ErrorHandler(err, fatal)
	var questions []question.Question
	err = json.Unmarshal(jsonFile, &questions)
	if err != nil {
		errorhandler.ErrorHandler(err, fatal)
		return
	}

	var textonly []string
	for _, question := range questions{
		textonly = append(textonly, question.Text)
	}
	
	marshquestionned, err := json.Marshal(textonly)
	if err != nil {	
		errorhandler.ErrorHandler(err, fatal)
		return
	}
	
	w.Write(marshquestionned)
}

var memory []float64

func handleAnswers(w http.ResponseWriter, r *http.Request){
	var err error
	var fatal bool = false
	w.Header().Set("Content-Type", "application/json")

	// fmt.Println("What")

	body, err := ioutil.ReadAll(r.Body)
	errorhandler.ErrorHandler(err, fatal)
	// fmt.Println(body)

	// Getting the answers
	var stringanswers []string
	err = json.Unmarshal(body, &stringanswers)
	errorhandler.ErrorHandler(err, fatal)

	// fmt.Println(stringanswers)
	
	// Reading the Json questions
	jsonFile, err := ioutil.ReadFile(config.QUESTIONPATH)
	errorhandler.ErrorHandler(err, fatal)
	var questions []question.Question
	err = json.Unmarshal(jsonFile, &questions)
	if err != nil {
		errorhandler.ErrorHandler(err, fatal)
		return
	}

	// Number of answers different from the number of Questions
	if len(stringanswers) != len(questions) {
		w.Write([]byte("Number of answers different from the number of Questions"))
		return
	}

	var res int
	for i := 0; i < len(stringanswers); i++ {
		answer, err := strconv.Atoi(stringanswers[i])
		if err != nil {
			w.Write([]byte("Please answers must be of type int [0-9]+"))
			return
		}
		if answer == questions[i].GetAnswer() {
			res += 1
		}
	}

	var percentscore float64 = float64(res) / float64(len(stringanswers))
	var beaten int = 0
	for i := 0; i < len(memory); i++ {
		if memory[i] < percentscore{
			beaten += 1
		} 
	}
	var result string = fmt.Sprintf("Thank you for taking this Quizz ! You have %d correct answers out of %d questions\n", res, len(stringanswers))
	result += fmt.Sprintf("You were better than %.0f percent of all quizzers", 100 * (float64(beaten) / float64(len(memory))))

	memory = append(memory, percentscore)

	w.Write([]byte(result))
}