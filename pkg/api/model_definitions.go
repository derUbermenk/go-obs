package api

// initial definition, will add other fields later.
type User struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPass"`
}

type Bid struct {
	ID     int `json:"id"`
	Amount int `json:"amount"`
}

type Bidding struct{}
