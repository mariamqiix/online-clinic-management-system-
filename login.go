package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

var LogedUser *User

func logIn() bool {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("				*Enter Your Id/Email:")
	scanner.Scan()
	username := scanner.Text()

	fmt.Println("				*Enter Your Password:")
	password := readPassword()

	// Use GetUserByUsername to get the user
	user, err := GetUserByUsername(username)
	if err != nil || user == nil {
		user2, err := GetUserByEmail(username)
		if err != nil || user2 == nil {
			// Handle error
			return false
		}
		// Compare the hashed password with the given password
		err = bcrypt.CompareHashAndPassword([]byte(user2.hashed_password), []byte(password))
		if err != nil {
			return false
		} else {
			LogedUser = user2
			return true
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.hashed_password), []byte(password))
	if err != nil {
		return false
	}

	LogedUser = user

	// If everything is okay, return true
	return true
}

func SignUp() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("				*Enter Your Id:")
	scanner.Scan()
	userId := scanner.Text()
	fmt.Println("				*Enter Your First Name:")
	scanner.Scan()
	FirstName := scanner.Text()
	fmt.Println("				*Enter Your last Name:")
	scanner.Scan()
	lastName := scanner.Text()
	fmt.Println("				*Enter Your Gender 'Female/Male' :")
	scanner.Scan()
	Gender := scanner.Text()
	fmt.Println("				*Enter Your Birth date 'DD-MM-YYYY':")
	scanner.Scan()
	BirthDate := scanner.Text()
	fmt.Println("				*Enter Your email:")
	scanner.Scan()
	email := scanner.Text()
	fmt.Println("				*Enter Your Password:")
	Password := readPassword()
	hashedPassword, err := hashPassword(Password)
	user, err := GetUserByUsername(userId)
	if user != nil {
		fmt.Print("this user with the id" + userId + " already exiset")
		return false
	}
	TemUser := User{
		UserId:          userId,
		first_name:      FirstName,
		last_name:       lastName,
		Gender:          Gender,
		date_of_birth:   FormatDate(BirthDate),
		email:           email,
		hashed_password: hashedPassword,
		Rule:            "patient"}

	err = CreateUser(TemUser)
	if err != nil {
		return false
	}

	LogedUser = &TemUser

	// If everything is okay, return true
	return true
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func readPassword() string {
	bytePassword, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return string(bytePassword)
}
