package utils

import "fmt"

type Score struct {
	correct   int
	incorrect int
}

func (s *Score) Increment(wasCorrect bool) {
	if wasCorrect {
		s.correct++
	} else {
		s.incorrect++
	}

}

func (s Score) Print() {
	fmt.Println("")
	fmt.Println("That's all the words. Here are your statistics:")
	fmt.Printf("Total words: %d\nCorrect: %d\nIncorrect: %d\n", s.correct+s.incorrect, s.correct, s.incorrect)
}
