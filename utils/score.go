package utils

import "fmt"

type Score struct {
	Correct   int
	Incorrect int
}

func (s *Score) Increment(wasCorrect bool) {
	if wasCorrect {
		s.Correct++
	} else {
		s.Incorrect++
	}

}

func (s Score) Print() {
	fmt.Println("")
	fmt.Println("Here are your statistics:")
	fmt.Printf("Total words: %d\nCorrect: %d\nIncorrect: %d\n", s.Correct+s.Incorrect, s.Correct, s.Incorrect)
}
