package token

import "time"

// interface for making token
type Maker interface {
	//create token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	//check is the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
