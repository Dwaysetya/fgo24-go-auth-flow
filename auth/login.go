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
	models.ClearScreen()

	fmt.Println("\n=== LOGIN ===")

	for i := 1; i <= 3; i++ {
		fmt.Print("Email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)
		models.ClearScreen()

		for _, user := range users {
			if user.Email == email && user.Password == EncryptPassword(password) {
				fmt.Println("\n✅ Login berhasil Selamat datang", user.Username)

				fmt.Print("Tekan ENTER untuk melanjutkan...")
				reader.ReadString('\n')
				models.ClearScreen()

				showLoginMenu(users)
				return
			}
		}
		fmt.Println("Email atau password salah.")
	}
	fmt.Println("Login gagal")
	os.Exit(1)
}

func showLoginMenu(users []models.User) {
	for {
		fmt.Println(`
=== MENU LOGIN ===
1. Lihat daftar user
2. Kembali
		`)
		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)
		models.ClearScreen()

		switch pilihan {
		case "1":
			fmt.Println("\n=== DAFTAR USER TERDAFTAR ===")
			for i, u := range users {
				fmt.Printf("%d. %s (%s)\n", i+1, u.Username, u.Email)
			}
		case "2":
			fmt.Println(" Logout berhasil. Kembali ke menu utama...")
			return
		default:
			fmt.Println(" Pilihan tidak valid.")
		}
	}
}

