package token

import "time"

// it is an interface for managing tokens
type Maker interface {
	// This method creates a new token for specific username
	CreateToken(username string, duration time.Duration) (string, error)

	// This method verifies the token and returns the payload if varified successfully
	VerifyToken(token string) (*Payload, error)
}
