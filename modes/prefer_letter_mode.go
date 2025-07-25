package modes

import (
	"strings"

	"github.com/jan-re/sovorel/utils"
)

type PreferLetterMode struct {
	GameCore
	PreferredLetter string
}

func (m *PreferLetterMode) PlayRound() bool {
	var wordWithLetter *utils.Word
	for i, word := range m.Words[m.index:] {
		if strings.Contains(word.Armenian, m.PreferredLetter) {
			wordWithLetter = &word
			m.index = i
			break
		}
	}

	if wordWithLetter == nil {
		return false
	}

	wasCorrect := playTranslationGame(m.Reader, wordWithLetter.Armenian, wordWithLetter.English)

	m.score.Increment(wasCorrect)

	// Stop the game if we've reached the last word.
	if m.index == len(m.Words)-1 {
		return false
	}

	m.index++
	return true
}

func (m *PreferLetterMode) GetScore() utils.Score {
	return m.score
}
