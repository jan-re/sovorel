package modes

import "github.com/jan-re/sovorel/utils"

type EngToArmMode struct {
	GameCore
}

func (m *EngToArmMode) PlayRound() bool {
	word := m.Words[m.index]

	wasCorrect := playTranslationGame(m.Reader, word.English, word.Armenian)

	m.score.Increment(wasCorrect)

	// Stop the game if we've reached the last word.
	if m.index == len(m.Words)-1 {
		return false
	}

	m.index++
	return true
}

func (m *EngToArmMode) GetScore() utils.Score {
	return m.score
}
