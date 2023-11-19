package utils

import "golang.org/x/crypto/bcrypt"

type PasswordEncoder interface {
	Encode(rawPwd string) (string, error)
	Matches(rawPwd, encodePwd string) bool
}

type PasswordUtil struct {
}

func (PasswordUtil) Encode(rawPwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (PasswordUtil) Matches(rawPwd, encodePwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(rawPwd)) == nil
}
