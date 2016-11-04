package tld

import (
	"errors"
	"regexp"
	"strings"
)

var validMail = regexp.MustCompile(`^(?i)[a-z]+\@([a-z]+\.)+[a-z]{2,4}$`)

// LocateEmail takes the tld from given email adress and, returnes a country it
// represents based on the information in CountryTLD interface
func LocateEmail(mail string, cslice CountryTLD) (string, error) {
	if !validateEmail(mail) {
		return "", errors.New("invalid email format")
	}
	dot := strings.LastIndex(mail, ".")
	tld := mail[dot+1:]
	return cslice.CheckTLD(tld), nil
}

func validateEmail(mail string) bool {
	return validMail.MatchString(mail)
}
