package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Success() {
	fmt.Println("\n🎉 Login berhasil!")
	fmt.Println("1. Kembali ke menu utama")
	fmt.Println("0. Keluar")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pilih opsi: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			return
		case "0":
			fmt.Println("Keluar...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
