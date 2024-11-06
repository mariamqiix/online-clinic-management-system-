package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	// "strings"
	"time"
)

func main() {
	dbPath := "/home/mariam/clinic-management-system/database.sqlite"
	err := Connect(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// 	user := Apointment{
	// 		Doctor_id:      "1",
	// 		user_id:       "5",
	// 		date_of_Apointment:   FormatDate("15-12-2004"),
	// 		TheTime:           "noali@gmail.com"}

	// CreateAppointment()

	// user, _ := GetApointmentsByDatrAndTiem("2024-05-08", "10:10")
	// fmt.Print(user.date_of_Apointment)
	// currentTime := time.Now()
	// currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)
	// fmt.Println("Current Date:", strings.ReplaceAll(currentDate.String(), " 00:00:00 +0000 UTC", ""))
	// fmt.Println("Current Date:", currentDate)

	server()
	// fmt.Println("Database created successfully!")
	// User1, err := GetUserByUsername("041206789")
	// fmt.Println(User1)

	// // user := User{
	// // 	UserId:          "041208412",
	// // 	first_name:      "noor",
	// // 	last_name:       "ali",
	// // 	Gender:          "Famle",
	// // 	date_of_birth:   FormatDate("15-12-2004"),
	// // 	email:           "noali@gmail.com",
	// // 	hashed_password: "1234567890",
	// // 	Rule:            "patient"}

	// // CreateUser(user)
	// updateUserID("0412067899", "041206789")
	// User1, err = GetUserByUsername("041206789")
	// fmt.Println(User1)

}

func FormatDate(date1 string) string {

	// Parse the date string into a time.Time value
	t, err := time.Parse("2006-01-02", date1)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	// Format the time value using the desired layout
	formattedStr := t.Format("2006-01-02T15:04:05Z")

	return formattedStr // Specify the desired date format
}

var db *sql.DB

// Connects to the database, if an error happens exists with status 1
func Connect(dbPath string) error {
	// sql.Open wont error if file not found
	fi, err := os.Stat(dbPath)
	if err != nil || fi.IsDir() {
		return errors.New("database file not found")
	}
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", dbPath)
	ldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		msg := fmt.Sprintf("can't connect to database: %s", err.Error())
		return errors.New(msg)
	}
	db = ldb
	return nil
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// "time"
// 	_ "github.com/mattn/go-sqlite3"
// )

// var db *sql.DB

// func FormatDate(date1 string) string {
// 	layout := "02-01-2006" // Specify the layout of the input date string

// 	// Parse the date string into a time.Time value
// 	date, _ := time.Parse(layout, date1)
// 	return date.Format("02/01/2006") // Specify the desired date format
// }

// func main() {
// 	// Open a connection to the SQLite database
// 	db, err := sql.Open("sqlite3", "database.sqlite")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Read the SQL file
// 	sqlFile, err := ioutil.ReadFile("/home/mariam/clinic-management-system/database/create_db.sql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Execute the SQL statements in the file
// 	queries := string(sqlFile)
// 	_, err = db.Exec(queries)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Database created successfully.")
// }
