package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func GetUserIdFromJwt(token string) (*int, error) {
	tokenSpit := strings.Split(token, ".")

	base64payload := tokenSpit[1]
	bytePayload, err := base64.RawURLEncoding.DecodeString(base64payload)
	if err != nil {
		return nil, errors.New("Could not decode base64payload to bytePayload")
	}

	var UsrId struct {
		UserId int `json:"sub"`
	}

	err = json.Unmarshal(bytePayload, &UsrId)
	if err != nil {
		return nil, errors.New("Could not convert bytePayload to usrId")
	}

	return &UsrId.UserId, nil

}
