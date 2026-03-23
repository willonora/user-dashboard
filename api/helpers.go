package user_dashboard

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v9/v9"
)

func generateRandomString() string {
	max := big.NewInt(int64(65536))
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	n := big.NewInt(0).SetBytes(bytes)
	n.Mod(n, max)
	return hex.EncodeToString(n.Bytes())
}

func generateJWTToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(30 * time.Minute).Unix(),
		"sub":   user.ID,
		"email": user.Email,
	})
	return token.SignedString([]byte("secret"))
}

func verifyJWTToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error {
		_, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error {
			return []byte("secret"), nil
		}))
		return err
	})
}

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func sendSMS(phoneNumber string, message string) error {
	resp, err := http.Post("https://api.example.com/sms", "application/json", strings.NewReader(`{"message": "`+message+`","phone_number": "`+phoneNumber+`"}`))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}