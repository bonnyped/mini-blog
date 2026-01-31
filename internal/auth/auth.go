package auth

import (
	"encoding/json"
	"log/slog"
	"mini-blog/pkg/sl"
	"strconv"

	"github.com/go-chi/jwtauth/v5"
)

type JWTManager struct {
	*jwtauth.JWTAuth
}

type AccessToken struct {
	JWTToken string `json:"access_token"`
}

func (jm JWTManager) getAccessToken(userID int) (string, string, error) {
	const op = "internal.auth.getAccessToken"

	claims := map[string]any{"user_id": strconv.Itoa(userID)}

	_, accessToken, err := jm.Encode(claims)

	return accessToken, op, err
}

func (jm JWTManager) GetMarshaledToken(logger *slog.Logger, userID int) ([]byte, error) {
	const op = "internal.auth.GetMarshaledToken"

	accessToken, cause, err := jm.getAccessToken(userID)
	if err != nil {
		return nil, sl.Err(cause, err)
	}

	marshaledToken, err := json.Marshal(AccessToken{JWTToken: accessToken})
	if err != nil {
		return nil, sl.Err(op, err)
	}

	return marshaledToken, nil
}
