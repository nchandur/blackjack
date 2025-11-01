package game

type GameManager interface {
	Save(game Game) error
	Load(gameID int) (Game, error)
}
