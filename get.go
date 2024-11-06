package main

import (
	"database/sql"
	"log"
)

func GetUserByUsername(userId string) (*User, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,UserId, first_name, last_name,Gender, date_of_birth, email, hashed_password, Rule FROM User WHERE UserId = ?`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u User
	err = stmt.QueryRow(userId).Scan(
		&u.Id,
		&u.UserId,
		&u.first_name,
		&u.last_name,
		&u.Gender,
		&u.date_of_birth,
		&u.email,
		&u.hashed_password,
		&u.Rule)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetUserById(userId string) (*User, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,UserId, first_name, last_name,Gender, date_of_birth, email, hashed_password, Rule FROM User WHERE id = ?`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u User
	err = stmt.QueryRow(userId).Scan(
		&u.Id,
		&u.UserId,
		&u.first_name,
		&u.last_name,
		&u.Gender,
		&u.date_of_birth,
		&u.email,
		&u.hashed_password,
		&u.Rule)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetUserByEmail(email string) (*User, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,UserId, first_name, last_name,Gender, date_of_birth, email, hashed_password, Rule FROM User WHERE email = ?`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u User
	err = stmt.QueryRow(email).Scan(
		&u.Id,
		&u.UserId,
		&u.first_name,
		&u.last_name,
		&u.Gender,
		&u.date_of_birth,
		&u.email,
		&u.hashed_password,
		&u.Rule)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetVacationByVacationId(email string) (*vacation, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,Doctor_id, patient_Id, Number_Of_Days,Start_from, VacationReason FROM vacation WHERE id = ?`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u vacation
	err = stmt.QueryRow(email).Scan(
		&u.id,
		&u.Doctor_id,
		&u.patient_Id,
		&u.Number_Of_Days,
		&u.Start_from,
		&u.VacationReason)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetUserNoByUserId(UserId string) (int, error) {
	// Prepare the SQL statement to select the password hash for the given username
	stmt, err := db.Prepare("SELECT id FROM User WHERE UserId = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return 0, err
	}
	defer stmt.Close()

	var id int

	err = stmt.QueryRow(UserId).Scan(&id)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return 0, err
	}

	return id, nil
}

func GetUsersByGender(Gender, rule string) ([]User, error) {
	// Prepare the SQL statement to get appointments with specified userId
	stmt, err := db.Prepare(`SELECT id,UserId, first_name,
		 last_name,Gender, date_of_birth, email
		  FROM User WHERE Gender = ? AND Rule = ?`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(Gender, rule)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var Users []User
	for rows.Next() {
		var User User
		err = rows.Scan(&User.Id,
			&User.UserId,
			&User.first_name,
			&User.last_name,
			&User.Gender,
			&User.date_of_birth,
			&User.email)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		Users = append(Users, User)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return Users, nil
}

func GetLabatory() ([]User, error) {
	// Prepare the SQL statement to get appointments with specified userId
	stmt, err := db.Prepare(`SELECT id,UserId, first_name,
		 last_name,Gender, date_of_birth, email
		  FROM User WHERE Rule = 'Lab Techment'`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var Users []User
	for rows.Next() {
		var User User
		err = rows.Scan(&User.Id,
			&User.UserId,
			&User.first_name,
			&User.last_name,
			&User.Gender,
			&User.date_of_birth,
			&User.email)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		Users = append(Users, User)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return Users, nil
}

func GetDoctor() ([]User, error) {
	// Prepare the SQL statement to get appointments with specified userId
	stmt, err := db.Prepare(`SELECT id,UserId, first_name,
		 last_name,Gender, date_of_birth, email
		  FROM User WHERE Rule = 'doctor`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var Users []User
	for rows.Next() {
		var User User
		err = rows.Scan(&User.Id,
			&User.UserId,
			&User.first_name,
			&User.last_name,
			&User.Gender,
			&User.date_of_birth,
			&User.email)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		Users = append(Users, User)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return Users, nil
}

func GetPatients() ([]User, error) {
	// Prepare the SQL statement to get appointments with specified userId
	stmt, err := db.Prepare(`SELECT id,UserId, first_name,
		 last_name,Gender, date_of_birth, email
		  FROM User WHERE Rule = 'patient'`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var Users []User
	for rows.Next() {
		var User User
		err = rows.Scan(&User.Id,
			&User.UserId,
			&User.first_name,
			&User.last_name,
			&User.Gender,
			&User.date_of_birth,
			&User.email)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		Users = append(Users, User)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return Users, nil
}

func GetAppointmentsByUser(userId int) ([]Appointment, error) {
	// Prepare the SQL statement to get appointments with specified userId
	stmt, err := db.Prepare(`SELECT id, Doctor_id, user_id, date_of_Appointment, TheTime
							 FROM Appointment
							 WHERE user_id = ?
							 ORDER BY date_of_Appointment DESC`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var appointments []Appointment
	for rows.Next() {
		var appointment Appointment
		err = rows.Scan(&appointment.Id,
			&appointment.doctor_id,
			&appointment.user_id,
			&appointment.date_of_Appointment,
			&appointment.TheTime)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return appointments, nil
}

func GetAppointmentsByDoctor(userId int) ([]Appointment, error) {
	// Prepare the SQL statement to get appointments with specified userId
	stmt, err := db.Prepare(`SELECT id, Doctor_id, user_id, date_of_Appointment, TheTime
							 FROM Appointment
							 WHERE Doctor_id = ?
							 ORDER BY date_of_Appointment DESC`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var appointments []Appointment
	for rows.Next() {
		var appointment Appointment
		err = rows.Scan(&appointment.Id,
			&appointment.doctor_id,
			&appointment.user_id,
			&appointment.date_of_Appointment,
			&appointment.TheTime)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return appointments, nil
}

func GetTestByUserIdAndDate(userId, testDate string) (Test, error) {
	// Prepare the SQL statement to select the test information for the given user ID and test date
	stmt, err := db.Prepare("SELECT id, Doctor_id, user_id, lab_Tech, Test_Type, Test_Date FROM Test WHERE user_id = ? AND Test_Date = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return Test{}, err
	}
	defer stmt.Close()

	var test Test

	err = stmt.QueryRow(userId, testDate).Scan(&test.Id, &test.doctor_id, &test.user_id, &test.lab_Tech, &test.Test_Type, &test.Test_Date)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return Test{}, err
	}

	return test, nil
}

func GetTestById(TestId string) (Test, error) {
	// Prepare the SQL statement to select the test information for the given user ID and test date
	stmt, err := db.Prepare("SELECT id, Doctor_id, user_id, lab_Tech, Test_Type, Test_Date FROM Test WHERE id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return Test{}, err
	}
	defer stmt.Close()

	var test Test

	err = stmt.QueryRow(TestId).Scan(&test.Id, &test.doctor_id, &test.user_id, &test.lab_Tech, &test.Test_Type, &test.Test_Date)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return Test{}, err
	}

	return test, nil
}

func GetTestsByUserId(userId string) ([]Test, error) {
	// Prepare the SQL statement to select the tests for the given user ID
	stmt, err := db.Prepare("SELECT id, Doctor_id, user_id, lab_Tech, Test_Type, Test_Date FROM Test WHERE user_id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var test Test
		err = rows.Scan(&test.Id, &test.doctor_id, &test.user_id, &test.lab_Tech, &test.Test_Type, &test.Test_Date)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		tests = append(tests, test)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return tests, nil
}

func GetTestsByLabTech(lab_Tech string) ([]Test, error) {
	// Prepare the SQL statement to select the tests for the given user ID
	stmt, err := db.Prepare("SELECT id, Doctor_id, user_id, lab_Tech, Test_Type, Test_Date FROM Test WHERE lab_Tech = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(lab_Tech)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var test Test
		err = rows.Scan(&test.Id, &test.doctor_id, &test.user_id, &test.lab_Tech, &test.Test_Type, &test.Test_Date)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		tests = append(tests, test)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return tests, nil
}

func GetTestReslutByTestId(TestId string) (string, error) {
	// Prepare the SQL statement to select the test information for the given user ID and test date
	stmt, err := db.Prepare("SELECT TestResults WHERE Test_Id = ? ")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return "", err
	}
	defer stmt.Close()

	var result string

	err = stmt.QueryRow(TestId).Scan(&result)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return "", err
	}

	return result, nil
}

func GetTestTrackByTestId(TestId string) (string, error) {
	// Prepare the SQL statement to select the test information for the given user ID and test date
	stmt, err := db.Prepare("SELECT TestTrack WHERE Test_Id = ? ")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return "", err
	}
	defer stmt.Close()

	var result string

	err = stmt.QueryRow(TestId).Scan(&result)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return "", err
	}

	return result, nil
}

func GetPrescriptionBypatientId(patient_Id string) ([]prescription, error) {
	// Prepare the SQL statement to select the tests for the given user ID
	stmt, err := db.Prepare("SELECT id, Doctor_id, patient_Id,  _id, date_of_prescription, prescription FROM prescription WHERE patient_Id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(patient_Id)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var prescriptions []prescription
	for rows.Next() {
		var Prescription prescription
		err = rows.Scan(&Prescription.Id, &Prescription.doctor_id, &Prescription.patient_Id, &Prescription.Medicine_id, &Prescription.date_of_prescription, &Prescription.prescription)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		prescriptions = append(prescriptions, Prescription)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return prescriptions, nil
}

func GetvacationBypatientId(patient_Id string) ([]vacation, error) {
	// Prepare the SQL statement to select the tests for the given user ID
	stmt, err := db.Prepare("SELECT id, Doctor_id,patient_Id, Number_Of_Days, Start_from, VacationReason FROM vacation WHERE patient_Id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(patient_Id)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var vacations []vacation
	for rows.Next() {
		var Vacation vacation
		err = rows.Scan(&Vacation.id, &Vacation.Doctor_id, &Vacation.patient_Id, &Vacation.Number_Of_Days, &Vacation.Start_from, &Vacation.VacationReason)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		vacations = append(vacations, Vacation)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return vacations, nil
}

func GetDiseases() ([]diseases, error) {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`SELECT id, diseaseName, disease_description, disease_Medicine_id FROM diseases`)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the diseases
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var diseasesList []diseases
	for rows.Next() {
		var u diseases
		err := rows.Scan(
			&u.id,
			&u.diseaseName,
			&u.disease_description,
			&u.disease_Medicine_id)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		diseasesList = append(diseasesList, u)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return diseasesList, nil
}

func GetAppointmentsByid(id string) (*Appointment, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,Doctor_id, user_id, date_of_Appointment,TheTime FROM Appointment WHERE id = ?`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u Appointment
	err = stmt.QueryRow(id).Scan(
		&u.Id,
		&u.doctor_id,
		&u.user_id,
		&u.date_of_Appointment,
		&u.TheTime)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetPaitentDesiesBypatientId(patient_Id int) ([]paitentDesies, error) {
	// Prepare the SQL statement to select the tests for the given user ID
	stmt, err := db.Prepare("SELECT id, patient_Id, diseases_Id FROM paitentDesies WHERE patient_Id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(patient_Id)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var paitentDesiess []paitentDesies
	for rows.Next() {
		var paitentDesies paitentDesies
		err = rows.Scan(&paitentDesies.id, &paitentDesies.patient_Id, &paitentDesies.diseases_Id)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		paitentDesiess = append(paitentDesiess, paitentDesies)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return paitentDesiess, nil
}

func GetAppointmentsByDatrAndTimeAndDoctor(date, time string, id int) (*Appointment, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,Doctor_id, user_id, date_of_Appointment,TheTime FROM Appointment WHERE date_of_Appointment = ? AND TheTime = ? AND Doctor_id = ?`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u Appointment
	err = stmt.QueryRow(date, time, id).Scan(
		&u.Id,
		&u.doctor_id,
		&u.user_id,
		&u.date_of_Appointment,
		&u.TheTime)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetTests() ([]Test, error) {
	// Prepare the SQL statement to select all tests
	stmt, err := db.Prepare("SELECT id, Doctor_id, user_id, lab_Tech, Test_Type, Test_Date FROM Test")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var test Test
		err = rows.Scan(&test.Id, &test.doctor_id, &test.user_id, &test.lab_Tech, &test.Test_Type, &test.Test_Date)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		tests = append(tests, test)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return tests, nil
}

func GetMedicineByMedicineId(id int) (string, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,MedicineName, Medicine_description FROM Medicine WHERE id = ? `)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return "", err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u Medicine
	err = stmt.QueryRow(id).Scan(
		&u.Id,
		&u.MedicineName,
		&u.Medicine_description)
	if err == sql.ErrNoRows {
		return "", nil // User doesn't exist, return nil with no error
	}

	return u.MedicineName, nil
}

func GetDiseaseByDiseaseName(Name string) (*diseases, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,diseaseName, disease_Medicine_id FROM diseases WHERE diseaseName = ? `)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u diseases
	err = stmt.QueryRow(Name).Scan(
		&u.id,
		&u.diseaseName,
		&u.disease_description,
		&u.disease_Medicine_id)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetMedicineByMedicineName(id string) (*Medicine, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,MedicineName, Medicine_description FROM Medicine WHERE diseaseName = ? `)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u Medicine
	err = stmt.QueryRow(id).Scan(
		&u.Id,
		&u.MedicineName,
		&u.Medicine_description)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetDiseaseByDiseaseId(Name int) (*diseases, error) {
	// Prepare the SQL statement with a placeholder
	stmt, err := db.Prepare(`SELECT id,diseaseName, disease_Medicine_id FROM diseases WHERE id = ? `)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the user information
	var u diseases
	err = stmt.QueryRow(Name).Scan(
		&u.id,
		&u.diseaseName,
		&u.disease_description,
		&u.disease_Medicine_id)
	if err == sql.ErrNoRows {
		return nil, nil // User doesn't exist, return nil with no error
	}

	return &u, nil
}

func GetMedicine() ([]Medicine, error) {
	// Prepare the SQL statement to select all Medicines
	stmt, err := db.Prepare("SELECT id, MedicineName, Medicine_description FROM Medicine")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var Medicines []Medicine
	for rows.Next() {
		var Medicine Medicine
		err = rows.Scan(&Medicine.Id, &Medicine.MedicineName, &Medicine.Medicine_description)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		Medicines = append(Medicines, Medicine)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return Medicines, nil
}

func GetPrescriptionByDoctorId(patient_Id string) ([]prescription, error) {
	// Prepare the SQL statement to select the tests for the given user ID
	stmt, err := db.Prepare("SELECT id, Doctor_id, patient_Id,  _id, date_of_prescription, prescription FROM prescription WHERE Doctor_id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(patient_Id)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()

	var prescriptions []prescription
	for rows.Next() {
		var Prescription prescription
		err = rows.Scan(&Prescription.Id, &Prescription.doctor_id, &Prescription.patient_Id, &Prescription.Medicine_id, &Prescription.date_of_prescription, &Prescription.prescription)
		if err != nil {
			log.Printf("Error scanning row: %s\n", err.Error())
			return nil, err
		}
		prescriptions = append(prescriptions, Prescription)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %s\n", err.Error())
		return nil, err
	}

	return prescriptions, nil
}
