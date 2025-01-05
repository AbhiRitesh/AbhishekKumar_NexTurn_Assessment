package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Question struct to store each question and options and correct answer
type Question struct {
	Question string
	Options  []string
	Answer   int 
}

func main() {
	// Question Bank
	questions := []Question{
		{"Which is the tallest building of Dubai?", []string{"Burj Al Arab", "Dubai Frame", "Burj Khalifa", "Dubai Opera"}, 3},
		{"Which programming language is most widly used?", []string{"Python", "C++", "Go", "Java"}, 2},
		{"What is 5 * 3?", []string{"5", "18", "9", "15"}, 4},
	}

	fmt.Println("Welcome to the Online Examination System!")
	score := 0

	reader := bufio.NewReader(os.Stdin)
	timerDuration := 15 * time.Second

	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}
		fmt.Println("Enter your choice (or type 'exit' to quit):")

		answerChan := make(chan int)
		errChan := make(chan error)

		// Timer for each question
		go func() {
			for {
				fmt.Print("> ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if strings.ToLower(input) == "exit" {
					errChan <- errors.New("exit")
					return
				}

				choice, err := strconv.Atoi(input)
				if err != nil || choice < 1 || choice > len(q.Options) {
					fmt.Println("Invalid input. Please enter a valid option number.")
					continue
				}

				answerChan <- choice
				return
			}
		}()

		select {
		case answer := <-answerChan:
			if answer == q.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong!")
			}
		case err := <-errChan:
			if err.Error() == "exit" {
				fmt.Println("Exiting the quiz...")
				showResults(score, len(questions))
				return
			}
		case <-time.After(timerDuration):
			fmt.Println("\nTime's up for this question!")
		}
	}

	// Display results
	showResults(score, len(questions))
}

func showResults(score, total int) {
	fmt.Printf("\nQuiz Over! Your score: %d/%d\n", score, total)
	percentage := (float64(score) / float64(total)) * 100

	if percentage >= 80 {
		fmt.Println("Performance: Excellent")
	} else if percentage >= 50 {
		fmt.Println("Performance: Good")
	} else {
		fmt.Println("Performance: Needs Improvement")
	}
}
