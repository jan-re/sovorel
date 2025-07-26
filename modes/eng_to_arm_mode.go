package modes

import "fmt"

type EngToArmMode struct {
	GameCore
}

func (m *EngToArmMode) PlayRound() bool {
	word := m.Words[m.index]

	wasCorrect, err := playTranslationGame(m.Reader, word.English, word.Armenian)
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
