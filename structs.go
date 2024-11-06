package main

type User struct {
	Id              int
	UserId          string
	first_name      string
	last_name       string
	Gender          string
	date_of_birth   string
	email           string
	hashed_password string
	Rule            string
}

type Appointment struct {
	Id                  int
	doctor_id           int
	user_id             int
	date_of_Appointment string
	TheTime             string
}

type Test struct {
	Id        int
	doctor_id int
	user_id   int
	lab_Tech  int
	Test_Type string
	Test_Date string
}

type TestTrack struct {
	Id           int
	Test_id      int
	user_id      int
	Test_Statues string
}

type Medicine struct {
	Id                   int
	MedicineName         string
	Medicine_description string
}

type prescription struct {
	Id                   int
	doctor_id            int
	patient_Id           int
	Medicine_id          int
	date_of_prescription string
	prescription         string
}

type TestResults struct {
	id      int
	Test_id int
	result  string
}

type vacation struct {
	id             int
	Doctor_id      int
	patient_Id     int
	Number_Of_Days int
	Start_from     string
	VacationReason string
}

type diseases struct {
	id                  int
	diseaseName         string
	disease_description string
	disease_Medicine_id int
}

type paitentDesies struct {
	id          int
	patient_Id  int
	diseases_Id int
}
