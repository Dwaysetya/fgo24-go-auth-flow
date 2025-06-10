package login

import (
	"bufio"
	"crypto/md5"
	"fgo24-go-auth-flow/user"
	"fmt"
	"os"
	"strings"
)

func Login(users []user.User) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== LOGIN ===")

	for i := 1; i <= 3; i++ {

		fmt.Print("Email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		if users[0].Email == email && users[0].Password == encryptPassword(password) {
			fmt.Println("Login berhasil Selamat datang", users[0].Username)
			return
		}
		fmt.Println("Login gagal Email atau password salah.")
	}
	fmt.Println("Gagal login")
	os.Exit(1)
}

func encryptPassword(password string) string {
	encrypt := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", encrypt)
}
