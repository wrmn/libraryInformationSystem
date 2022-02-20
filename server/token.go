package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"librarySysfo/database/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func createToken(data models.User) (t string, e error) {
	isEmployee := false
	if data.Employee.EmployeeNumber != "" {
		isEmployee = true
	}
	claims := credential{
		Id:         data.Id,
		Username:   data.Username,
		IsEmployee: isEmployee,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	if isEmployee {
		claims.Division = data.Employee.Division
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, e = token.SignedString(key)

	return
}

func createRefreshToken(data tokenCred) (token string, e error) {
	claims := jwt.MapClaims{}
	claims["token"] = data.AccessToken
	claims["exp"] = time.Now().Add(time.Hour * 24 * 90).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, e = refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if e != nil {
		return token, e
	}

	return token, nil
}

func validateToken(bearerToken string) (token *jwt.Token, err error) {
	if len(bearerToken) < 1 {
		err = errors.New("empty")
		return
	}
	tokenString := strings.Split(bearerToken, " ")[1]
	token, err = jwt.ParseWithClaims(tokenString, &credential{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return
}

func validateRefreshToken(model tokenCred, uname credential) (models.User, error) {
	token, err := jwt.Parse(model.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	user := models.User{}
	if err != nil {
		return user, err
	}
	payload, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return user, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err = parser.ParseUnverified(payload["token"].(string), claims)
	if err != nil {
		return user, err
	}

	payload, ok = token.Claims.(jwt.MapClaims)
	if !ok {
		return user, errors.New("invalid token")
	}

	user.Username = payload["username"].(string)
	if user.Username != uname.Username && uname.Id != user.Id {
		err = errors.New("invalid user token for user")
		return user, err
	}

	return user, nil
}

func inspectToken(w http.ResponseWriter, r *http.Request) (token *jwt.Token, isErr bool) {
	var ve *jwt.ValidationError

	bearerToken := r.Header.Get("Authorization")
	refresh := r.Header.Get("X-Refresh-Token")
	token, err := validateToken(bearerToken)
	if err != nil && err.Error() == "empty" {
		response := responseParam{
			W:      w,
			Body:   respToByte("error", "Token Required", http.StatusUnauthorized),
			Status: http.StatusUnauthorized,
		}
		responseFormatter(response, "")
		isErr = true
		return
	}
	uname := token.Claims.(*credential)
	isErr = false

	if errors.As(err, &ve) && refresh != "" {

		creds := tokenCred{
			AccessToken:  bearerToken,
			RefreshToken: refresh,
		}

		user, err := validateRefreshToken(creds, *uname)
		if err != nil {
			response := responseParam{
				W:      w,
				Body:   respToByte("error", "Invalid refresh token", http.StatusUnauthorized),
				Status: http.StatusUnauthorized,
			}
			responseFormatter(response, "")
			isErr = true
			return
		}

		newToken, err := createToken(user)
		if err != nil {
			json.NewEncoder(w).Encode("Unable to create access token")
			isErr = true
			return
		}
		w.Header().Add("X-New-Token", newToken)
		return
	} else if err != nil {
		if err == jwt.ErrSignatureInvalid {
			unAuthorized(w)
			isErr = true
		} else if refresh == "" {
			response := responseParam{
				W:      w,
				Body:   respToByte("error", "Access Token is expire!", http.StatusUnauthorized),
				Status: http.StatusUnauthorized,
			}
			responseFormatter(response, "")
			isErr = true
			return
		} else {
			badRequest(w)
			isErr = true
		}
		return
	}
	return
}
