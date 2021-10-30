package question

// import (
// 	"fmt"
// )

type Question struct {
	Text string `json:"text"`
	Answer int `json:"answer"`
}

// func (question *Question) String() string {
// 	return fmt.Sprintf("Question:{text:%s, answer:%d}", question.Text, question.Answer)
// }