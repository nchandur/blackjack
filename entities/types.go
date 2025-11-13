package entities

type Stats struct {
	Played int `json:"played"`
	Won    int `json:"won"`
	Lost   int `json:"lost"`
	Pushed int `json:"pushed"`
}
