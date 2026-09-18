package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
	"url_short/config"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	SUB      int    `json:"sub"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	IssuedAt int64  `json:"iat"`
	Expire   int64  `json:"exp"`
}

func CreateJwt(payload Payload) (*string, error) {
	headerStruct := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteHeader, err := json.Marshal(headerStruct)

	if err != nil {
		return nil, errors.New("Failed to create byte array of jwt header")
	}

	// time
	payload.IssuedAt = time.Now().Unix()
	payload.Expire = time.Now().Add(time.Duration(config.GetConfig().JWT_ACCESS_TOKEN_LIFE_TIME) * time.Minute).Unix()

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("Failed to create byte array of jwt payload")
	}

	message := base64encoding(byteHeader) + "." + base64encoding(bytePayload)

	jwt_secret := config.GetConfig().JWT_SECRET

	h := hmac.New(sha256.New, []byte(jwt_secret))

	_, err1 := h.Write([]byte(message))
	if err1 != nil {
		return nil, errors.New("Failed to create JWT signature")
	}

	byteSignature := h.Sum(nil)

	base64Signature := base64encoding(byteSignature)

	jwtToken := message + "." + base64Signature

	return &jwtToken, nil
}

func base64encoding(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
