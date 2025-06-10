package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Success(){
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n Anda berhasil login \n")
		fmt.Println("0. Kembali")

		fmt.Print("Pilih menu")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "0":
			fmt.Println("Keluar dari program")
		}
	}
}