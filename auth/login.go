package auth

import (
	"bufio"
	"fgo24-go-auth-flow/models"
	"fmt"
	"os"
	"strings"
)

func Login(users []models.User) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n=== LOGIN ===")

	for i := 1; i <= 3; i++ {
		fmt.Print("Email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		encrypted := EncryptPassword(password)

		// Cek ke seluruh users
		for _, user := range users {
			if user.Email == email && user.Password == encrypted {
				fmt.Println("Login berhasil! Selamat datang", user.Username)
				return
			}
		}

		fmt.Println("Login gagal! Email atau password salah.")
	}

	fmt.Println("Gagal login setelah 3 kali percobaan.")
	os.Exit(1)
}

