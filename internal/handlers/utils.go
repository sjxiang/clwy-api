package handlers


import (
	"math/rand"
	"strings"
	"time"
	"errors"
 
	"golang.org/x/crypto/bcrypt"
)


var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)


var r *rand.Rand

func init() {
	// 随机种子
	r = rand.New(rand.NewSource(time.Now().Unix()))
}


func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	
	b := make([]byte, length)
	
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	
	return string(b)
}

func RandomString(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func GeneratePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}


func ComparePasswordAndHash(password string, hash []byte) (bool, error) {

	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		} else {
			return false, ErrInvalidCredentials
		}
	} 
	
	return true, nil
}

