package game

type GameManager interface {
	Save(game Game) error
	Load() (Game, error)
}
