package userdashboard

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// GetUserIDFromSession gets the user ID from the session.
func GetUserIDFromSession(r *http.Request) (int, error) {
	session := GetSession(r)
	userID, err := strconv.Atoi(session.Get("user_id"))
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// ValidateUserToken validates the user token.
func ValidateUserToken(r *http.Request) (bool, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return false, nil
	}
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}
	return false, nil
}

// GetSession gets the session from the request.
func GetSession(r *http.Request) *session {
	session, err := r.Cookie("session_id")
	if err != nil {
		log.Println(err)
		return nil
	}
	return &session
}