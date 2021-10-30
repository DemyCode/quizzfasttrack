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
	"log"
	// "net"
	"net/http"
	"io/ioutil"

	"github.com/DemyCode/quizzfasttrack"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "questions",
	Short: "Displays the Questions",
	Long: `Usage : questions
Description : Displays the Questions containted in the description of questions.go file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var adress = "http://" + config.HOST + ":" + config.PORT
		var resp, errgetquestions = http.Get(adress + "/questions")
		if (errgetquestions != nil) {
			log.Fatalln(errgetquestions)
		}
		var body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))
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
