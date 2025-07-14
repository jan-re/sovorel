package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

const (
	databaseFilename = "db.json"
)

const (
	modeEngToArm = "1"
	modeArmToEng = "2"
	modeCombo    = "3"
)

func main() {
	words := loadAndShuffleWords()
	fmt.Println("Word database successfully loaded.")

	reader := bufio.NewReader(os.Stdin)
	gameModeFunc := selectGameMode(reader)
	fmt.Println("Mode selected. The game begins.")

	var correct, incorrect int
	for _, word := range words {
		wasCorrect := playGame(word, reader, gameModeFunc)

		if wasCorrect {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Println("")
	fmt.Println("That's all the words. Here are your statistics:")
	fmt.Printf("Total words: %d\nCorrect: %d\nIncorrect: %d\n", correct+incorrect, correct, incorrect)
}

func engToArmFunc(w word) (string, string) {
	return w.English, w.Armenian
}

func armToEngFunc(w word) (string, string) {
	return w.Armenian, w.English
}

func randomModeFunc(w word) (string, string) {
	n := rand.IntN(2)

	if n == 0 {
		return w.Armenian, w.English
	}

	return w.English, w.Armenian
}

func loadAndShuffleWords() []word {
	bs, err := os.ReadFile(databaseFilename)
	if err != nil {
		panic(err.Error())
	}

	var words []word
	err = json.Unmarshal(bs, &words)
	if err != nil {
		panic(err.Error())
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return words
}

func selectGameMode(reader *bufio.Reader) func(w word) (string, string) {
	fmt.Printf("Enter %q for English to Armenian.\n", modeEngToArm)
	fmt.Printf("Enter %q for Armenian to English.\n", modeArmToEng)
	fmt.Printf("Enter %q for a random combination of %q and %q.\n", modeCombo, modeEngToArm, modeArmToEng)
	fmt.Print("Enter choice: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	switch strings.TrimSpace(input) {
	case "1":
		return engToArmFunc
	case "2":
		return armToEngFunc
	case "3":
		return randomModeFunc
	default:
		panic("Unknown input: " + input)
	}
}

func playGame(w word, reader *bufio.Reader, gameModeFunc func(w word) (string, string)) (wasCorrect bool) {
	query, answer := gameModeFunc(w)

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
