package main

func updateUserFirstName(userID int, newFirstName string) error {
	query := "UPDATE User SET first_name = ? WHERE UserId = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(newFirstName, userID)
	if err != nil {
		return err
	}

	return nil
}

func updateUserlasttName(userID string, newFirstName string) error {
	query := "UPDATE User SET last_name = ? WHERE UserId = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(newFirstName, userID)
	if err != nil {
		return err
	}

	return nil
}

func updateUserID(userID string, newUserId string) error {
	query := "UPDATE User SET UserId = ? WHERE UserId = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(newUserId, userID)
	if err != nil {
		return err
	}

	return nil
}

func updateUserBirthDay(userID string, newUserId string) error {
	query := "UPDATE User SET UserId = ? WHERE UserId = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(newUserId, userID)
	if err != nil {
		return err
	}

	return nil
}

func updateUserEmail(userID string, newUserId string) error {
	query := "UPDATE User SET email = ? WHERE UserId = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(newUserId, userID)
	if err != nil {
		return err
	}

	return nil
}

func updatePassword(userID string, newUserId string) error {
	query := "UPDATE User SET hashed_password = ? WHERE UserId = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(newUserId, userID)
	if err != nil {
		return err
	}

	return nil
}

func ChangeLabTech(TestId string, lab_Tech string) error {
	query := "UPDATE Test SET lab_Tech = ? WHERE id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(lab_Tech, TestId)
	if err != nil {
		return err
	}

	return nil
}

func changeTestStatues(TestId string, test_Statues string) error {
	query := "UPDATE TestTrack SET test_Statues = ? WHERE Test_id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(test_Statues, TestId)
	if err != nil {
		return err
	}

	return nil
}

func changePrescription(id string, prescription string) error {
	query := "UPDATE prescription SET prescription = ? WHERE id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(prescription, id)
	if err != nil {
		return err
	}

	return nil
}

func changeTestResults(id string, TestResults string) error {
	query := "UPDATE TestResults SET Result = ? WHERE id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(TestResults, id)
	if err != nil {
		return err
	}

	return nil
}

func changeVacation(id string, Number_Of_Days int) error {
	query := "UPDATE vacation SET Number_Of_Days = ? WHERE id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(Number_Of_Days, id)
	if err != nil {
		return err
	}

	return nil
}

func changeVacation2(id string, Number_Of_Days string) error {
	query := "UPDATE vacation SET VacationReason = ? WHERE id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(Number_Of_Days, id)
	if err != nil {
		return err
	}

	return nil
}

func changeDiseasesDescription(id int, disease_description string) error {
	query := "UPDATE diseases SET disease_description = ? WHERE id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(disease_description, id)
	if err != nil {
		return err
	}

	return nil
}

func changeDiseasesMedicine(id int, disease_Medicine_id int) error {
	query := "UPDATE diseases SET disease_Medicine_id = ? WHERE id = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(disease_Medicine_id, id)
	if err != nil {
		return err
	}

	return nil
}

func ChangeMedicineDescription(MedicineName string, NewDescription string) error {
	query := "UPDATE Medicine SET Medicine_description = ? WHERE MedicineName = ?"

	// Prepare the statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(NewDescription, MedicineName)
	if err != nil {
		return err
	}

	return nil
}
