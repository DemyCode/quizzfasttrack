/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/DemyCode/quizzfasttrack"
	"github.com/DemyCode/quizzfasttrack/lib/errorhandler"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "questions",
	Short: "Displays the Questions",
	Long: `Usage : questions
Description : Displays the Questions containted in the description of questions.go file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var fatal bool = true
		var adress = "http://" + config.HOST + ":" + config.PORT
		resp, err := http.Get(adress + "/questions")
		errorhandler.ErrorHandler(err, fatal)

		body, err := ioutil.ReadAll(resp.Body)
		errorhandler.ErrorHandler(err, fatal)
		
		var textonly []string
		err = json.Unmarshal(body, &textonly)
		errorhandler.ErrorHandler(err, fatal)

		fmt.Println("Here are the questions !\n")
		for _, text := range textonly {
			fmt.Println(text + "\n")
		}

		defer resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
