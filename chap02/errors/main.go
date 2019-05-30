package main

import (
	"errors"
	"fmt"
)

func Login(User map[string]string) (map[string]string, error) {
	if _, found := User["userName"]; found == false {
		return nil, errors.New("No UserName Data")
	}
	if userName := User["userName"]; userName != "gauarv" {
		errorObj := fmt.Errorf("User %s does not exist", userName)
		return nil, errorObj
	}
	return User, nil
}

func main() {
	user := map[string]string{}
	_, err := Login(user)
	fmt.Println("Error:", err, "\nError Message:", err.Error())
	user["userName"] = "testName"
	_, err = Login(user)
	fmt.Println("Error:", err, "\nError Message:", err.Error())
	fmt.Println("Variable err value:", err)
}
