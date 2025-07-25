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
