package modes

import (
	"fmt"
	"math/rand/v2"
)

type ShuffleComboMode struct {
	GameCore
}

func (m *ShuffleComboMode) PlayRound() bool {
	word := m.Words[m.index]

	n := rand.IntN(2)

	var query, answer string
	if n == 0 {
		query, answer = word.English, word.Armenian
	} else {
		query, answer = word.Armenian, word.English
	}

	wasCorrect, err := playTranslationGame(m.Reader, query, answer)
	if err != nil {
		fmt.Println(logErrorReadingInput)
		return false
	}

	m.score.Increment(wasCorrect)

	// Stop the game if we've reached the last word.
	if m.index == len(m.Words)-1 {
		return false
	}

	m.index++
	return true
}
