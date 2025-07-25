package modes

import (
	"math/rand/v2"

	"github.com/jan-re/sovorel/utils"
)

type ShuffleComboMode struct {
	GameCore
}

func (m *ShuffleComboMode) PlayRound() bool {
	word := m.Words[m.index]

	n := rand.IntN(2)

	var wasCorrect bool
	if n == 0 {
		wasCorrect = playTranslationGame(m.Reader, word.English, word.Armenian)
	} else {
		wasCorrect = playTranslationGame(m.Reader, word.Armenian, word.English)
	}

	m.score.Increment(wasCorrect)

	// Stop the game if we've reached the last word.
	if m.index == len(m.Words)-1 {
		return false
	}

	m.index++
	return true
}

func (m *ShuffleComboMode) GetScore() utils.Score {
	return m.score
}
