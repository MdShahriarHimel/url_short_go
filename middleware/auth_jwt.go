package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"url_short/config"
)

type expireTime struct {
	IssuedAt int64 `json:"iat"`
	Expire   int64 `json:"exp"`
}

func AuthJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		AuthHeaderSplit := strings.Split(r.Header.Get("Authorization"), " ")

		if len(AuthHeaderSplit) != 2 {
			http.Error(w, "Invalid Auth Header", http.StatusUnauthorized)
			return
		}

		if len(AuthHeaderSplit[1]) == 0 {
			http.Error(w, "EMPTY JWT TOKEN", http.StatusUnauthorized)
			return
		}

		jwt_secret := config.GetConfig().JWT_SECRET

		jwtTokenSplit := strings.Split(AuthHeaderSplit[1], ".")
		if len(jwtTokenSplit) != 3 {
			http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		base64Header := jwtTokenSplit[0]
		base64Payload := jwtTokenSplit[1]
		jwt_signature := jwtTokenSplit[2]

		bytePayload, err := base64decoding(base64Payload)
		if err != nil {
			http.Error(w, "Oops! Something Went Wrong!", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		var expireStruct expireTime
		err = json.Unmarshal(bytePayload, &expireStruct)

		if err != nil {
			http.Error(w, "Oops! Something Went Wrong!", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if checkExpiration(expireStruct) {
			http.Error(w, "JWT SESSION EXPIRED!", http.StatusUnauthorized)
			return
		}

		jwt_created_signature := JwtHelper(base64Header, base64Payload, jwt_secret)

		if hmac.Equal([]byte(jwt_created_signature), []byte(jwt_signature)) == false {
			http.Error(w, "NOT VALID JWT", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}

func JwtHelper(base64urlHeader string, base64urlPayload string, jwt_secret string) string {
	message := base64urlHeader + "." + base64urlPayload

	h := hmac.New(sha256.New, []byte(jwt_secret))
	h.Write([]byte(message))
	byteHash := h.Sum(nil)

	jwt_created_signature := base64encoding(byteHash)

	return jwt_created_signature
}

func base64encoding(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
func base64decoding(data string) ([]byte, error) {
	decodedByte, err := base64.RawURLEncoding.DecodeString(data)

	if err != nil {
		return nil, err
	}

	return decodedByte, nil
}

func checkExpiration(jwtTime expireTime) bool {
	if time.Now().Unix() > jwtTime.Expire {
		return true
	}

	return false
}
