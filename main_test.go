package bcrypt_benchmark

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func BenchmarkEncryption(B *testing.B) {

	testPassword := "Z1BMflfqmfwJz8sdzqp9Pnc3DYGAG35sn58ZSbALUHggrPuppxNFHHQZEjZ9Lep6"
	hash8, _ := bcrypt.GenerateFromPassword([]byte(testPassword), 8)
	hash4, _ := bcrypt.GenerateFromPassword([]byte(testPassword), 4)
	hash6, _ := bcrypt.GenerateFromPassword([]byte(testPassword), 6)

	B.ResetTimer()
	B.Run("bcrypt_8", func(b *testing.B) {
		B.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			VerifyPassword(testPassword, string(hash8))
		}

	})

	B.Run("bcrypt_6", func(b *testing.B) {
		B.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			VerifyPassword(testPassword, string(hash6))
		}

	})

	B.Run("bcrypt_4", func(b *testing.B) {
		B.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			VerifyPassword(testPassword, string(hash4))
		}

	})
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
