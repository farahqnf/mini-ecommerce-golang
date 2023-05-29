package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tugasmeilyanto/go-trial-class/config"
	"github.com/tugasmeilyanto/go-trial-class/entity"
)

func LoginMenu() {
	consoleReader := bufio.NewReader(os.Stdin)

	var email, password string

	fmt.Println("---Login Menu---")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Print("Select an option: ")

	option, _ := consoleReader.ReadString('\n')
	option = sanitizeInput(option)

	switch option {
	case "1":
		fmt.Print("Enter your email: ")
		email, _ = consoleReader.ReadString('\n')
		email = sanitizeInput(email)

		fmt.Print("Enter your password: ")
		password, _ = consoleReader.ReadString('\n')
		password = sanitizeInput(password)

		if login(email, password) {
			// Redirect to the main menu or perform other actions
			MainMenu()
		} else {
			fmt.Println("Invalid email or password. Please try again.")
			LoginMenu()
		}
	case "2":
		fmt.Print("Enter your email: ")
		email, _ = consoleReader.ReadString('\n')
		email = sanitizeInput(email)

		fmt.Print("Enter your password: ")
		password, _ = consoleReader.ReadString('\n')
		password = sanitizeInput(password)

		if register(email, password) {
			fmt.Println("Registration successful. You can now login.")
			LoginMenu()
		} else {
			fmt.Println("Registration failed. Please try again.")
			LoginMenu()
		}
	default:
		fmt.Println("Invalid option. Please try again.")
		LoginMenu()
	}
}

func login(email, password string) bool {
	// Query the database to check if the provided email and password match any user
	var customer entity.Customer
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&customer).Error
	if err != nil {
		// Error occurred or no matching user found
		return false
	}

	// Login successful
	return true
}

func register(email, password string) bool {
	// Create a new Customer object with the given email and password
	newCustomer := entity.Customer{
		Email:    email,
		Password: password,
	}

	// Insert the new customer into the database
	err := config.DB.Create(&newCustomer).Error
	if err != nil {
		ErrorHandler(err.Error())
		return false
	}

	// Return true to indicate a successful registration
	return true
}

func sanitizeInput(input string) string {
	// Remove leading/trailing whitespaces and newlines from the input
	input = strings.TrimSpace(input)
	input = strings.TrimSuffix(input, "\n")
	return input
}
