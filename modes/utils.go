package modes

import (
	"bufio"
	"fmt"
	"strings"
)

func playTranslationGame(reader *bufio.Reader, query, answer string) (wasCorrect bool) {
	fmt.Println("")
	fmt.Println("Translate: " + query)

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	if strings.TrimSpace(input) == answer {
		fmt.Println("Correct!")
		return true
	}

	fmt.Println("Wrong. Correct answer is: " + answer)
	return false
}
