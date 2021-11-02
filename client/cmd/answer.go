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
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"

	"github.com/spf13/cobra"

	"github.com/DemyCode/quizzfasttrack"
	"github.com/DemyCode/quizzfasttrack/lib/errorhandler"
)

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Used to answer the questions received (please use questions command before)",
	Long: `Usage : ./client answer ([0-9] )*
result : Is the integer corresponding to the answer of your question.
Answer questions in respective order

Example : 
$ ./client questions
A) What is the capital of France ?
1) Paris 2) Amsterdam 3) London 4) New York
B) What is the capital of Italy ?
1) Firenze 2) Vatican 3) Roma 4) Venise
$ ./client answer 1 3`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var fatal bool = true

		var adress = "http://" + config.HOST + ":" + config.PORT
		// resp, err := http.Get(adress + "/answer")

		jsondata, err := json.Marshal(args)
		errorhandler.ErrorHandler(err, fatal)
		resp, err := http.Post(adress + "/answer", "application/json", bytes.NewBuffer(jsondata))
		errorhandler.ErrorHandler(err, fatal)
		body, err := ioutil.ReadAll(resp.Body)
		errorhandler.ErrorHandler(err, fatal)

		fmt.Println(string(body))

		defer resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(answerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// answerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// answerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
