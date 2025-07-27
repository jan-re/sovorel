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

	return m.finalizeRound(wasCorrect)
}
