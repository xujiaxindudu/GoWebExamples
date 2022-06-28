package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//bcrypt:加盐值和自适应性，增加迭代次数使得加密变慢递与暴力搜索

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func main() {
	password := "dudu "
	hash, _ := HashPassword(password)
	fmt.Println("p:", password)
	fmt.Println("hash:", hash)
	match := CheckPasswordHash(password, hash)
	fmt.Println("match", match)
}
