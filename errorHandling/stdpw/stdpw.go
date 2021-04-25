package stdpw

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type PasswordError struct {
	msg string
}

func (pwe *PasswordError) Error() string {
	return pwe.msg
}

func PasswordValidation(pw string) error {

	pwError := &PasswordError{}

	// Check length
	if e := checkLength(pw); e != nil {
		// ....concat error message.
		pwError.msg = e.Error() + "\n"
	}

	// check contain small letter
	if e := containSmallLetter(pw); e != nil {
		// ....concat error message.
		pwError.msg = pwError.msg + e.Error() + "\n"
	}
	// check contain capital letter
	if e := containCapitalLetter(pw); e != nil {
		// ....concat error message.
		pwError.msg = pwError.msg + e.Error() + "\n"
	}
	// check contain digit
	if e := containDigit(pw); e != nil {
		// ....concat error message.
		pwError.msg = pwError.msg + e.Error() + "\n"
	}

	if pwError.msg != "" {
		return pwError
	}
	return nil
}

func checkLength(pw string) error {
	pwLen := utf8.RuneCountInString(pw)
	if pwLen < 7 || pwLen > 16 {
		return fmt.Errorf("password length should be 7-16. Your is %d", pwLen)
	}
	return nil

}

func containSmallLetter(pw string) error {
	return regexHelper(pw, `[a-z]`, "password must contain small letter")
}

func containCapitalLetter(pw string) error {
	return regexHelper(pw, `[A-Z]`, "password must contain capital letter")
}

func containDigit(pw string) error {
	return regexHelper(pw, `[0-9]`, "password must contain digit")
}

func regexHelper(pw, pettern, msg string) error {
	re := regexp.MustCompile(pettern)
	result := re.FindString(pw)
	if result == "" {
		return fmt.Errorf(msg)
	}
	return nil
}
