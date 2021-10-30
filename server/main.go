package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"

	"github.com/DemyCode/QuizzBackEnd/config"
	"github.com/DemyCode/QuizzBackEnd/questionpack"
)

func main() {
	// Checking if the questions.json file is not corrupted
	jsonFile, errbyte := ioutil.ReadFile(config.QUESTIONPATH)
	// jsonFile = []byte(strings.ReplaceAll(string(jsonFile), "\n", ""))
	if errbyte != nil{
		fmt.Println("errbyte", errbyte)
		return
	}

	// Checking if the questions.json file is not corrupted
	var questions []questionpack.Question
	var errjson = json.Unmarshal(jsonFile, &questions)
	if errjson != nil{
		fmt.Println("errjson", errjson)
		return
	}

	for _, question := range questions {
		fmt.Println(question)
	}

	var adress = config.HOST + ":" + config.PORT
	http.HandleFunc("/questions", handleQuestions)
	fmt.Println("Server running on ", adress)
	var err = http.ListenAndServe(adress, nil)
	if err != nil{
		fmt.Println(err)
	}
}

func handleQuestions(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	jsonFile, err := ioutil.ReadFile(config.QUESTIONPATH)
	if err != nil{
		fmt.Println(err)
		return
	}
	w.Write(jsonFile)
}