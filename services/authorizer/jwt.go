package authorizer

import (
	"errors"
	"fmt"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	uuid "github.com/satori/go.uuid"
)

const issuer = "scores.spacerabbits.com"
const expiration = 24 * 30 * time.Hour

// GetAccessToken generates a new JWT for the user with id
func GetAccessToken(id string) ([]byte, error) {
	privKey, err := getPrivKey()

	if err != nil {
		fmt.Printf("Error getAccT: %v", err)
		return nil, err
	}

	claims := jws.Claims{}
	claims.SetSubject(id)
	claims.SetIssuer(issuer)
	claims.SetIssuedAt(time.Now())
	claims.SetExpiration(time.Now().Add(expiration))
	claims.SetJWTID(uuid.NewV4().String())

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS512)

	serialized, err := jwt.Serialize(privKey)

	return serialized, err
}

func GetPlayerID(s string) (string, error) {
	jwt, err := jws.ParseJWT([]byte(s))
	claims := jwt.Claims()

	if err != nil {
		return "", err
	}

	pubkey, err := getPubKey()

	if err != nil {
		fmt.Printf("Error: %v", err)
		return "", err
	}

	err = jwt.Validate(pubkey, crypto.SigningMethodRS512)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return "", err
	}

	userID, ok := claims.Subject()

	if !ok {
		return "", errors.New("Could not get Subject")
	}

	return userID, nil
}
