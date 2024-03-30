package main

import (
	"fmt"
	"log"
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
	num, err := fmt.Scan(&input)

	if err != nil {
		log.Fatal(num, err)
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

	var ans bool

	return ans
}

func (ref *Answers) setEnablePasswordAttacks() bool {
	var ans bool

	return ans
}
