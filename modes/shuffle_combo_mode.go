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

	return m.finalizeRound(wasCorrect)
}
