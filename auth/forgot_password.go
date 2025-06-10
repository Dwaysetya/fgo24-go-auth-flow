package auth

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fgo24-go-auth-flow/models"
	"fmt"
	"os"
	"strings"
)

func ForgotPassword(users *[]models.User) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== FORGOT PASSWORD ===")
	fmt.Print("Masukkan email: ")
	emailInput, _ := reader.ReadString('\n')
	emailInput = strings.TrimSpace(emailInput)

	for i, u := range *users {
		if u.Email == emailInput {
			fmt.Print("Masukkan password baru: ")
			newPass, _ := reader.ReadString('\n')
			newPass = strings.TrimSpace(newPass)

			(*users)[i].Password = encryptPassworded(newPass)
			fmt.Println("Password berhasil direset")
			return
		}
	}

	fmt.Println("Email tidak ditemukan.")
}

func encryptPassworded(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
