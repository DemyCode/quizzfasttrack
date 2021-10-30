package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	// "strings"

	"github.com/DemyCode/quizzfasttrack"
	"github.com/DemyCode/quizzfasttrack/lib/question"
	"github.com/DemyCode/quizzfasttrack/lib/errorhandler"
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