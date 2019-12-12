package password

import (
	"strconv"
	"strings"
)

// Password is an integer that must be valid
type Password struct {
	Value  []int
	Valid  bool
	Reason []string
}

//New intacs a string and outputs a valid password if the string qualifies as a password.
func New(potentialPassword string) Password {
	passString := strings.Split(potentialPassword, "")
	passInt := []int{}
	for i := 0; i < len(passString); i++ {
		j, err := strconv.Atoi(passString[i])
		if err != nil {
			panic(err)
		}
		passInt = append(passInt, j)
	}
	newPassWord := Password{Value: passInt, Valid: false, Reason: []string{}}
	newPassWord.Valid = newPassWord.Verify()
	return newPassWord
}

//Verify checks a potentialPassword to ensure that it is valid
func (potentialPassword Password) Verify() bool {
	valid := true
	if !checkRuleDoubleReq(potentialPassword) {
		valid = false
	}

	if !checkRuleNeverDecrease(potentialPassword) {
		valid = false
	}
	return valid
}

func checkRuleNeverDecrease(potentialPassword Password) bool {
	valid := true
	for i := 0; i < len(potentialPassword.Value)-1 && valid != false; i++ {
		if potentialPassword.Value[i] > potentialPassword.Value[i+1] {
			//fmt.Printf("\nComparing digit: %d To digit %d", potentialPassword.Value[i], potentialPassword.Value[i+1])
			valid = false
		}
	}
	return valid
}

func checkRuleDoubleReq(potentialPassword Password) bool {
	valid := false
	for i := 0; i < len(potentialPassword.Value)-1 && valid != true; i++ {
		if potentialPassword.Value[i] == potentialPassword.Value[i+1] {
			if i > 0 && potentialPassword.Value[i] == potentialPassword.Value[i-1] {

			} else if i < len(potentialPassword.Value)-2 && potentialPassword.Value[i] == potentialPassword.Value[i+2] {

			} else {
				valid = true
			}
		}
	}
	return valid
}
