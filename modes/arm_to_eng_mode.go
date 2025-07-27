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

	return m.finalizeRound(wasCorrect)
}
