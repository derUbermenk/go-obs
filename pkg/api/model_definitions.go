package api

// initial definition, will add other fields later.
type User struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPass"`
}

type Bidding struct{}
