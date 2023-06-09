package cli

import (
	"fmt"
	"os"

	"github.com/tugasmeilyanto/go-trial-class/helpers"
)

func MainMenu(username string) {
	helpers.ClearScreen()
	fmt.Println("Selamat Datang di Mini Ecommerce", username, "!")
	fmt.Println("--------------------------------")

	var input string
	fmt.Println("Tekan (1) untuk melanjutkan ke list product")
	fmt.Println("Tekan (2) untuk melanjutkan ke list order")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "1":
		ListProduct(username)
	case "2":
		ListOrder(username)
	case "q":
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		MainMenu(username)
	}
}
