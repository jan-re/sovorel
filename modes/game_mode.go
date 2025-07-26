package modes

type GameMode interface {
	PlayRound() bool
	PrintScore()
}
