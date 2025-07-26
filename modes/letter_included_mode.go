package modes

import (
	"fmt"
	"strings"

	"github.com/jan-re/sovorel/utils"
)

type LetterIncludedMode struct {
	GameCore
	IncludeLetter   string
	wordsWithLetter []utils.Word
}

func (m *LetterIncludedMode) PlayRound() bool {
	if m.wordsWithLetter == nil {
		for _, word := range m.Words {
			if strings.Contains(word.Armenian, m.IncludeLetter) {
				m.wordsWithLetter = append(m.wordsWithLetter, word)
			}
		}

		if len(m.wordsWithLetter) == 0 {
			fmt.Println("There are no words in the database with that letter. Sorry!")
			return false
		}
	}

	word := m.wordsWithLetter[m.index]

	wasCorrect := playTranslationGame(m.Reader, word.Armenian, word.English)

	m.score.Increment(wasCorrect)

	// Stop the game if we've reached the last word.
	if m.index == len(m.wordsWithLetter)-1 {
		return false
	}

	m.index++
	return true
}

func (m *LetterIncludedMode) GetScore() utils.Score {
	return m.score
}
