package database

import "url_short/util"

type Link struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	ShortCode string `json:"short_code"`
	LongUrl   string `json:"long_url"`
}

var linkList []Link

func check(longUrl string, usrId int) (*Link, bool) {
	for i := range linkList {
		if linkList[i].LongUrl == longUrl && linkList[i].UserId == usrId {
			return &linkList[i], true
		}
	}

	return nil, false
}

func (link *Link) StoreShortLink() (*Link, bool) {

	l, exist := check(link.LongUrl, link.UserId)
	if exist {
		return l, false
	}

	if len(linkList) == 0 {
		link.ID = 1
	} else {
		link.ID = linkList[len(linkList)-1].ID + 1
	}

	link.ShortCode = util.Base62Encode(link.ID)

	linkList = append(linkList, *link)

	return nil, true
}

func (link *Link) GetList() []Link {
	var userLinkList []Link

	for _, links := range linkList {
		if link.UserId == links.UserId {
			userLinkList = append(userLinkList, links)
		}
	}

	return userLinkList
}

func (link *Link) GetById() *Link {
	for i := range linkList {
		if link.UserId == linkList[i].UserId && link.ID == linkList[i].ID {
			return &linkList[i]
		}
	}

	return nil
}

func (link *Link) Update() bool {

	for i := range linkList {
		if linkList[i].UserId == link.UserId && linkList[i].ID == link.ID {
			linkList[i].LongUrl = link.LongUrl
			return true
		}
	}

	return false

}

func (link *Link) Delete() bool {
	for i, l := range linkList {
		if l.ID == link.ID {
			linkList = append(linkList[:i], linkList[i+1:]...)
			return true
		}
	}
	return false
}

func (link *Link) GetLongUrl() (string, bool) {
	for i := range linkList {
		if link.UserId == linkList[i].UserId && link.ShortCode == linkList[i].ShortCode {
			return linkList[i].LongUrl, true
		}
	}

	return "", false
}
