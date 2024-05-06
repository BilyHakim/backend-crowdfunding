package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(userID int64) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("crowdfunding_bh")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int64) (string, error) {
	// Payload kadang disebut claim
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
