package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func server() {
	clearScreen()
	fmt.Println("Hello there! Select an option:")
	fmt.Println("								1) LogIn")
	fmt.Println("								2) SignUp")
	fmt.Println("								3) exit")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()          // Read the input from the user
	casee := scanner.Text() // Get the user's input as a string

	switch casee {
	case "1":
		// Handle LogIn option
		fmt.Println("LogIn selected.")
		if logIn() {
			fmt.Println("CorrectPassword.")
			System()
		}
	case "2":
		// Handle SignUp option
		fmt.Println("SignUp selected.")
		if SignUp() {
			fmt.Println("Nice.")
			System()
		}
	case "3":
		// Handle LogOut option
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func System() {
	switch LogedUser.Rule {
	case "admin":
		Admin() // done
	case "patient":
		paitent() //done
	case "doctor":
		doctor() // done
	case "Lab Techment":
		LabTech() // done
	default:
		fmt.Println("Who Are You!")
	}
}

func AddUser(Rule string) bool {
	clearScreen()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("				*Enter the " + Rule + "'s' Id:")
	scanner.Scan()
	userId := scanner.Text()
	fmt.Println("				*Enter the " + Rule + "'s' First Name:")
	scanner.Scan()
	FirstName := scanner.Text()
	fmt.Println("				*Enter the " + Rule + "'s' last Name:")
	scanner.Scan()
	lastName := scanner.Text()
	fmt.Println("				*Enter the " + Rule + "'s' Gender 'Female/Male' :")
	scanner.Scan()
	Gender := scanner.Text()
	fmt.Println("				*Enter the " + Rule + "'s' Birth date 'YYYY-MM-DD':")
	scanner.Scan()
	BirthDate := scanner.Text()
	fmt.Println("				*Enter the " + Rule + "'s' email:")
	scanner.Scan()
	email := scanner.Text()
	fmt.Println("				*Enter the " + Rule + "'s' Password:")
	Password := readPassword()
	hashedPassword, err := hashPassword(Password)
	user, err := GetUserByUsername(userId)
	if user != nil {
		fmt.Println("this user with the id " + userId + " already exiset")
		return false
	}
	TemUser := User{
		UserId:          userId,
		first_name:      FirstName,
		last_name:       lastName,
		Gender:          Gender,
		date_of_birth:   BirthDate,
		email:           email,
		hashed_password: hashedPassword,
		Rule:            Rule}

	err = CreateUser(TemUser)
	if err != nil {
		return false
	}

	return true
}
