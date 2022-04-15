package utils

import (
	"HLRJ/gin_learn/model"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
func RandomString(n int) string {
	var (
		letters = []byte("qwertyuiopassdfghjklzxcvvbbnm")
	)
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
