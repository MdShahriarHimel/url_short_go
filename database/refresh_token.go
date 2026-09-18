package database

import (
	"strings"
)

type RefreshToken struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	TokenHash string `json:"refresh_token"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
	Revoked   bool   `json:"revoked"`
}

var refreshTokenList []RefreshToken

func (r *RefreshToken) Save() {
	if len(refreshTokenList) == 0 {
		r.ID = 1
	} else {
		r.ID = refreshTokenList[len(refreshTokenList)-1].ID + 1
	}

	refreshTokenList = append(refreshTokenList, *r)

	// fmt.Println(refreshTokenList)

}

func (r *RefreshToken) IsValid() bool {
	for i := range refreshTokenList {
		if strings.Compare(refreshTokenList[i].TokenHash, r.TokenHash) == 0 && !refreshTokenList[i].Revoked {
			r.ID = refreshTokenList[i].ID
			r.UserId = refreshTokenList[i].UserId
			r.TokenHash = refreshTokenList[i].TokenHash
			// fmt.Println(r)
			return true
		}
	}
	return false
}

func (r *RefreshToken) RevokeRefreshToken() {
	for i := range refreshTokenList {
		if r.UserId == refreshTokenList[i].UserId {
			refreshTokenList[i].Revoked = true
		}
	}
}
