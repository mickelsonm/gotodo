package email

import (
	"regexp"
)

func IsEmail(emailString string) bool {
	valid, _ := regexp.MatchString("\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*", emailString)
	return valid
}
