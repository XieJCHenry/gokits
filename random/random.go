package random

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

var shortLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

// NewId 获取当前微秒时间戳，插入随机小写字母
func NewId() string {
	return newIdFromLetters(letters, 13)
}

func NewShortId() string {
	return newIdFromLetters(shortLetters, 6)
}

func newIdFromLetters(letters []rune, idLength int) string {
	letterSize := len(letters)
	b := make([]rune, idLength)
	for i := range b {
		b[i] = letters[rand.Intn(letterSize)]
	}
	return string(b)
}
