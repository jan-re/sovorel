package modes

import (
	"fmt"
	"strings"

	"github.com/jan-re/sovorel/utils"
)

type LetterIncludedMode struct {
	GameCore
	IncludeLetter string
	wordsFiltered bool
}

func (m *LetterIncludedMode) PlayRound() bool {
	if !m.wordsFiltered {
		wordsWithLetter := make([]utils.Word, 0, len(m.Words))

		for _, word := range m.Words {
			if strings.Contains(word.Armenian, m.IncludeLetter) {
				wordsWithLetter = append(wordsWithLetter, word)
			}
		}

		if len(wordsWithLetter) == 0 {
			fmt.Println("There are no words in the database with that letter. Sorry!")
			return false
		}

		m.Words = wordsWithLetter
		m.wordsFiltered = true
	}

	word := m.Words[m.index]

	fmt.Println("")
	fmt.Println("Read the following word and press enter: " + word.Armenian)

	m.Reader.ReadString('\n') // for press enter to continue

	return m.finalizeRound(true)
}

func (gc *LetterIncludedMode) PrintScore() {
	fmt.Printf("\nAll done. Total words read: %d\n", gc.score.Correct)
}
