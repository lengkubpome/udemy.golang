package main

import (
	"fmt"

	// "udemy.golang/errorHandling/stdpw"
	"example.com/errorhandling/stdpw"
)

type PasswordError struct {
	msg string
}

func (pwe *PasswordError) Error() string {
	return pwe.msg
}

func main() {
	if err := stdpw.PasswordValidation("$"); err != nil {
		// fmt.Println("%v", err.Error())

		// ถ้าต้องการเข้าถึงโครงสร้างของ password error เพื่อที่จะได้จัดการต่อ
		fmt.Printf("%+v \n", *(err.(*stdpw.PasswordError)))
		// handling error when invalid length
		// handling error when missing small character
		// handling error when missing capital character
		// handling error when missing digit
	}
	fmt.Println("Done")
}
