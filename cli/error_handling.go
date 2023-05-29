package cli

import (
	"fmt"
	"os"
)

func ErrorHandler(msg string, username string) {
	fmt.Println("terjadi kesalahan dalam aplikasi")
	fmt.Println(msg)

	var input string
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	switch input {
	case "m":
		MainMenu(username)
	case "q":
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		ErrorHandler(msg, username)
	}
}
