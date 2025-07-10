package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	modeEngToArm = "eng2arm"
	modeArmToEng = "arm2eng"
)

func main() {
	bs, err := os.ReadFile("db.json")
	if err != nil {
		panic(err.Error())
	}

	var words []word
	err = json.Unmarshal(bs, &words)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Word database successfully loaded.")
	fmt.Println("Enter 1 for English to Armenian.")
	fmt.Println("Enter 2 for Armenian to English.")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter choice: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	var mode string
	switch strings.TrimSpace(input) {
	case "1":
		mode = modeEngToArm
	case "2":
		mode = modeArmToEng
	default:
		panic("Unknown input: " + input)
	}

	fmt.Println("Mode selected.")

	// TODO Shuffle words

	for _, word := range words {
		var query, answer string

		if mode == modeEngToArm {
			query = word.English
			answer = word.Armenian
		} else {
			query = word.Armenian
			answer = word.English
		}
		fmt.Println("")
		fmt.Println(query)

		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}

		if strings.TrimSpace(input) == answer {
			fmt.Println("Correct")
			continue
		}

		fmt.Println("Wrong. Correct answer is: " + answer)
	}
}
