package main

import (
	"bufio"
	"fmt"
	"os"
)

func runQuiz(q *Quiz) {
	reader := bufio.NewReader(os.Stdin)
	for _, v := range q.questions {
		fmt.Println(v.Question)
		ans, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		s := ans[:len(ans)-2]

		if s == v.Answer.Text {
			fmt.Println("Correct!")
			q.correct++
		} else {
			fmt.Println("Wrong!  The answer was: " + v.Answer.Text)
			q.wrong++
		}
	}

	fmt.Printf("Quiz ended with %d right, %d wrong!\n", q.correct, q.wrong)
}
