package modes

import "fmt"

type ArmToEngMode struct {
	GameCore
}

func (m *ArmToEngMode) PlayRound() bool {
	word := m.Words[m.index]

	wasCorrect, err := playTranslationGame(m.Reader, word.Armenian, word.English)
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
