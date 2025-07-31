package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/jan-re/sovorel/modes"
	"github.com/jan-re/sovorel/utils"
)

const (
	databaseFilename = "db.json"
)

const (
	choiceModeEngToArm       = "1"
	choiceModeArmToEng       = "2"
	choiceModeCombo          = "3"
	choiceModeLetterIncluded = "4"
)

func main() {
	words, err := loadAndShuffleWords()
	if err != nil {
		log.Fatal("Failed to initialize word database. Error: " + err.Error())
	}

	fmt.Println("Word database successfully loaded.")

	reader := bufio.NewReader(os.Stdin)
	mode, err := selectGameMode(reader, words)
	if err != nil {
		log.Fatal("Failed to initialize gamemode. Error: " + err.Error())
	}

	fmt.Println("Mode selected. The game begins.")

	continuePlaying := true
	for continuePlaying {
		continuePlaying = mode.PlayRound()
	}

	mode.PrintScore()
}

func selectGameMode(reader *bufio.Reader, words []utils.Word) (modes.GameMode, error) {
	presentChoices()

	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	core := modes.GameCore{Reader: reader, Words: words}

	switch strings.TrimSpace(input) {
	case "1":
		return &modes.EngToArmMode{GameCore: core}, nil
	case "2":
		return &modes.ArmToEngMode{GameCore: core}, nil
	case "3":
		return &modes.ShuffleComboMode{GameCore: core}, nil
	case "4":
		fmt.Print("Enter letter: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read input: %w", err)
		}

		trimmedInput := strings.TrimSpace(input)

		if trimmedInput == "" {
			return nil, errors.New("The letter cannot be empty. Sorry.")
		}

		return &modes.LetterIncludedMode{GameCore: core, IncludeLetter: trimmedInput}, nil
	default:
		return nil, errors.New("You've entered an unsupported choice. Sorry.")
	}
}

func presentChoices() {
	fmt.Printf("Enter %q for English to Armenian.\n", choiceModeEngToArm)
	fmt.Printf("Enter %q for Armenian to English.\n", choiceModeArmToEng)
	fmt.Printf("Enter %q for a random combination of %q and %q.\n", choiceModeCombo, choiceModeEngToArm, choiceModeArmToEng)
	fmt.Printf("Enter %q for a reading exercise of Armenian words with a specific letter included.\n", choiceModeLetterIncluded)
	fmt.Print("Enter choice: ")
}

func loadAndShuffleWords() ([]utils.Word, error) {
	bs, err := os.ReadFile(databaseFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file with words: %w", err)
	}

	var words []utils.Word
	err = json.Unmarshal(bs, &words)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file with words: %w", err)
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return words, nil
}
