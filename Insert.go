package main

import (
	"log"
)

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func CreateUser(u User) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO User 
							(UserId, first_name, last_name,Gender, date_of_birth, email, hashed_password, Rule) 
							VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.UserId,
		u.first_name,
		u.last_name,
		u.Gender,
		u.date_of_birth,
		u.email,
		u.hashed_password,
		u.Rule)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func CreateAppointment(u Appointment) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO Appointment 
							(Doctor_id, user_id, date_of_Appointment,TheTime) 
							VALUES (?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.doctor_id,
		u.user_id,
		u.date_of_Appointment,
		u.TheTime)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func RequestTest(u Test) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO Test 
							(Doctor_id, user_id,lab_Tech, date_of_Appointment,TheTime) 
							VALUES (?, ?, ?, ?,?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.doctor_id,
		u.user_id,
		u.lab_Tech,
		u.Test_Type,
		u.Test_Date)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}
	stmt, err = db.Prepare("SELECT id FROM Test WHERE user_id = ? AND Test_Date = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	var testID int

	err = stmt.QueryRow(u.user_id, u.Test_Date).Scan(&testID)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	// Prepare the SQL statement
	stmt, err = db.Prepare(`INSERT INTO TestTrack 
							(Test_id,test_Statues) 
							VALUES (?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		testID,
		"test requested")

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}
	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func CreateMedicine(u Medicine) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO Medicine 
							(MedicineName, Medicine_description) 
							VALUES (?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.MedicineName,
		u.Medicine_description)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func CreatePrescription(u prescription) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO prescription 
							(Doctor_id, patient_Id,Medicine_id,prescription) 
							VALUES (?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.doctor_id,
		u.patient_Id,
		u.Medicine_id,
		u.date_of_prescription,
		u.prescription)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func Creatwvacation(u vacation) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO vacation 
							(Doctor_id,patient_Id, Number_Of_Days,Start_from,VacationReason) 
							VALUES (?,?, ?, ?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.Doctor_id,
		u.patient_Id, u.Number_Of_Days,
		u.Start_from,
		u.VacationReason)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func CreatDiseases(u diseases) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO diseases 
							(diseaseName, disease_description,disease_Medicine_id) 
							VALUES (?, ?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.diseaseName,
		u.disease_description,
		u.disease_Medicine_id)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func CreatPaitentDesies(u paitentDesies) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO paitentDesies 
							(patient_Id, diseases_Id) 
							VALUES (?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.patient_Id,
		u.diseases_Id)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}

// Append New User Info the database, if any error occurs while appending (preparing and executing the SQL statements) it will return an error
func CreatTestResults(u TestResults) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare(`INSERT INTO TestResults 
							(Test_Id, Result) 
							VALUES (?, ?)`)

	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the user's data
	_, err = stmt.Exec(
		u.Test_id,
		u.result)

	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}

	return nil
}
