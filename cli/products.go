package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tugasmeilyanto/go-trial-class/config"
	"github.com/tugasmeilyanto/go-trial-class/entity"
	"github.com/tugasmeilyanto/go-trial-class/helpers"
)

func ListProduct(username string) {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)

	var products []entity.Product
	err := config.DB.Find(&products).Error

	if err != nil {
		ErrorHandler(err.Error(), username)
		return
	}

	fmt.Println("---List Product---")
	for _, product := range products {
		product.PrintDetail()
	}

	var input string
	fmt.Println("Masukan Id Product untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, err = consoleReader.ReadString('\n')
	switch input {
	case "m\n":
		MainMenu(username)
	case "q\n":
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		OrderProduct(input, username)
	}
}

func OrderProduct(id string, username string) {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)

	var product entity.Product
	err := config.DB.Where("ID = ?", strings.TrimSpace(id)).First(&product).Error
	if err != nil {
		ErrorHandler(err.Error(), username)
		return
	}

	product.PrintDetail()

	var input string
	fmt.Println("Tekan (y) untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, _ = consoleReader.ReadString('\n')

	switch strings.TrimSpace(input) {
	case "y":
		CreateOrder(product, username)
	case "m":
		MainMenu(username)
	case "q":
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		OrderProduct(input, username)
	}
}

func CreateOrder(product entity.Product, username string) {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)
	var email string
	var address string
	fmt.Println("Masukan email anda")
	email, _ = consoleReader.ReadString('\n')

	fmt.Println("Masukan alamat anda : ")
	address, _ = consoleReader.ReadString('\n')

	order := entity.Order{
		ProductId:    int(product.ID),
		BuyerEmail:   email,
		BuyerAddress: address,
		OrderDate:    time.Now(),
	}

	err := config.DB.Create(&order).Error
	if err != nil {
		ErrorHandler(err.Error(), username)
		return
	}

	fmt.Println("Order berhasil dibuat")

	var input string
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err = fmt.Scanln(&input)
	if err != nil {
		MainMenu(username)
	}
	switch input {
	case "q\n":
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		MainMenu(username)
	}
}
