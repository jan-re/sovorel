package modes

import (
	"bufio"

	"github.com/jan-re/sovorel/utils"
)

type GameCore struct {
	Reader *bufio.Reader
	Words  []utils.Word
	index  int
	score  utils.Score
}

func (gc *GameCore) PrintScore() {
	gc.score.Print()
}

func (gc *GameCore) finalizeRound(wasCorrect bool) bool {
	gc.score.Increment(wasCorrect)

	// Stop the game if we've reached the last word.
	if gc.index == len(gc.Words)-1 {
		return false
	}

	gc.index++
	return true
}
