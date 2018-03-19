package main

/*Stories
- Load questions from file
- Ask all questions in random order
- Answers given on cli
*/

/*Algorithms
- Read file
- Load into json struct
- Produce a random ordering of questions to ask
- Ask a question
- Wait for answer
- Grade answer
- Produce report at end with total correct
*/

const quizFile = "questions.json"

type report struct {
	correct int
	total   int
}

func main() {
	qb := load(quizFile)

	qhints := QuizPolicy{
		"selection": "all",
		"ordering":  "shuffle",
	}

	q := newQuiz(qb, qhints)

	runQuiz(q)
}
