package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username string
	Email    string
	Password string
}

func (u *User) VerifLogin(email, password string) bool {
	return u.Email == email && u.Password == encryptPassword(password)
}

func encryptPassword(password string) string {
	encrypt := md5.Sum([]byte(password))
	return hex.EncodeToString(encrypt[:])
}

func AddUser(users *[]User, user User) {
	*users = append(*users, user)
}

func main() {
	defer fmt.Println("\nTerima kasih telah menggunakan program ini 😊")

	var users []User
	reader := bufio.NewReader(os.Stdin)
	isRegister := false

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		if !isRegister {
			fmt.Println("1. Register")
		}
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "1" && !isRegister {
			fmt.Println("\n=== REGISTER ===")
			fmt.Print("Username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			encrypt := encryptPassword(password)
			newUser := User{
				Username: username,
				Email:    email,
				Password: encrypt,
			}

			AddUser(&users, newUser)
			isRegister = true
			fmt.Println("Registrasi berhasil")
		} else if input == "2" {
			if len(users) == 0 {
				fmt.Println("email dan Password tidak terdaftar")
				continue
			}

			fmt.Println("\n=== LOGIN ===")

			for i := 1; i <= 3; i++ {

				fmt.Print("Email: ")
				email, _ := reader.ReadString('\n')
				email = strings.TrimSpace(email)

				fmt.Print("Password: ")
				password, _ := reader.ReadString('\n')
				password = strings.TrimSpace(password)

				user := users[0] 

				if user.VerifLogin(email, password) {
					fmt.Println("Login berhasil elamat datang,", user.Username)
					return
				} else {
					fmt.Println("Login gagal email atau password salah.")
				}
			}

			fmt.Println("\nlogin gagal, password salah.")
			os.Exit(1)
		} else if input == "0" {
			fmt.Println("Keluar dari program...")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
