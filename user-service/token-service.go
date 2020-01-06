package main

// TokenEncoder --
type TokenEncoder interface {
	Encode(data interface{}) (string, error)
	Decode(token string) (interface{}, error)
}

// TokenService --
type TokenService struct {
	repo repository
}

// Encode --
func (t *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}

// Decode --
func (t *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}
