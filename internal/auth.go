package internal

import (
	"encoding/json"
	"net/http"
	"time"
	"watchman/schema"
	"watchman/utils"

	"github.com/golang-jwt/jwt/v5"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	utils.HandleMethodNotAllowed(w, r, http.MethodPost)

	config := utils.ReadConfig()

	var user schema.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.HandleError(w, r, http.StatusBadRequest, "Error decoding JSON: ", nil)
	}

	if user.Username != config.Admin.Username || user.Password != config.Admin.Password {
		response := schema.ResponseType{
			Status:    "ERROR",
			Message:   "Invalid credentials",
			RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		}
		w.WriteHeader(http.StatusUnauthorized)
		utils.SendResponse(w, r, response)
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &schema.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		utils.HandleError(w, r, http.StatusInternalServerError, "Error signing token: ", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		Expires:  expirationTime,
		MaxAge:   1800,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	response := schema.ResponseType{
		Status:    "OK",
		Message:   "Login successful",
		RequestID: r.Context().Value(schema.RequestIDKey{}).(string),
		Data: map[string]string{
			"token":      tokenString,
			"expires_at": expirationTime.String(),
		},
	}
	utils.SendResponse(w, r, response)
}
