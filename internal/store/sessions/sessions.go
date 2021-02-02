package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"trm/internal/model"
)

//Sessions - глобальное хранилище сессий
var Sessions map[string]*model.User = make(map[string]*model.User)

//GetUniqSessionID - генератор ID новой сессии (проверка на уникальность)
func GetUniqSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
