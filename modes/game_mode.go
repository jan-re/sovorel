package modes

import "github.com/jan-re/sovorel/utils"

type GameMode interface {
	PlayRound() bool
	GetScore() utils.Score
}
