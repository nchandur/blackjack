package users

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	SignedIn bool   `json:"signedIn"`
	Games    []struct {
		GameID int `json:"game_id"`
		Round  struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
			Pushed int `json:"pushed"`
		} `json:"round"`
	} `json:"games"`
}

type UserManager interface {
	SignIn(username, password string) error
	Create(username, password string) error
	Retrieve(username string) (User, error)
	Delete(username string) error
}
