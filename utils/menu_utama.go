package utils

import (
	"fgo24-go-auth-flow/auth"
	"fgo24-go-auth-flow/models"
	"fmt"
	"os"
)


var users []models.User
var isRegister bool = false

func MenuUtama() {
	for {
		fmt.Println("\n=== MENU UTAMA ===")
		if !isRegister {
			fmt.Println("1. Register")
		}
		fmt.Println(`2. Login
3. Lupa password


0. Keluar
		`)

		fmt.Print("Pilih menu: ")
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			if !isRegister {
				auth.Register(&users)
				isRegister = true
			} else {
				fmt.Println("Sudah terdaftar.")
			}
		case "2":
			if len(users) == 0 {
				fmt.Println("Data tidak ditemukan")
			} else {
				auth.Login(users)
			}
		case "3":
			if len(users) == 0 {
				fmt.Println("Data tidak ditemukan")
			}else {
				auth.ForgotPassword(&users)
			}
		case "0":
			fmt.Println("Keluar dari program...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}