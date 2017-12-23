package authorizer

import (
	"errors"
	"fmt"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	uuid "github.com/satori/go.uuid"
	"github.com/tehsis/rabbitscore/services/players"
)

const issuer = "scores.spacerabbits.com"
const expiration = 24 * 30 * time.Hour

// GetAccessToken generates a new JWT for the user with id
func GetAccessToken(player players.Player) ([]byte, error) {
	privKey, err := getPrivKey()

	if err != nil {
		fmt.Printf("Error getAccT: %v", err)
		return nil, err
	}

	claims := jws.Claims{}
	claims.SetSubject(player.ID)
	claims.Set("name", player.Name)
	claims.SetIssuer(issuer)
	claims.SetIssuedAt(time.Now())
	claims.SetExpiration(time.Now().Add(expiration))
	claims.SetJWTID(uuid.NewV4().String())

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS512)

	serialized, err := jwt.Serialize(privKey)

	return serialized, err
}

func GetPlayer(s string) (players.Player, error) {
	jwt, err := jws.ParseJWT([]byte(s))
	claims := jwt.Claims()

	if err != nil {
		return players.Player{}, err
	}

	pubkey, err := getPubKey()

	if err != nil {
		fmt.Printf("Error: %v", err)
		return players.Player{}, err
	}

	err = jwt.Validate(pubkey, crypto.SigningMethodRS512)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return players.Player{}, err
	}

	userID, ok := claims.Subject()
	username := claims.Get("name").(string)

	if !ok {
		return players.Player{}, errors.New("Could not get Subject")
	}

	return players.Player{
		ID:       userID,
		Name:     username,
		SocialID: players.SocialPlayer{ID: "", Provider: ""},
	}, nil
}
