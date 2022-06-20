package core

import (
	"authorization/custom_models"
	"authorization/database"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Core struct {
}

func (u *Core) CreateJwt(id string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	} else {
		db := database.Connect()
		defer db.Disconnect()
		tokenmodel := custom_models.Token{UserId: id, Token: tokenString, Timestamp: time.Now().Unix()}
		_, err = db.InsertToken(&tokenmodel)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
}

func (u *Core) DeleteJwt(userid string) (bool, error) {
	db := database.Connect()
	defer db.Disconnect()
	return db.FindOneAndDeleteTokenById(userid)
}

func (u *Core) ValidateJwt(token string) (string, error) {
	db := database.Connect()
	defer db.Disconnect()
	istokenexists, err := db.IsTokenExists(token)
	if err != nil {
		return "", err
	} else {
		if istokenexists {
			secretKey := os.Getenv("SECRET_KEY")
			tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("an error occurred while parsing token")
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				return "", err
			}
			if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
				return fmt.Sprint(claims["userid"]), nil
			} else {
				return "", fmt.Errorf("access denied")
			}
		} else {
			return "", fmt.Errorf("dccess denied")
		}
	}

}
