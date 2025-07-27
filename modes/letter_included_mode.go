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
	}

	word := m.Words[m.index]

	wasCorrect, err := playTranslationGame(m.Reader, word.Armenian, word.English)
	if err != nil {
		fmt.Println(logErrorReadingInput)
		return false
	}

	return m.finalizeRound(wasCorrect)
}
