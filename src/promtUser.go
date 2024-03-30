package main

import (
	"fmt"
	"log"
	"strings"
)

type Answers struct {
	tartgetFile           string
	targetDomains         []string
	enableEnumSubDomains  bool
	enablePasswordAttacks bool
}

func (ref *Answers) initAnswers() {
	ref.tartgetFile = ref.setTartgetFile()
	ref.targetDomains = ref.setTartgetDomains()
	ref.enableEnumSubDomains = ref.setEnableSubDomains()
	ref.enablePasswordAttacks = ref.setEnablePasswordAttacks()
}

func (ref *Answers) setTartgetFile() string {

	input := ""
	fmt.Print("Target File Location: ")
	isGood := false

	for !isGood {
		isGood = checkInput(&input)
	}

	return input
}

func (ref *Answers) setTartgetDomains() []string {

	input := make([]string, 0)
	fmt.Print("Target File Location: ")
	num, err := fmt.Scan(&input)

	if err != nil {
		log.Fatal(num, err)
	}

	return input
}

func (ref *Answers) setEnableSubDomains() bool {

	return ref.getUserInputForYesNoQuestions("Would you like to enumerate sub domains (Y/N): ")
}

func (ref *Answers) setEnablePasswordAttacks() bool {

	return ref.getUserInputForYesNoQuestions("Would you like to do password attacks (Y/N): ")
}

func (ref *Answers) getUserInputForYesNoQuestions(question string) bool {
	var ans bool
	var input string
	sentinal := false

	for sentinal {

		fmt.Print(question)
		fmt.Scan(&input)

		input = strings.ToLower(input)
		input = strings.TrimSpace(input)

		if input == "y" || input == "yes" {
			ans = true
			sentinal = false
		} else if input == "n" || input == "no" {
			ans = false
			sentinal = false
		} else {
			fmt.Println("Please enter Y for yes or N for no.")
		}
	}

	return ans
}

func checkInput(input *string) bool {
	num, err := fmt.Scan(input)
	isGood := true

	if err != nil {
		log.Fatal(num, err)
		fmt.Println("Please try again")
		isGood = false
	}

	return isGood
}
