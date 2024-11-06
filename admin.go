package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func Admin() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		fmt.Println("Enter Your Choice:")
		fmt.Println("				1) Add doctor")
		fmt.Println("				2) Add patient")
		fmt.Println("				3) Add Lab Techment")
		fmt.Println("				4) Add Appointment")
		fmt.Println("				5) Show tests")
		fmt.Println("				6) Modefi Doctor's Information")
		fmt.Println("				7) Modefi patient's Information")
		fmt.Println("				8) Modefi Lab Techment's information")
		fmt.Println("				9) modifi Medicine description")
		fmt.Println("				10) add Medicine")
		fmt.Println("				11) Generate Reports")
		fmt.Println("				12) Track Test")
		fmt.Println("				13) Show paitent's Tests")
		fmt.Println("				14) Show diseases")
		fmt.Println("				15) add diseases")
		fmt.Println("				16) modifi disease")
		fmt.Println("				17) Show paitent's Desies")

		scanner.Scan()          // Read the input from the user
		casee := scanner.Text() // Get the user's input as a string

		switch casee {
		case "1":
			AddUser("doctor")
		case "2":
			AddUser("patient")
		case "3":
			AddUser("Lab Techment")
		case "4":
			AddAppointment()
		case "5":
			ShowTests()
		case "6":
			ModefiInformation("doctor")
		case "7":
			ModefiInformation("patient")
		case "8":
			ModefiInformation("Lab Techment")
		case "9":
			ChangeMedicineDesc()
		case "10":
			addMed()
		case "11":
			// not yet
		case "12":
			TrackTest()
		case "13":
			ShowPaitentTests()
		case "14":
			ShowDiseases()
		case "15":
			addDiseases()
		case "16":
			modifiDisease()
		case "17":
			ShowPaitentDesies()
		default:
			fmt.Println("Invalid option. Please try again.")
		}

	}
}

func AddAppointment() {
	scanner := bufio.NewScanner(os.Stdin)
	clearScreen()

	fmt.Println("				*Enter patient id:")
	scanner.Scan()
	UserId := scanner.Text()
	user_id, err := GetUserNoByUserId(UserId)
	if err != nil || user_id == 0 {
		fmt.Print("this user isnt available for appointments/dose not exiset")
		BeforeReturn()
		return
	}
	fmt.Println("				*Enter The doctor's Gender")
	scanner.Scan()
	Gender := scanner.Text()
	doctors, _ := GetUsersByGender(Gender, "doctor")
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
				user_id:             user_id,
				date_of_Appointment: formattedDate,
				TheTime:             FormattedTime,
			}
			err := CreateAppointment(appoinment)
			if err != nil {
				fmt.Print("there is smth rong in the db")
				BeforeReturn()
				return
			} else {
				fmt.Print("Appointment addedd succsessfuly")
			}
			BeforeReturn()
			return
		}
	}

	fmt.Print("could not add appointment in this time with a " + Gender + "doctor , try another time or gende")

}

func clearScreen() {
	// Clear the screen command based on the operating system
	var clearCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default:
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func ShowAppointment() {
	appointments, _ := GetAppointmentsByDoctor(LogedUser.Id)
	for i := 0; i < len(appointments); i++ {
		paitent, _ := GetUserById(strconv.Itoa(appointments[i].user_id))
		fmt.Println("patient name : " + paitent.first_name + " " + paitent.last_name)
		fmt.Println("date : " + strings.ReplaceAll(appointments[i].date_of_Appointment, "T00:00:00Z", ""))
		fmt.Println("time : " + (appointments[i].TheTime))
		if i != len(appointments)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
}

func ShowTests() {
	tests, _ := GetTests()
	for i := 0; i < len(tests); i++ {
		paitent, _ := GetUserById(strconv.Itoa(tests[i].user_id))
		LabTech, _ := GetUserById(strconv.Itoa(tests[i].lab_Tech))
		doctor, _ := GetUserById(strconv.Itoa(tests[i].doctor_id))
		fmt.Sprintf("Test Id : #%d\n", tests[i].Id)
		fmt.Println("patient name : " + paitent.first_name + " " + paitent.last_name)
		fmt.Println("Lab tech name : " + LabTech.first_name + " " + LabTech.last_name)
		fmt.Println("doctor name : " + doctor.first_name + " " + doctor.last_name)
		fmt.Println("test date : " + strings.ReplaceAll(tests[i].Test_Date, "T00:00:00Z", ""))
		fmt.Println("test type : " + (tests[i].Test_Type))
		if i != len(tests)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
	if len(tests) == 0 {
		fmt.Println("no testes yet")
	}
			BeforeReturn()
		return
}

func ModefiInformation(user string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("				*Enter " + user + "'s Id:")
	scanner.Scan()
	UserId := scanner.Text()
	id, _ := GetUserNoByUserId(UserId)
	fmt.Println("				*Select what to update:")
	fmt.Println("		1) birth Day")
	fmt.Println("		2) password")
	fmt.Println("		3) email")
	fmt.Println("		4) id")
	fmt.Println("		5) first name and last name ")
	scanner.Scan()
	casee := scanner.Text()
	switch casee {
	case "1":
		fmt.Println("		Enter the new birth day ")
		scanner.Scan()
		bd := scanner.Text()
		err := updateUserBirthDay(UserId, bd)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
	case "2":
		fmt.Println("		Enter the new password ")
		scanner.Scan()
		pass := scanner.Text()
		err := updateUserEmail(UserId, pass)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
	case "3":
		fmt.Println("		Enter the new email ")
		scanner.Scan()
		email := scanner.Text()
		err := updatePassword(UserId, email)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
	case "4":
		fmt.Println("		Enter the new Id  ")
		scanner.Scan()
		Uid := scanner.Text()
		useer, _ := GetUserByUsername(Uid)
		if useer != nil {
			fmt.Println("this user with this id is already exist")
			BeforeReturn()
			return
		}
		err := updateUserID(UserId, Uid)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
	case "5":
		fmt.Println("		Enter the new first name ")
		scanner.Scan()
		Fname := scanner.Text()
		fmt.Println("		Enter the new LAST name ")
		scanner.Scan()
		lname := scanner.Text()
		err := updateUserFirstName(id, Fname)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
		err = updateUserlasttName(UserId, lname)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func ChangeMedicineDesc() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("		Enter Medicine Name  ")
	scanner.Scan()
	Mname := scanner.Text()
	fmt.Println("		Enter The new discription  ")
	scanner.Scan()
	NewDescription := scanner.Text()
	err := ChangeMedicineDescription(Mname, NewDescription)
	if err != nil {
		fmt.Print(err)
		BeforeReturn()
		return
	}
}

func addMed() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("		Enter Medicine Name  ")
	scanner.Scan()
	Mname := scanner.Text()
	fmt.Println("		Enter The Medicine discription  ")
	scanner.Scan()
	NewDescription := scanner.Text()
	med := Medicine{
		MedicineName:         Mname,
		Medicine_description: NewDescription,
	}
	err := CreateMedicine(med)
	if err != nil {
		fmt.Print(err)
		BeforeReturn()
		return
	}
}

func ShowPaitentTests() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter Patient Id :\n#  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, _ := GetUserByUsername(UsreId)
	if useer != nil {
		fmt.Println("this user with this id is not exist")
		BeforeReturn()
		return
	}
	tests, err := GetTestsByUserId(UsreId)
	if err != nil {
		fmt.Print(err)
		BeforeReturn()
		return
	}

	for i := 0; i < len(tests); i++ {
		paitent, _ := GetUserById(strconv.Itoa(tests[i].user_id))
		LabTech, _ := GetUserById(strconv.Itoa(tests[i].lab_Tech))
		doctor, _ := GetUserById(strconv.Itoa(tests[i].doctor_id))
		fmt.Sprintf("Test Id : #%d\n", tests[i].Id)
		fmt.Println("patient name : " + paitent.first_name + " " + paitent.last_name)
		fmt.Println("Lab tech name : " + LabTech.first_name + " " + LabTech.last_name)
		fmt.Println("doctor name : " + doctor.first_name + " " + doctor.last_name)
		fmt.Println("test date : " + strings.ReplaceAll(tests[i].Test_Date, "T00:00:00Z", ""))
		fmt.Println("test type : " + (tests[i].Test_Type))
		if i != len(tests)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
	if len(tests) == 0 {
		fmt.Println("no testes yet")
	}

		BeforeReturn()
		return

}

func ShowDiseases() {
	diseases, err := GetDiseases()
	if err != nil {
		fmt.Print(err)
		BeforeReturn()
		return
	}
	for i := 0; i < len(diseases); i++ {
		Medicine, _ := GetMedicineByMedicineId(diseases[i].disease_Medicine_id)
		fmt.Println("disease Name : " + diseases[i].diseaseName)
		fmt.Println("disease description : " + diseases[i].disease_description)
		fmt.Println("disease Medicine : " + Medicine)
		if i != len(diseases)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
		if len(diseases) == 0 {
		fmt.Println("no diaeases in the data base yet , try to add new one")
	}
			BeforeReturn()
		return
}
func addDiseases() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("		Enter Diseases Name :")
	scanner.Scan()
	DisName := scanner.Text()
	useer, _ := GetDiseaseByDiseaseName(DisName)
	if useer != nil {
		fmt.Println("this Disease with this id is already exist")
		BeforeReturn()
		return
	}

	fmt.Println("		Enter Diseases discription :")
	scanner.Scan()
	DisDesc := scanner.Text()

	fmt.Println("		Enter Diseases Medicine Name :")
	scanner.Scan()
	DisMed := scanner.Text()

	MedId, _ := GetMedicineByMedicineName(DisMed)
	if MedId == nil {
		fmt.Println("there is no Medicine with this name")
		BeforeReturn()
		return
	}

	NewDis := diseases{
		diseaseName:         DisName,
		disease_description: DisDesc,
		disease_Medicine_id: MedId.Id,
	}
	err := CreatDiseases(NewDis)
	if err != nil {
		fmt.Print(err)
		BeforeReturn()
		return
	}
}

func modifiDisease() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("		Enter Diseases Name :")
	scanner.Scan()
	DisName := scanner.Text()
	useer, _ := GetDiseaseByDiseaseName(DisName)
	if useer == nil {
		fmt.Println("this Disease with this name is not exist")
		BeforeReturn()
		return
	}
	fmt.Println("		Enter your chice :")
	fmt.Println("		1) disease description")
	fmt.Println("		2) disease Medicine")
	scanner.Scan()
	Choice := scanner.Text()

	switch Choice {
	case "1":
		fmt.Println("		Enter the new description : ")
		scanner.Scan()
		des := scanner.Text()
		err := changeDiseasesDescription(useer.id, des)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
	case "2":
		fmt.Println("		Enter the new Diseases Medicine : ")
		scanner.Scan()
		des := scanner.Text()
		MedId, _ := GetMedicineByMedicineName(des)
		if MedId == nil {
			fmt.Println("there is no Medicine with this name")
			BeforeReturn()
			return
		}

		err := changeDiseasesMedicine(useer.id, MedId.Id)
		if err != nil {
			fmt.Print(err)
			BeforeReturn()
			return
		}
	default:
		fmt.Println("Invalid option. Please try again.")
	}
}

func ShowPaitentDesies() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter Patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, _ := GetUserByUsername(UsreId)
	if useer != nil {
		fmt.Println("this user with this id is already exist")
		BeforeReturn()
		return
	}
	diseases, err := GetPaitentDesiesBypatientId(useer.Id)
	if err != nil {
		fmt.Print(err)
		BeforeReturn()
		return
	}
	for i := 0; i < len(diseases); i++ {
		disease, _ := GetDiseaseByDiseaseId(diseases[i].diseases_Id)
		Medicine, _ := GetMedicineByMedicineId(disease.disease_Medicine_id)
		fmt.Println("disease Name : " + disease.diseaseName)
		fmt.Println("disease description : " + disease.disease_description)
		fmt.Println("disease Medicine : " + Medicine)
		if i != len(diseases)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
		if len(diseases) == 0 {
		fmt.Println("this pateint have no diseases")
	}
			BeforeReturn()
		return
}
