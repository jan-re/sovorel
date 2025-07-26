package modes

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	logErrorReadingInput = "An error occurred while reading your input."
)

func playTranslationGame(reader *bufio.Reader, query, answer string) (wasCorrect bool, err error) {
	fmt.Println("")
	fmt.Println("Translate: " + query)

	input, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read input: %w", err)
	}

	if strings.TrimSpace(input) == answer {
		fmt.Println("Correct!")
		return true, nil
	}

	fmt.Println("Wrong. Correct answer is: " + answer)
	return false, nil
}
