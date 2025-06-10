package auth

import (
	"bufio"
	"fgo24-go-auth-flow/models"
	"fmt"
	"os"
	"strings"
)

func Register(users *[]models.User) {
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

	encrypted := EncryptPassword(password)

	newUser := models.User{
		Username: username,
		Email:    email,
		Password: encrypted,
	}

	*users = append(*users, newUser)
	fmt.Println("Registrasi berhasil")
	
}



