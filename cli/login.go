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

	var username, password string

	fmt.Println("---Login Menu---")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Print("Select an option: ")

	option, _ := consoleReader.ReadString('\n')
	option = sanitizeInput(option)

	switch option {
	case "1":
		fmt.Print("Enter your username: ")
		username, _ = consoleReader.ReadString('\n')
		username = sanitizeInput(username)

		fmt.Print("Enter your password: ")
		password, _ = consoleReader.ReadString('\n')
		password = sanitizeInput(password)

		if login(username, password) {
			// Redirect to the main menu or perform other actions
			MainMenu(username)
		} else {
			fmt.Println("Invalid username or password. Please try again.")
			LoginMenu()
		}
	case "2":
		fmt.Print("Enter your username: ")
		username, _ = consoleReader.ReadString('\n')
		username = sanitizeInput(username)

		fmt.Print("Enter your password: ")
		password, _ = consoleReader.ReadString('\n')
		password = sanitizeInput(password)

		if register(username, password) {
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

func login(username, password string) bool {
	// Query the database to check if the provided username and password match any user
	var user entity.User
	err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		// Error occurred or no matching user found
		return false
	}

	// Login successful
	return true
}

func register(username, password string) bool {
	// Create a new User object with the given username and password
	newUser := entity.User{
		Username: username,
		Password: password,
	}

	// Insert the new User into the database
	err := config.DB.Create(&newUser).Error
	if err != nil {
		ErrorHandler(err.Error(), username)
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
