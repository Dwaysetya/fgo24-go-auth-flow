package main

import (
	"bufio"
	"fgo24-go-auth-flow/login"
	"fgo24-go-auth-flow/register"
	"fgo24-go-auth-flow/user"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer fmt.Println("\nTerima kasih telah menggunakan program ini 😊")

	var users []user.User
	isRegistered := false
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		if !isRegistered {
			fmt.Println("1. Register")
		}
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			if !isRegistered {
				register.Register(&users)
			} 
		case "2":
			if len(users) == 0 {
				fmt.Println("Data tidak terdaftar")
			} else {
				login.Login(users)
			}
		case "0":
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
