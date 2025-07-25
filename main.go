package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	choiceModeEngToArm = "1"
	choiceModeArmToEng = "2"
	choiceModeCombo    = "3"
)

func main() {
	words := loadAndShuffleWords()
	fmt.Println("Word database successfully loaded.")

	reader := bufio.NewReader(os.Stdin)
	mode := selectGameMode(reader, words)
	fmt.Println("Mode selected. The game begins.")

	continuePlaying := true
	for continuePlaying {
		continuePlaying = mode.PlayRound()
	}

	mode.GetScore().Print()
}

func selectGameMode(reader *bufio.Reader, words []utils.Word) modes.GameMode {
	fmt.Printf("Enter %q for English to Armenian.\n", choiceModeEngToArm)
	fmt.Printf("Enter %q for Armenian to English.\n", choiceModeArmToEng)
	fmt.Printf("Enter %q for a random combination of %q and %q.\n", choiceModeCombo, choiceModeEngToArm, choiceModeArmToEng)
	fmt.Print("Enter choice: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	core := modes.GameCore{Reader: reader, Words: words}

	switch strings.TrimSpace(input) {
	case "1":
		return &modes.EngToArmMode{GameCore: core}
	case "2":
		return &modes.ArmToEngMode{GameCore: core}
	case "3":
		return &modes.ShuffleComboMode{GameCore: core}
	default:
		panic("Unknown input: " + input)
	}
}

func loadAndShuffleWords() []utils.Word {
	bs, err := os.ReadFile(databaseFilename)
	if err != nil {
		panic(err.Error())
	}

	var words []utils.Word
	err = json.Unmarshal(bs, &words)
	if err != nil {
		panic(err.Error())
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return words
}
