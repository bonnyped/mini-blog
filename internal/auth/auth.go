package auth

import (
	"encoding/json"
	"log/slog"
	"mini-blog/pkg/sl"

	"github.com/go-chi/jwtauth/v5"
)

type JWTManager struct {
	*jwtauth.JWTAuth
}

type AccessToken struct {
	JWTToken string `json:"access_token"`
}

func (jwt JWTManager) getAccessToken(userID int) (string, error, string) {
	const op = "internal.auth.getAccessToken"
	claims := map[string]any{"user_id": userID}

	_, accessToken, err := jwt.Encode(claims)

	return accessToken, err, op
}

func (jwt JWTManager) GetMarshaledToken(logger *slog.Logger, userID int) ([]byte, error) {
	const op = "internal.auth.GetMarshaledToken"

	accessToken, err, cause := jwt.getAccessToken(userID)
	if err != nil {
		return nil, sl.Err(cause, err)
	}

	marshaledToken, err := json.Marshal(AccessToken{JWTToken: accessToken})
	if err != nil {
		return nil, sl.Err(op, err)
	}

	return marshaledToken, nil
}
