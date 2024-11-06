package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func doctor() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		fmt.Println("Enter Your Choice:")
		fmt.Println("				1) show  Appointments")
		fmt.Println("				2) show patient's infornation")
		fmt.Println("				3) test request")
		fmt.Println("				4) test tracking")
		fmt.Println("				5) test results")
		fmt.Println("				6) add patient vacation")
		fmt.Println("				7) Modefi patient's vacation")
		fmt.Println("				8) Modefi patient's desies")
		fmt.Println("				9) add patient prescription")
		fmt.Println("				10) show Medicines")
		fmt.Println("				11) show patient prescriptions")
		fmt.Println("				12) show  patients")
		fmt.Println("				13) show  prescriptions added by you")
		fmt.Println("				14) show patient's tests")

		scanner.Scan()          // Read the input from the user
		casee := scanner.Text() // Get the user's input as a string

		switch casee {
		case "1":
			showAppointments()
		case "2":
			showPatientInfornation()
		case "3":
			testRequest()
		case "4":
			TrackTest()
		case "5":
			testResults()
		case "6":
			addPatientVacation()
		case "7":
			ModefiPatientVacation()
		case "8":
			ModefiPatientDesies()
		case "9":
			addPatientPrescription()
		case "10":
			showMedicines()
		case "11":
			showPatientPrescriptions()
		case "12":
			showPatients()
		case "13":
			showPrescriptionsAddedByYou()
		case "14":
			ShowPaitentTests()

		default:
			fmt.Println("Invalid option. Please try again.")
		}

	}
}

func TrackTest() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter Test Id :\n#  ")
	scanner.Scan()
	TestId := scanner.Text()

	tracking, err := GetTestTrackByTestId(TestId)
	if err != nil {
		fmt.Print(err)
		BeforeReturn()
		return
	}

	fmt.Print("Test statues :" + tracking)

}

func showAppointments() {
	clearScreen()
	appointments, _ := GetAppointmentsByDoctor(LogedUser.Id)

	for i := 0; i < len(appointments); i++ {
		user, _ := GetUserById(strconv.Itoa(appointments[i].user_id))
		fmt.Println("patient Name : " + user.first_name + " " + user.last_name)
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

func showPatientInfornation() {
	clearScreen()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter the patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, _ := GetUserByUsername(UsreId)
	if useer == nil || useer.Rule != "patient" {
		fmt.Println("this patient with this id is not exist")

		fmt.Println("\n\n               **enter anuthing to return")
		scanner.Scan()
		casee := scanner.Text()
		if casee != "" {
			return
		}
		return
	}
	Desies, _ := GetPaitentDesiesBypatientId(useer.Id)
	fmt.Println("patient name : " + useer.first_name + " " + useer.last_name)
	fmt.Println("patient Age :" + ReturnAge(strings.ReplaceAll(useer.date_of_birth, "T00:00:00Z", "")))
	fmt.Println("patient Birth day :" + ReturnAge(strings.ReplaceAll(useer.date_of_birth, "T00:00:00Z", "")))
	fmt.Println("patient Desieses :")
	for i := 0; i < len(Desies); i++ {
		d, _ := GetDiseaseByDiseaseId(Desies[i].diseases_Id)
		fmt.Print(d.diseaseName)
		if i != len(Desies)-1 {
			fmt.Println("      -      ")
		}
	}
	if len(Desies) == 0 {
		fmt.Println("thie paitent have no desieas")
	}
	fmt.Println("\n\n               **enter anuthing to return")
	scanner.Scan()
	casee := scanner.Text()
	if casee != "" {
		return
	}

}

func ReturnAge(birthdateStr string) string {
	currentTime := time.Now()
	// Parse birth date string into a time.Time value
	birthdate, err := time.Parse("2006-01-02", birthdateStr)
	if err != nil {
		fmt.Println("Error parsing birthdate:", err)
		return ""
	}

	// Calculate age
	age := currentTime.Year() - birthdate.Year()

	// Check if the birthday has occurred this year or not
	if currentTime.Before(time.Date(currentTime.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, time.UTC)) {
		age--
	}

	return strconv.Itoa(age)
}

func testRequest() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter the patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, _ := GetUserByUsername(UsreId)
	if useer == nil || useer.Rule != "patient" {
		fmt.Println("this patient with this id is not exist")
		BeforeReturn()
		return
	}

	labs, _ := GetLabatory()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random index within the range of the array
	randomIndex := rand.Intn(len(labs))

	// Select the random element
	lab := labs[randomIndex]

	fmt.Print("		Enter Tet Type :\n  ")
	scanner.Scan()
	TestType := scanner.Text()
	currentTime := time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)
	test := Test{
		doctor_id: LogedUser.Id,
		user_id:   useer.Id,
		lab_Tech:  lab.Id,
		Test_Type: TestType,
		Test_Date: strings.ReplaceAll(currentDate.String(), " 00:00:00 +0000 UTC", ""),
	}
	err := RequestTest(test)
	if err != nil {
		fmt.Println(err)
		BeforeReturn()
		return
	}
}

func testResults() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter Test Id :\n#  ")
	scanner.Scan()
	TestId := scanner.Text()

	_, err := GetTestById(TestId)
	if err != nil {
		fmt.Print("this test with this id dose not exist")
		BeforeReturn()
		return
	}

	result, err := GetTestReslutByTestId(TestId)
	if err != nil || result == "" {
		fmt.Print("The result is not out yet")
		BeforeReturn()
		return
	}

	fmt.Println("test #" + TestId + " result : " + result)
	BeforeReturn()
	return
}

func addPatientVacation() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, _ := GetUserByUsername(UsreId)
	if useer == nil || useer.Rule != "patient" {
		fmt.Println("this patient with this id is not exist")
		BeforeReturn()
		return
	}

	fmt.Print("		Enter days number :\n  ")
	scanner.Scan()
	days := scanner.Text()
	intDays, err := strconv.Atoi(days)
	if err != nil {
		fmt.Println("invalid days")
		BeforeReturn()
		return
	}

	fmt.Print("		Enter Vacation Reason:\n  ")
	scanner.Scan()
	VacationReason := scanner.Text()
	date := strings.ReplaceAll(time.Now().Format("2006-01-02"), "T00:00:00Z", "")
	vacation := vacation{
		Doctor_id:      LogedUser.Id,
		patient_Id:     useer.Id,
		Number_Of_Days: intDays,
		Start_from:     date,
		VacationReason: VacationReason,
	}

	err = Creatwvacation(vacation)
	if err != nil {
		fmt.Println(err)
		BeforeReturn()
		return

	}

	BeforeReturn()
	return

}

func ModefiPatientVacation() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, _ := GetUserByUsername(UsreId)
	if useer == nil || useer.Rule != "patient" {
		fmt.Println("this patient with this id is not exist")
		BeforeReturn()
		return
	}

	vacations, err := GetvacationBypatientId(strconv.Itoa(useer.Id))
	if err != nil {
		fmt.Println(err)
		BeforeReturn()
		return
	}

	for i := 0; i < len(vacations); i++ {
		fmt.Println("vcation no." + strconv.Itoa(vacations[i].id))
		fmt.Println("number of days" + strconv.Itoa(vacations[i].Number_Of_Days))
		fmt.Println("date" + vacations[i].Start_from)
		fmt.Println("Vacation Reason : " + vacations[i].VacationReason)

		if i != len(vacations)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}

	fmt.Print("		Enter vacation Id :\n  ")
	scanner.Scan()
	VacationId := scanner.Text()

	vacation, err := GetVacationByVacationId(VacationId)
	if err != nil || vacation.patient_Id != useer.Id {
		fmt.Println("invalid vacation Id")
		return
	}

	fmt.Println("Enter Your choice :")
	fmt.Println("			1) vacation reason ")
	fmt.Println("			2) number of dayse ")

	scanner.Scan()
	casee := scanner.Text()

	switch casee {
	case "1":
		fmt.Print("		Enter vacation reason :\n  ")
		scanner.Scan()
		days := scanner.Text()
		err = changeVacation2(VacationId, days)

		if err != nil {
			fmt.Println(err)
			return
		}
	case "2":

		fmt.Print("		Enter days number :\n  ")
		scanner.Scan()
		days := scanner.Text()
		intDays, err := strconv.Atoi(days)
		err = changeVacation(VacationId, intDays)
		if err != nil {
			fmt.Println("invalid days")
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	default:
		fmt.Println("Invalid option. Please try again.")

	}

}

func addPatientPrescription() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, err := GetUserByUsername(UsreId)
	if useer == nil || useer.Rule != "patient" {
		fmt.Println("this patient with this id is not exist")
		return
	}
	fmt.Print("		Enter Medicine Name Id :\n  ")
	scanner.Scan()
	medId := scanner.Text()
	med, err := GetMedicineByMedicineName(medId)
	if med == nil || err != nil {
		fmt.Println("this medicine with this name dose not exist")
		return
	}

	fmt.Print("		Enter prescription :\n  ")
	scanner.Scan()
	prescriptions := scanner.Text()

	medi := prescription{
		doctor_id:            LogedUser.Id,
		patient_Id:           useer.Id,
		Medicine_id:          med.Id,
		date_of_prescription: strings.ReplaceAll(time.Now().Format("2006-01-02"), "T00:00:00Z", ""),
		prescription:         prescriptions,
	}

	err = CreatePrescription(medi)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func ModefiPatientDesies() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, err := GetUserByUsername(UsreId)
	if useer == nil || useer.Rule != "patient" {
		fmt.Println("this patient with this id is not exist")
		return
	}
	fmt.Print("		Enter the desies you want to delete/add :\n  ")
	scanner.Scan()
	desId := scanner.Text()
	des, err := GetDiseaseByDiseaseName(desId)
	if err != nil || des == nil {
		fmt.Print("err to find")
	}

	fmt.Println("Enter Your choice :")
	fmt.Println("			1) delete ")
	fmt.Println("			2) add ")

	scanner.Scan()
	casee := scanner.Text()

	switch casee {
	case "1":

		err := removePaitentDesies(strconv.Itoa(useer.Id), strconv.Itoa(des.id))
		if err != nil {
			fmt.Print("err to delete")
			return
		}
	case "2":

		paiDes := paitentDesies{
			patient_Id:  useer.Id,
			diseases_Id: des.id,
		}

		err := CreatPaitentDesies(paiDes)
		if err != nil {
			fmt.Print("err to add")
			return
		}

	default:
		fmt.Println("Invalid option. Please try again.")

	}
}

func showMedicines() {
	med, err := GetMedicine()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(med); i++ {
		fmt.Println("Medicine Id: #%d", med[i].Id)
		fmt.Println("[i]icine Name: ", med[i].MedicineName)
		fmt.Println("Medicine description: ", med[i].Medicine_description)

		if i != len(med)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
		if len(med) == 0 {
		fmt.Println("no medisined yet in the data base")
	}
			BeforeReturn()
		return
}

func showPatientPrescriptions() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter patient Id :\n  ")
	scanner.Scan()
	UsreId := scanner.Text()
	useer, err := GetUserByUsername(UsreId)
	if useer.Id == 0 || useer.Rule != "patient" || err != nil {
		fmt.Println("this patient with this id is not exist")
		return
	}
	pres, err := GetPrescriptionBypatientId(strconv.Itoa(useer.Id))
	if err != nil {
		fmt.Println("no patient with this id")
		return
	}
	for i := 0; i < len(pres); i++ {
		fmt.Print("the prescription date :" + pres[i].date_of_prescription)
		med, _ := GetMedicineByMedicineId(pres[i].Medicine_id)
		fmt.Print("the Medicine :" + med)
		fmt.Print("the prescription :" + pres[i].prescription)
		if i != len(pres)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}

		if len(pres) == 0 {
		fmt.Println("no prescriptions for this patient")
	}
			BeforeReturn()
		return

}

func showPatients() {
	clearScreen()

	patients, _ := GetPatients()
	for i := 0 ; i < len(patients) ; i++ {
		fmt.Println("patient Id : " + patients[i].UserId)
		fmt.Println("patient name : " + patients[i].first_name + " " + patients[i].last_name)
		fmt.Println("patient birthday :" +strings.ReplaceAll(patients[i].date_of_birth, "T00:00:00Z", ""))
		fmt.Println("patient Age :" + ReturnAge(strings.ReplaceAll(patients[i].date_of_birth, "T00:00:00Z", "")))
				if i != len(patients)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
	if len(patients) == 0 {
		fmt.Println("no patients")
	}
			BeforeReturn()
		return
}

func showPrescriptionsAddedByYou() {
	pres, err := GetPrescriptionByDoctorId(strconv.Itoa(LogedUser.Id))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(pres); i++ {
		userById, _ := GetUserById(strconv.Itoa(pres[i].patient_Id))
		fmt.Println("patient name : " + userById.first_name + " " + userById.last_name)
		fmt.Print("the prescription date :" + pres[i].date_of_prescription)
		med, _ := GetMedicineByMedicineId(pres[i].Medicine_id)
		fmt.Print("the Medicine :" + med)
		fmt.Print("the prescription :" + pres[i].prescription)
		if i != len(pres)-1 {
			fmt.Println("-------------------------------------------------------------")
		}
	}
	if len(pres) == 0 {
		fmt.Println("no pres yet")
	}
			BeforeReturn()
		return
}
