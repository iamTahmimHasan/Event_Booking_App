package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // Linux, macOS, etc.
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func login() {
	fmt.Println("Login function called")
}

func createEvent() {
	fmt.Println("Create Event function called")
}
func updateEvent() {
	fmt.Println("Update Event function called")
}

func showEvents() {
	fmt.Println("Show Events function called")
}

func showBookings() {
	fmt.Println("Show Bookings function called")
}

func options() {
	fmt.Println("Options function called")
}

func main() {
	clearConsole()

	fmt.Println("Welcome to Event Booking App")
	fmt.Println(" ------------------------------")
	fmt.Println("1. Create Event")
	fmt.Println("2. Update Event")
	fmt.Println("3. Show Events")
	fmt.Println("4. Show Bookings")
	fmt.Println("5. Login")
	fmt.Println("6. Exit")
	fmt.Println(" ------------------------------")
	fmt.Println("Please select an option:")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		createEvent()
	case 2:
		updateEvent()
	case 3:
		showEvents()
	case 4:
		showBookings()
	case 5:
		login()
	case 6:
		fmt.Println("Exiting...")
		clearConsole()
		return
	default:
		fmt.Println("Invalid choice, please try again.")
	}

	options()

}
