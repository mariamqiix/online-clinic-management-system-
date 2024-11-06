package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func paitent() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		fmt.Println("Enter Your Choice:")
		fmt.Println("				1) book an appointment")
		fmt.Println("				2) update information")
		fmt.Println("				3) show appointments")
		scanner.Scan()          // Read the input from the user
		casee := scanner.Text() // Get the user's input as a string

		switch casee {
		case "1":
			bookAnAppointment()
		case "2":
			ModefiPaitentInformation()
		case "3":
			showPatientAppointments()
		default:
			fmt.Println("Invalid option. Please try again.")
		}

	}
}

func showPatientAppointments() {
	appointments, _ := GetAppointmentsByUser(LogedUser.Id)
	for i := 0; i < len(appointments); i++ {
		user, _ := GetUserById(strconv.Itoa(appointments[i].doctor_id))
		fmt.Println("doctor Name : " + user.first_name + " " + user.last_name)
		fmt.Println("date : " + strings.ReplaceAll(appointments[i].date_of_Appointment, "T00:00:00Z", ""))
		fmt.Println("The Time : " + appointments[i].TheTime)

		if i != len(appointments)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
	fmt.Println("\n\n               **enter anuthing to return")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	casee := scanner.Text()
	if casee != "" {
		return
	}
}

func BeforeReturn() {
	fmt.Println("\n\n               **enter anuthing to return")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	casee := scanner.Text()
	if casee != "" {
		return
	}
}

func ModefiPaitentInformation() {
	fmt.Println("				*Select what to update:")
	fmt.Println("		1) birth Day")
	fmt.Println("		2) password")
	fmt.Println("		3) email")
	fmt.Println("		4) id")
	fmt.Println("		5) first name and last name ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	casee := scanner.Text()
	switch casee {
	case "1":
		fmt.Println("		Enter the new birth day ")
		scanner.Scan()
		bd := scanner.Text()
		err := updateUserBirthDay(strconv.Itoa(LogedUser.Id), bd)
		if err != nil {
			fmt.Print(err)
			return
		}
	case "2":
		fmt.Println("		Enter the new password ")
		scanner.Scan()
		pass := scanner.Text()
		err := updateUserEmail(strconv.Itoa(LogedUser.Id), pass)
		if err != nil {
			fmt.Print(err)
			return
		}
	case "3":
		fmt.Println("		Enter the new email ")
		scanner.Scan()
		email := scanner.Text()
		err := updatePassword(strconv.Itoa(LogedUser.Id), email)
		if err != nil {
			fmt.Print(err)
			return
		}
	case "4":
		fmt.Println("		Enter the new Id  ")
		scanner.Scan()
		Uid := scanner.Text()
		useer, _ := GetUserByUsername(Uid)
		if useer != nil {
			fmt.Println("this user with this id is already exist")
			return
		}
		err := updateUserID(strconv.Itoa(LogedUser.Id), Uid)
		if err != nil {
			fmt.Print(err)
			return
		}
	case "5":
		fmt.Println("		Enter the new first name ")
		scanner.Scan()
		Fname := scanner.Text()
		fmt.Println("		Enter the new LAST name ")
		scanner.Scan()
		lname := scanner.Text()
		err := updateUserFirstName(LogedUser.Id, Fname)
		if err != nil {
			fmt.Print(err)
			return
		}
		err = updateUserlasttName(strconv.Itoa(LogedUser.Id), lname)
		if err != nil {
			fmt.Print(err)
			return
		}
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func bookAnAppointment() {
	scanner := bufio.NewScanner(os.Stdin)

	clearScreen()
	fmt.Println("				*Enter The doctor's Gender")
	scanner.Scan()
	Gender := scanner.Text()
	doctors, _ := GetUsersByGender(Gender, "doctor")
	fmt.Println(doctors)
	fmt.Println("				*Enter The date 'YYYY-MM-DD':")
	scanner.Scan()
	TheDate := scanner.Text()
	formattedDate := TheDate
	fmt.Println("				*Enter The time 'HH:MM':")
	scanner.Scan()
	FormattedTime := scanner.Text()
	var doctor_id int
	for i := 0; i < len(doctors); i++ {
		appointment, _ := GetAppointmentsByDatrAndTimeAndDoctor(formattedDate, FormattedTime, doctors[i].Id)
		if appointment == nil {
			doctor_id = doctors[i].Id
			appoinment := Appointment{
				doctor_id:           doctor_id,
				user_id:             LogedUser.Id,
				date_of_Appointment: formattedDate,
				TheTime:             FormattedTime,
			}
			err := CreateAppointment(appoinment)
			if err != nil {
				fmt.Print("there is smth rong in the db")
				scanner.Scan()
				casee := scanner.Text()
				if casee != "" {
					return
				}
			} else {
				fmt.Print("Appointments addedd succsessfuly")
				scanner.Scan()
				casee := scanner.Text()
				if casee != "" {
					return
				}
			}
		}
	}
	fmt.Print("could not add appointment in this time with a " + Gender + " doctor , try another time or gende")
	scanner.Scan()
	casee := scanner.Text()
	if casee != "" {
		return
	}
}
