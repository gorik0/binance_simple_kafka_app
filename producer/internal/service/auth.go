package service

import (
	"bina/internal/core"
	"bina/internal/storage"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`

}
type AuthService struct {
	AuthRepo storage.Authorization
}

var tokenTTL = time.Hour*60

var signingKey *ecdsa.PrivateKey

func init() {
	signingKey,_=ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
}
func (a AuthService) GenerateToken(username, password string) (string, error) {
	user,err:= a.AuthRepo.GetUSer(username,GenerateHashPassword(password))
	if err != nil {
		return "",nil
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodES256,&tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId:           user.Id,
	})
	return token.SignedString(signingKey)
}

func (a AuthService) CreateUser(user *core.User) (int, error) {
	user.Password = GenerateHashPassword(user.Password)
	return a.AuthRepo.CreateUser(user)
}

var salt = "SALT"

func GenerateHashPassword(password string) string {
	hash:= sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x",hash.Sum([]byte(salt)))

}


func (a AuthService) ParseToken(token string) (int, error) {
	tok,err:= jwt.ParseWithClaims(token,&tokenClaims{}, func(*jwt.Token) (interface{}, error){
		return signingKey,nil
	})
	if err != nil {
		return 0,nil
	}
	claim,ok:= tok.Claims.(*tokenClaims)
	if !ok {
		return 0,errors.New("Claim of invalid type(not token claim)")
	}
	return claim.UserId,nil
}

func NewAuthService(s *storage.Storage) *AuthService {
	return &AuthService{
		AuthRepo:     s.Authorization  ,
	}
}
