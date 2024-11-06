package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func LabTech() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		fmt.Println("Enter Your Choice:")
		fmt.Println("				1) show  Tests")
		fmt.Println("				2) show paitent's infornation")
		fmt.Println("				3) update test tracking")
		fmt.Println("				4) add test results")
		fmt.Println("				5) update test results")
		fmt.Println("				6) add diseases")
		fmt.Println("				7) modifi patient diseases")
		scanner.Scan()          // Read the input from the user
		casee := scanner.Text() // Get the user's input as a string

		switch casee {
		case "1":
			ShowTests()
		case "2":
			showPatientInfornation()
		case "3":
			updateTestTracking()
		case "4":
			addTestResults()
		case "5":
			updateTestResults()
		case "6":
			addDiseases()
		case "7":
			ModefiPatientDesies()
		default:
			fmt.Println("Invalid option. Please try again.")
		}

	}
}

func updateTestTracking() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter Test Id :\n#  ")
	scanner.Scan()
	TestId := scanner.Text()
	fmt.Print("		Enter The new statues :\n  ")
	scanner.Scan()
	statues := scanner.Text()
	err := changeTestStatues(TestId, statues)
	if err != nil {
		fmt.Println("couldn't update , id test wrong or some inner problem")
		BeforeReturn()
		return
	}

}

func addTestResults() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter Test Id :\n#  ")
	scanner.Scan()
	TestId := scanner.Text()
	tid, err := strconv.Atoi(TestId)
	if err != nil {
		fmt.Println("invalid test id")
		BeforeReturn()
		return
	}
	_, err = GetTestById(TestId)
	if err != nil {
		fmt.Println("error to find this test/ this test dose not exist")
		BeforeReturn()
		return
	}
	OldResult, _ := GetTestReslutByTestId(TestId)
	if OldResult != "" {
		fmt.Println("this test result already out , try to update it if you want to put a new one")
		BeforeReturn()
		return
	}
	fmt.Print("		Enter Test results :\n  ")
	scanner.Scan()
	result := scanner.Text()

	tr := TestResults{
		Test_id: tid,
		result:  result,
	}
	err = CreatTestResults(tr)
	if err != nil {
		fmt.Println("error to find this test/ this test dose not exist")
		BeforeReturn()
		return
	}
}

func updateTestResults() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("		Enter Test Id :\n#  ")
	scanner.Scan()
	TestId := scanner.Text()
	_, err := strconv.Atoi(TestId)

	if err != nil {
		fmt.Println("invalid test id")
		BeforeReturn()
		return
	}
	_, err = GetTestById(TestId)
	if err != nil {
		fmt.Println("error to find this test/ this test dose not exist")
		BeforeReturn()
		return
	}
	OldResult, err := GetTestReslutByTestId(TestId)
	if OldResult == "" || err != nil {
		fmt.Println("this test result is not out yet , try to add new one")
		BeforeReturn()
		return
	}
	fmt.Print("		Enter Test results :\n  ")
	scanner.Scan()
	result := scanner.Text()
	err = changeTestResults(TestId, result)
	if err != nil {
		fmt.Println("error to find this test/ this test dose not exist")
		BeforeReturn()
		return
	}
}
