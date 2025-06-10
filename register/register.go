package register

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fgo24-go-auth-flow/user"
	"fmt"
	"os"
	"strings"
)

func Register(users *[]user.User) {
	reader := bufio.NewReader(os.Stdin)

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

	encrypted := encryptPassword(password)

	newUser := user.User{
		Username: username,
		Email:    email,
		Password: encrypted,
	}

	*users = append(*users, newUser)
	fmt.Println("Registrasi berhasil")
}

func encryptPassword(password string) string {
	encrypt := md5.Sum([]byte(password))
	return hex.EncodeToString(encrypt[:])
}
