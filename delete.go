package main

import (
	"log"
)

// Remove reaction from Post by taking ReactionPost ID
func removeUser(UserId string) error {
	stmt, err := db.Prepare("DELETE FROM User WHERE UserId = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(UserId)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}
	return nil
}

// Remove reaction from Post by taking ReactionPost ID
func removeAppointments(UserId, Appointment, date_of_Appointment, time string) error {
	stmt, err := db.Prepare("DELETE FROM Appointment WHERE UserId = ? AND date_of_Appointment = ? AND TheTime = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(UserId, date_of_Appointment, time)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}
	return nil
}

// Remove reaction from Post by taking ReactionPost ID
func removePaitentDesies(patient_Id, diseases_Id string) error {
	stmt, err := db.Prepare("DELETE FROM paitentDesies WHERE patient_Id = ? AND diseases_Id = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(patient_Id, diseases_Id)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}
	return nil
}

// Remove reaction from Post by taking ReactionPost ID
func removeVacation(User_id, Start_from string) error {
	stmt, err := db.Prepare("DELETE FROM vacation WHERE User_id = ? AND Start_from = ?")
	if err != nil {
		log.Printf("Error preparing SQL statement: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(User_id, Start_from)
	if err != nil {
		log.Printf("Error executing SQL statement: %s\n", err.Error())
		return err
	}
	return nil
}
